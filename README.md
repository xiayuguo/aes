# aes
AES/CBC/PKCS7Padding 加/解密的实现


## Installation
```golang
go get github.com/xiayuguo/aes
```

## Quick Start
```golang
package main

import (
    "fmt"
    "github.com/xiayuguo/aes"
)

func main(){
    const key = "your key"
    // encrypt
    output, err := aes.AesEncrypt("body", key)
    if err != nil {
        fmt.Println("AesEncrypt's result is ", output)
    }
    // decrypt
    output, err = aes.AesDecrypt("encrypt body", key)
    if err != nil {
        fmt.Println("AesDecrypt's result is ", output)
    }
}
```


package main

import (
"crypto/md5"
"fmt"
)

func main() {
h := md5.New()
fmt.Println(h)
}

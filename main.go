package main

import (
"fmt"
"os/exec"
)

func main() {
command := "echo hello"
output, _ := exec.Command("sh", "-c", command).Output()
fmt.Println(string(output))
}

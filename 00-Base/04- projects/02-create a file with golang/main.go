package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.WriteFile("filename.txt", []byte("Hello"), 0755)
    if err != nil {
        fmt.Print("unable to write file: \n", err)
    }
}
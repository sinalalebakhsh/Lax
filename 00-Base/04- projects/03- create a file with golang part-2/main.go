package main

import (
    "os"
    "log"
)


func main() {
    // If the file doesn't exist, create it, or append to the file
    f, err := os.OpenFile("txt.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
   
    _, err = f.Write([]byte("\nsss"))
    if err != nil {
        log.Fatal(err)
    }

    f.Close()
}
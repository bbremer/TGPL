package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args)
    var args []string
    if len(os.Args) == 0 {
        fmt.Println("No args. TODO: Read from STDIN.")
    } else {
        args = os.Args
    }
}

package main

import (
    "fmt"
    "os"
)

const EXIT_FAILURE = 1

func usage() {
    fmt.Println("Usage: proxy <port-number>")
    os.Exit(EXIT_FAILURE)
}

func main() {
    if len(os.Args) != 2 {
        usage()
    }
}

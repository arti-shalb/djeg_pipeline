package main

import (
    "bufio"
    "os"
    "fmt"
)

// To terminate the daemon use:
//  kill `cat pid`

func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    fmt.Printf(text)
}


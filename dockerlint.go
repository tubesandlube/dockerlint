package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        fmt.Println("error")
        panic(e)
    }
}

func main() {
    var buffer bytes.Buffer

    fmt.Println(len(os.Args), os.Args)

    file, err := os.Open("Dockerfile")
    check(err)

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        buffer.WriteString(scanner.Text())

        // strip whitespace
        // check end of line for '\', to determine breaks between commands
        partialCommand := strings.HasSuffix(strings.TrimSpace(buffer.String()), "\\")
        if !partialCommand {
            fmt.Println(buffer.String())
            buffer.Reset()
        }
    }

    Rules()


    file, err = os.Create("Dockerfile")
    check(err)

    w := bufio.NewWriter(file)
    n4, err := w.WriteString("buffered\n")
    fmt.Println("wrote %d bytes\n", n4)

    w.Flush()
}

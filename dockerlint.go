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
    var dockerfile string
    var writeString string

    fmt.Println(len(os.Args), os.Args)
    if len(os.Args) <= 1 {
        dockerfile = "Dockerfile"
    } else {
        dockerfile = os.Args[1]
    }

    readFile, err := os.Open(dockerfile)
    check(err)
    defer readFile.Close()

    scanner := bufio.NewScanner(readFile)
    for scanner.Scan() {
        buffer.WriteString(scanner.Text())

        // strip whitespace
        // check end of line for '\', to determine breaks between commands
        partialCommand := strings.HasSuffix(strings.TrimSpace(buffer.String()), "\\")
        if !partialCommand {
            if len(buffer.String()) > 0 {
                 writeString += Rules(buffer.String())
            }
            writeString += "\n"
            buffer.Reset()
        }
    }

    writeFile, err := os.Create(dockerfile)
    check(err)

    w := bufio.NewWriter(writeFile)
    _, err = w.WriteString(writeString)
    check(err)

    w.Flush()
}

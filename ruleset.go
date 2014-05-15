package main

import (
    "fmt"
    "strings"
)

func validate() {
    fmt.Println("validate the dockerfile is syntactically correct")
}

func capitalize(command string) string {
    fmt.Println("capitalize all of the commands")
    s := strings.SplitN(command, " ", 2)
    s[0] = strings.ToUpper(s[0])
    command = strings.Join(s, " ")
    fmt.Println(command)
    return command
}

func checkFirstCommand(command string) bool {
    fmt.Println("first command must be FROM")
    // TODO
    return true
}

func checkFromTag(command string) bool {
    fmt.Println("FROM command should have a tag")
    // TODO
    return true
}

func Rules(command string, line int) string {
    // Run against ruleset unless it's a comment
    if strings.HasPrefix(command, "#") {
        fmt.Println("Line is a comment")
    } else {
        validate()
        command = capitalize(command)
        if line == 1 {
            v := checkFirstCommand(command)
            fmt.Println(v)
            v = checkFromTag(command)
            fmt.Println(v)
        }
    }
    fmt.Println(line)
    return command
}

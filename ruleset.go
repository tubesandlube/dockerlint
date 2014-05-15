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
    command = strings.ToUpper(command)
    if strings.HasPrefix(command, "FROM ") {
        return true
    }
    return false
}

func checkFromTag(command string) bool {
    fmt.Println("FROM command should have a tag")
    if strings.Contains(command, ":") {
        return true
    }
    return false
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
        // check for too many layers
        if line == 128 {
            fmt.Println("Too many layers were found, reduce the number of commands")
        }
    }
    fmt.Println(line)
    return command
}

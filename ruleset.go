package main

import (
    "fmt"
    "regexp"
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


func Rules(command string) string {
    r, err := regexp.Compile(command)

    if err != nil {
        fmt.Println("There is a problem with your regexp.\n")
        return ""
    }

    // Will print 'Match'
    if r.MatchString("Hello Regular Expression.") == true {
        fmt.Println("Match ")
    } else {
        fmt.Println("No match ")
    }

    // Run against ruleset unless it's a comment
    if strings.HasPrefix(command, "#") {
        fmt.Println("Line is a comment")
    } else {
        validate()
        command = capitalize(command)
    }
    return command
}

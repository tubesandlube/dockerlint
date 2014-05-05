package main

import (
    "fmt"
    "regexp"
)

func validate() {
    fmt.Println("validate the dockerfile is syntactically correct")
}

func capitalize() {
    fmt.Println("capitalize all of the commands")
}


func Rules(command string) {
    r, err := regexp.Compile(command)

    if err != nil {
        fmt.Println("There is a problem with your regexp.\n")
        return
    }

    // Will print 'Match'
    if r.MatchString("Hello Regular Expression.") == true {
        fmt.Println("Match ")
    } else {
        fmt.Println("No match ")
    }
    validate()
    capitalize()
}

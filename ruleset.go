package main

import (
    "fmt"
    "regexp"
)

func validate() {
    fmt.Printf("validate the dockerfile is syntactically correct")
}

func capitalize() {
    fmt.Printf("capitalize all of the commands")
}


func Rules() {
    r, err := regexp.Compile(`Hello`)

    if err != nil {
        fmt.Printf("There is a problem with your regexp.\n")
        return
    }

    // Will print 'Match'
    if r.MatchString("Hello Regular Expression.") == true {
        fmt.Printf("Match ")
    } else {
        fmt.Printf("No match ")
    }
    validate()
    capitalize()
}

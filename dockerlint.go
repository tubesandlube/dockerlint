package main

import (
    "fmt"
    "io/ioutil"
    "os"
    //"strings"
)

func check(e error) {
    if e != nil {
        fmt.Printf("error")
        panic(e)
    }
}

func main() {
    fmt.Println(len(os.Args), os.Args)

    dat, err := ioutil.ReadFile("Dockerfile")
    check(err)
    fmt.Print(string(dat))


    // join dat into a string
    // split string into slices by newlines

    // split dat into commands
    // check end of line for '\', to determine breaks between commands
    for _,element := range dat {
        fmt.Print(string(element))
        fmt.Print("foo \n")
    }

    rules()

    // strip whitespace
    //endsWith := strings.HasSuffix("suffix", "\")

    d1 := []byte("hello\ngo\n")
    err = ioutil.WriteFile("Dockerfile", d1, 0644)
    check(err)
}

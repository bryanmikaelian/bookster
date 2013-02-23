package main

import (
    "bookster/collector"
    "fmt"
    "github.com/bitly/go-simplejson"
)

func main() {
    fmt.Println("Looking for books...")

    for i := 1; i < 10; i++ {
        r := make(chan *simplejson.Json)
        go collector.FindBooks("go", i, r)
        data := <- r
        fmt.Println(data)
    }
}



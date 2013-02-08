package collector

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/bitly/go-simplejson"
    "bookster/book"
)

const API_URL = "https://www.googleapis.com/books/v1/volumes?q="

func FindBooks(isbn string){
    data, size := fetch(isbn)

    for i := 0; i < size; i++ {
        c := make(chan book.Book)
        go book.Build(data.Get("items").GetIndex(i), c)
        google_book := <-c
        fmt.Println(google_book)
    }
}

func fetch(title string) (books *simplejson.Json, size int) {
    fmt.Println("Looking for books with an title of " + title)
    resp, err := http.Get(API_URL +  title + "&startIndex=1&maxResults=40")
    defer resp.Body.Close()

    if err != nil {
        panic(err)
    }

    body, err := ioutil.ReadAll(resp.Body)

    results, err := simplejson.NewJson(body)
    items, err := results.Get("items").Array()
    size = len(items)

    return results, size
}

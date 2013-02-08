package book

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/bitly/go-simplejson"
)

const API_URL = "https://www.googleapis.com/books/v1/volumes?q="

func GetBook(isbn string){
    fetch(isbn)
}

func fetch(isbn string) (books *simplejson.Json) {
    fmt.Println("Fetching book with ISBN of " + isbn)
    resp, err := http.Get(API_URL + "ISBN:" + isbn)
    defer resp.Body.Close()

    if err != nil {
        panic(err)
    }

    body, err := ioutil.ReadAll(resp.Body)

    results, err := simplejson.NewJson(body)

    return results
}

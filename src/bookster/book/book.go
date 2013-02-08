package book

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Book struct {
    Kind string 
    Totalitems string
}

const API_URL = "https://www.googleapis.com/books/v1/volumes?q="

func GetBook(isbn string){
    fetch(isbn)
}

func fetch(isbn string) {
    fmt.Println("Fetching book with ISBN of " + isbn)
    resp, err := http.Get(API_URL + "ISBN:" + isbn)
    defer resp.Body.Close()

    if err != nil {
    }

    body, err := ioutil.ReadAll(resp.Body)
    var r interface{}
    results := []byte(string(body))
    json.Unmarshal(results, &r)
    books := r.(map[string]interface{})
    fmt.Println(books["totalItems"])

}

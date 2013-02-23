package collector

import (
    "net/http"
    "io/ioutil"
    "github.com/bitly/go-simplejson"
    "bookster/book"
    "strconv"
)

const API_URL = "https://www.googleapis.com/books/v1/volumes?q="

func FindBooks(title string, pageNumber int, r chan *simplejson.Json){
    data, size := fetch(title, pageNumber)

    for i := 0; i < size; i++ {
        c := make(chan book.Book)
        go book.Build(data.Get("items").GetIndex(i), c)
    }
    r <- data
}

func fetch(title string, pageNumber int) (books *simplejson.Json, size int) {
    page := 10 * (pageNumber - 1)
    resp, err := http.Get(API_URL +  title + "&startIndex=" + strconv.Itoa(page) + "&maxResults=20")
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

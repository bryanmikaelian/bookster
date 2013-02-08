package book

import (
    "github.com/bitly/go-simplejson"
)


type Book struct {
    GoogleBookId string
    title string
    ISBN10 string
    ISBN13 string
}

func Build(j *simplejson.Json, c chan Book) {
    id, err := j.Get("id").String()
    title, err := j.Get("volumeInfo").Get("title").String()
    isbn10, err := j.Get("volumeInfo").Get("industryIdentifiers").GetIndex(0).Get("identifier").String()
    isbn13, err := j.Get("volumeInfo").Get("industryIdentifiers").GetIndex(1).Get("identifier").String()

    if err != nil {
    }

    book := Book{id, title, isbn10, isbn13}
    c <- book
}

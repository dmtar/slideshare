go-slideshare
=============

API wrapper for SlideShare written in Go.
Everything is in a very initial stage...preparation, arrangement and stuff.

Installation
------------

When everything is ready, package installation should be done via command bellow: 

`go get github.com/dishev/go-slideshare`

Examples
--------

    package main

    import (
        "fmt"
        "github.com/dishev/go-slideshare"
    )

    func main() {
        service := slideshare.Service{"API_KEY", "SHARED_SECRET"}
        slideshow, err := service.GetSlideshow(29551397)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("ID: ", slideshow.ID)
        fmt.Println("Title: ", slideshow.Title)
        fmt.Println("Description: ", slideshow.Description)
        fmt.Println("Username: ", slideshow.Username)
    }

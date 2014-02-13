slideshare
=============

API wrapper for SlideShare written in Go.
Everything is in a very initial stage...preparation, arrangement and stuff.

Installation
------------

When everything is ready, package installation should be done via command bellow: 

`go get github.com/dmtar/slideshare`

Methods
--------

- add_favorite
- check_favorite
- delete_slideshow
- edit_slideshow
- get_slideshow
- get_slideshows_by_tag
- get_slideshows_by_user
- get_user_contacts
- get_user_favorites
- search_slideshows
- upload_slideshow

- get_slideshow_by_group (500 Internal Server Error)
- get_user_groups (after redirecting returns "This is probably a temporary error, please check back in 5-10 minutes.")

Methods below require a PRO account to be used and tested, they are not implemented. There is no info about the XML structure on the slideshare website.

- get_user_campaign_leads
- get_user_campaigns
- get_user_leads

Examples
--------

    package main

    import (
        "fmt"
        "github.com/dishev/slideshare"
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

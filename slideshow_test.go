package slideshare

import (
	"testing"
)

const (
	ApiKey       = YOUR_API_KEY
	SharedSecret = YOUR_SHARED_SECRET
)

// Test basic getting of slideshow
func TestGetSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30975136)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30975136 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
}

// Test getting a slideshow with detailed info
func TestGetSlideshowDetailed(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30975136, true)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30975136 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
}
func TestGetSlideshowsByTag(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.GetSlideshowsByTag("db", false, 2)
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.TagName != "db" {
		t.Fail()
	}
	if slideshows.Slideshows[0].Title != "Memcache basics on google app engine" {
		t.Fail()
	}
	if slideshows.Slideshows[0].ID != 30796493 {
		t.Fail()
	}
	if slideshows.Slideshows[0].Format != "pdf" {
		t.Fail()
	}
}
func TestGetSlideshowsByUser(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.GetSlideshowsByUser("ddishev", true, 10)
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.UserName != "ddishev" {
		t.Fail()
	}
	if slideshows.Slideshows[0].ID != 30976468 {
		t.Fail()
	}
}

func TestEditSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isEdited := service.EditSlideshow("ddishev", YOUR_PASSWORD, 30975136, "Databeses")
	if !isEdited {
		t.Fail()
	}
}

func TestDeleteSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isDeleted := service.DeleteSlideshow("ddishev", YOUR_PASSWORD, 30976468)
	if !isDeleted {
		t.Fail()
	}
}

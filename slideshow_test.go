package slideshare

import (
	"testing"
)

const (
	ApiKey       =  //Your API key
	SharedSecret =  // Your shared secret
	Password     = 
	Username     =  
)

// Test basic getting of slideshow
func TestGetSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30130693)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30130693 {
		t.Fail()
	}
	if slideshow.Username != "kalinazdravkova" {
		t.Fail()
	}
}

// Test getting a slideshow with detailed info
func TestGetSlideshowDetailed(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30130693, true)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30130693 {
		t.Fail()
	}
	if slideshow.Username != "kalinazdravkova" {
		t.Fail()
	}
	if slideshow.UserID != 41191689 {
		t.Fail()
	}
	if slideshow.Format != "pptx" {
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
	slideshows, err := service.GetSlideshowsByUser("kalinazdravkova", true, 10)
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.UserName != "kalinazdravkova" {
		t.Fail()
	}
	if slideshows.Slideshows[0].ID != 30130693 {
		t.Fail()
	}
}

func TestEditSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isEdited := service.EditSlideshow(Username, Password, 30975136, "TestName")
	if !isEdited {
		t.Fail()
	}
}

/*
func TestDeleteSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isDeleted := service.DeleteSlideshow(Username, Password, 30976468)
	if !isDeleted {
		t.Fail()
	}
}
*/

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
	slideshow, err := service.GetSlideshow(29905515)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 29905515 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
}

// Test getting a slideshow with detailed info
func TestGetSlideshowDetailed(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(29905515, true)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 29905515 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
	if slideshow.NumSlides != 65 {
		t.Fail()
	}
	if slideshow.Tags[0] != "fmi" {
		t.Fail()
	}
	if slideshow.RelatedSlideshows[0] != 9551082 {
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
	if slideshows.Slideshows[0].ID != 29905515 {
		t.Fail()
	}
	if slideshows.Slideshows[0].Format != "ppt" {
		t.Fail()
	}
}
func TestSearchSlideshows(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	_, err := service.SearchSlideshows("db", false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEditSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isEdited := service.EditSlideshow("ddishev", PASSWORD, 29905515, "EditedTitleLastOne")
	if !isEdited {
		t.Fail()
	}
}

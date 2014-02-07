package slideshare

import (
	"testing"
)

const (
	ApiKey       = "5Pl6RFlI"
	SharedSecret = "X1lMfjPo"
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
	if slideshow.Title != "Databases some other lecture" {
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
	if slideshow.Title != "Databases some other lecture" {
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

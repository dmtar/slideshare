package slideshare

import (
	"testing"
)

func TestCheckFavorite(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.CheckFavorite("ddishev", YOUR_PASSWORD, 29430493)
	if err != nil {
		t.Fatal(err)
	}

	if slideshow.Favorited != true {
		t.Fail()
	}
	if slideshow.SlideshowID != 29430493 {
		t.Fail()
	}
	if slideshow.UserID != 59394728 {
		t.Fail()
	}
}

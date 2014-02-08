package slideshare

import (
	"testing"
)

func TestCheckFavorite(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.CheckFavorite(Username, Password, 30130693)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.Favorited != false {
		t.Fail()
	}
}

func TestAddFavorite(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isFavorited := service.AddFavorite(Username, Password, 9662505)
	if isFavorited != true {
		t.Fail()
	}
}

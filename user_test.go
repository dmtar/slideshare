package slideshare

import (
	"testing"
)

func TestGetUserFavorites(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.GetUserFavorites(Username)
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.Values[0].SlideshowID != 9662505 {
		t.Fail()
	}
}

/*
func TestGetUserContacts(t *testing.T)

func TestGetUserGroups(t *testing.T)

func TestGetUserTags(t *testing.T)
*/

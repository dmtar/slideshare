package slideshare

import (
	"testing"
)

func TestGetUserFavorites(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	favorites, err := service.GetUserFavorites(Username)
	if err != nil {
		t.Fatal(err)
	}
	if favorites.Values[0].SlideshowID != 9662505 {
		t.Fail()
	}
}

func TestGetUserContacts(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	contacts, err := service.GetUserContacts(Username)
	if err != nil {
		t.Fatal(err)
	}
	if contacts.Values[0].Username != "itseugene" {
		t.Fail()
	}
}

/*
func TestGetUserGroups(t *testing.T)

func TestGetUserTags(t *testing.T)
*/

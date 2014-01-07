package slideshare

import (
	"fmt"
)

type UserFavorites struct {
	Favorites []UserFavorite `xml:"favorites"`
}
type UserFavorite struct {
	XMLName     xml.Name `xml:"favorite"`
	SlideshowID uint64   `xml:"slideshow_id"`
	TagText     string   `xml:"tag_text"`
}
type Contacts struct {
	Contacts []Contact `xml:"Contacts"`
}
type Contact struct {
	XMLName       xml.Name `xml:"Contact"`
	Username      string   `xml:"Username"`
	NumSlideshows uint32   `xml:"NumSlideshows"`
	NumComments   uint32   `xml:"NumComments"`
}
type Groups struct {
	Groups []Group `xml:"group"`
}
type Group struct {
	Name          string `xml:"name"`
	NumPosts      uint32 `xml:"numposts"`
	NumSlideshows uint32 `xml:"numslideshows"`
	NumMembers    uint32 `xml:"nummembers"`
	Created       string `xml:"created"`
	QueryName     string `xml:"queryname"`
	Url           string `xml:"url"`
}

func getUserFavorites(username string) (UserFavorites, error)
func getUserContacts(username string, limit uint32) (Contacts, error)
func getUserGroups(username string) (Groups, error)
func getUserTags(username string, password string) (Tags, error)

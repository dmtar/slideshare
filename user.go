package slideshare

// UserFavorites keeps an array with favorited slideshows by a user.
type UserFavorites struct {
	Favorites []UserFavorite `xml:"favorites"`
}

// UserFavorite keeps information about a favorited slideshow by a user.
type UserFavorite struct {
	XMLName     xml.Name `xml:"favorite"`
	SlideshowID uint64   `xml:"slideshow_id"`
	TagText     string   `xml:"tag_text"`
}

// Contacts holds an array with all Contacts for a given user.
type Contacts struct {
	Contacts []Contact `xml:"Contacts"`
}

// Contact holds info like username, number of uploaded slideshows
// number of comments of a Contact for a given user.
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

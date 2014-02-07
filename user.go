package slideshare

// UserFavorites keeps an array with favorited slideshows by a user.
type UserFavorites struct {
	Favorites []UserFavorite `xml:"favorites"`
}

// UserFavorite keeps information about a favorited slideshow by a user.
type UserFavorite struct {
	SlideshowID uint64 `xml:"slideshow_id"`
	TagText     string `xml:"tag_text"`
}

// Contacts holds an array with all Contacts for a given user.
type Contacts struct {
	Contacts []Contact `xml:"Contacts"`
}

// Contact holds info like username, number of uploaded slideshows
// number of comments of a Contact for a given user.
type Contact struct {
	Username      string `xml:"Username"`
	NumSlideshows uint32 `xml:"NumSlideshows"`
	NumComments   uint32 `xml:"NumComments"`
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

/*
// Returns user favorites
// username_for required, username of user whose favorites are being requested.
func getUserFavorites(username_for string) (UserFavorites, error)

// Returns user contacts
// username_for required, username of user whose contacts are being requested
func getUserContacts(username_for string, limit uint32) (Contacts, error)

// Returns user groups
// username_for required, username of user whose groups are being requested
func getUserGroups(username_for string) (Groups, error)

// Returns user tags
// username required, username of user whose tags are being requested
// password required, password of user whose tags are being requested
func getUserTags(username string, password string) (Tags, error)
*/

package slideshare

type UserFavorites struct {
	Values []UserFavorite `xml:"favorite"`
}
type UserFavorite struct {
	SlideshowID uint64 `xml:"slideshow_id"`
	TagText     string `xml:"tag_text"`
}
type UserContacts struct {
	Values []UserContact `xml:"Contact"`
}
type UserContact struct {
	Username      string `xml:"Username"`
	NumSlideshows uint32 `xml:"NumSlideshows"`
	NumComments   uint32 `xml:"NumComments"`
}

type Groups struct {
	Values []Group `xml:"group"`
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
type Tags struct {
	Names []string `xml:"Tag"`
}

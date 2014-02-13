package slideshare

// Favorite type keeps the information about favorited slideshow by the user
type Favorite struct {
	SlideshowID int    `xml:"SlideshowID"`
	UserID      uint64 `xml:"User"`
	Favorited   bool   `xml:"Favorited"`
}

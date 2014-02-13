package slideshare

type Favorite struct {
	SlideshowID int    `xml:"SlideshowID"`
	UserID      uint64 `xml:"User"`
	Favorited   bool   `xml:"Favorited"`
}

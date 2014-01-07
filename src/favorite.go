package slideshare

import (
	"fmt"
)

type Favorite struct {
	XMLName     xml.Name `xml:"Slideshow"`
	SlideshowID uint64   `xml:"SlideshowID"`
	User        string   `xml:"User"`
	Favorited   bool     `xml:"Favorited"`
}

func (s *Service) checkFavorite(username string, password string, slideshow_id string) (Favorite, error) {
}
func (s *Service) addFavorite(username string, password string, slideshow_id string) (Favorite, error) {
}

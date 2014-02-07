package slideshare

type Favorite struct {
	SlideshowID uint64 `xml:"SlideshowID"`
	User        string `xml:"User"`
	Favorited   bool   `xml:"Favorited"`
}

/*
// Check user favorites
// username required, username of the  requesting user
// password required, password of the  requesting user
// slideshow_id required, Slideshow which is being checked
func (s *Service) checkFavorite(username string, password string, slideshow_id string) (Favorite, error) {
}

// Favorite a slideshow (identified by slideshow_id)
// username required, username of the  requesting user
// password required, password of the  requesting user
// slideshow_id required, the slideshow to be favorited
func (s *Service) addFavorite(username string, password string, slideshow_id string) (Favorite, error) {
}
*/

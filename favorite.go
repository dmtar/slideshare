package slideshare

import (
	"encoding/xml"

	"io/ioutil"
	"net/http"

	"strconv"
)

type Favorite struct {
	SlideshowID uint64 `xml:"SlideshowID"`
	UserID      uint64 `xml:"User"`
	Favorited   bool   `xml:"Favorited"`
}

// Check user favorites
// username required, username of the  requesting user
// password required, password of the  requesting user
// slideshow_id required, Slideshow which is being checked
func (s *Service) CheckFavorite(username string, password string, slideshow_id int) (Favorite, error) {
	args := make(map[string]string)
	args["username"] = username
	args["password"] = password
	args["slideshow_id"] = strconv.Itoa(slideshow_id)
	url := s.generateUrl("check_favorite", args)
	resp, err := http.Get(url)
	if err != nil {
		return Favorite{}, err
	}
	slideshow := Favorite{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
	}
	return slideshow, err
}

/*
// Favorite a slideshow (identified by slideshow_id)
// username required, username of the  requesting user
// password required, password of the  requesting user
// slideshow_id required, the slideshow to be favorited
func (s *Service) AddFavorite(username string, password string, slideshow_id string) (Favorite, error) {
}
*/

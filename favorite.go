package slideshare

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Check user favorite needs username and password of the requesting user
// and id of the slideshow which is being checked, whether is favorited by the user.
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

// Add favorite favorites a slideshow. Needs username and password of the requesting user
// and id of the slideshow to be favorited.
func (s *Service) AddFavorite(username string, password string, slideshow_id int) bool {
	args := make(map[string]string)
	args["username"] = username
	args["password"] = password
	args["slideshow_id"] = strconv.Itoa(slideshow_id)
	url := s.generateUrl("add_favorite", args)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	slideshow := Favorite{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
		if slideshow.SlideshowID == slideshow_id {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

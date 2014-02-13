package slideshare

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetSlideshow returns information about a slideshow
// needs id of the slideshow to be fetched, and detailed flag.
func (s *Service) GetSlideshow(id int, detailed bool) (Slideshow, error) {
	args := make(map[string]string)
	args["slideshow_id"] = strconv.Itoa(id)
	args["detailed"] = Btoa(detailed)
	url := s.generateUrl("get_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return Slideshow{}, err
	}
	slideshow := Slideshow{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
	}
	return slideshow, err
}

// GetSlideshowByTag need tag name and detailed flag
// limit and offset are optional if they are used, should be passed in that way limit first then offset.
func (s *Service) GetSlideshowsByTag(tag string, detailed bool, limitOffset ...int) (SlideshowsByTag, error) {
	args := make(map[string]string)
	if limitOffset != nil {
		switch len(limitOffset) {
		case 1:
			args["limit"] = strconv.Itoa(limitOffset[0])
			break
		case 2:
			args["limit"] = strconv.Itoa(limitOffset[0])
			args["offset"] = strconv.Itoa(limitOffset[1])
			break
		default:
		}
	}
	args["tag"] = tag
	args["detailed"] = Btoa(detailed)
	url := s.generateUrl("get_slideshows_by_tag", args)
	resp, err := http.Get(url)
	if err != nil {
		return SlideshowsByTag{}, err
	}
	slideshows := SlideshowsByTag{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshows)
	}
	return slideshows, err
}

// GetSlideshowByUser needs username and detailed flag arguments, others like
// limit and offset are optional.Should be passed in that way limit first then offset.
func (s *Service) GetSlideshowsByUser(username string, detailed bool, limitOffset ...int) (SlideshowsByUser, error) {
	args := make(map[string]string)
	if limitOffset != nil {
		switch len(limitOffset) {
		case 1:
			args["limit"] = strconv.Itoa(limitOffset[0])
			break
		case 2:
			args["limit"] = strconv.Itoa(limitOffset[0])
			args["offset"] = strconv.Itoa(limitOffset[1])
			break
		default:
		}
	}
	args["username_for"] = username
	args["detailed"] = Btoa(detailed)
	url := s.generateUrl("get_slideshows_by_user", args)
	resp, err := http.Get(url)
	if err != nil {
		return SlideshowsByUser{}, err
	}
	slideshows := SlideshowsByUser{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshows)
	}
	return slideshows, err
}

// Search for a slideshow needs query string and detailed flag arguments
// Optional are page default is 1, items per page default is 12,
// sort type default is relevance, and upload date default is any
// If you want to change the default values of additional parameters, you
// should pass them like that: page,items_per_page,sort,upload_date and values casted to string.
func (s *Service) SearchSlideshows(queryString string, detailed bool, additionalParams ...string) (SlideshowsSearch, error) {
	args := make(map[string]string)
	if additionalParams != nil {
		switch len(additionalParams) {
		case 1:
			args["page"] = additionalParams[0]
			args["items_per_page"] = "12"
			args["sort"] = "relevance"
			args["upload_date"] = "any"
			break
		case 2:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["sort"] = "relevance"
			args["upload_date"] = "any"
			break
		case 3:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["sort"] = additionalParams[2]
			args["upload_date"] = "any"
			break
		case 4:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["sort"] = additionalParams[2]
			args["upload_date"] = additionalParams[3]
			break
		default:
			args["page"] = "1"
			args["items_per_page"] = "12"
			args["sort"] = "relevance"
			args["upload_date"] = "any"
			break
		}
	} else {
		args["page"] = "1"
		args["items_per_page"] = "12"
		args["sort"] = "relevance"
		args["upload_date"] = "any"
	}
	args["q"] = queryString
	args["detailed"] = Btoa(detailed)
	url := s.generateUrl("search_slideshows", args)
	resp, err := http.Get(url)
	if err != nil {
		return SlideshowsSearch{}, err
	}
	slideshows := SlideshowsSearch{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshows)
	}
	return slideshows, err
}

// Edit a Slideshow needs username and password of the owner of the requesting user and slideshow id
// other parameters are new title, new tags and make the slideshow private,
// generate secret url, allow embeds and share with contacts are optinal.You have to pass their values
// like strigs.
func (s *Service) EditSlideshow(username string, password string, slideshowID int, additionalParams ...string) bool {
	args := make(map[string]string)
	if additionalParams != nil {
		switch len(additionalParams) {
		case 1:
			args["slideshow_title"] = additionalParams[0]
			break
		case 2:
			args["slideshow_title"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			break
		case 3:
			args["slideshow_title"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			args["make_slideshow_private"] = additionalParams[2]
			break
		case 4:
			args["slideshow_title"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			args["make_slideshow_private"] = additionalParams[2]
			args["generate_secret_url"] = additionalParams[3]
			break
		case 5:
			args["slideshow_title"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			args["make_slideshow_private"] = additionalParams[2]
			args["generate_secret_url"] = additionalParams[3]
			args["allow_ebeds"] = additionalParams[4]
			break
		case 6:
			args["slideshow_title"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			args["make_slideshow_private"] = additionalParams[2]
			args["generate_secret_url"] = additionalParams[3]
			args["allow_ebeds"] = additionalParams[4]
			args["share_with_contacts"] = additionalParams[5]
			break
		default:
		}
	}
	args["username"] = username
	args["password"] = password
	args["slideshow_id"] = strconv.Itoa(slideshowID)
	url := s.generateUrl("edit_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	slideshow := EditedSlideshow{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
		if slideshow.EditedID == slideshowID {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Delete a slideshow needs username and password of the requesting user and slideshow ID.
func (s *Service) DeleteSlideshow(username string, password string, slideshowID int) bool {
	args := make(map[string]string)
	args["username"] = username
	args["password"] = password
	args["slideshow_id"] = strconv.Itoa(slideshowID)
	url := s.generateUrl("delete_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	slideshow := DeletedSlideshow{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
		if slideshow.DeletedID == slideshowID {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// This method requires extra permissions.
// If you want to upload a file using SlideShare API, please send an email to api@slideshare.com with your developer account username describing the use case.
// The method requires username and password of the requesting user, title and upload url string containing an url pointing to the power point file:
// ex: http://domain.tld/directory/my_power_point.ppt
// The following urls are also acceptable
// http://www.domain.tld/directory/file.ppt
// http://www.domain.tld/directory/file.cgi?filename=file.ppt
// Optinal parameters are description of the slideshow and slideshow tags.
func (s *Service) UploadSlideshow(username string, password string, uploadURL string, slideshowTitle string, additionalParams ...string) (int, bool) {
	args := make(map[string]string)
	if additionalParams != nil {
		switch len(additionalParams) {
		case 1:
			args["slideshow_description"] = additionalParams[0]
			break
		case 2:
			args["slideshow_description"] = additionalParams[0]
			args["slideshow_tags"] = additionalParams[1]
			break
		default:
		}
	}
	args["username"] = username
	args["password"] = password
	args["upload_url"] = uploadURL
	args["slideshow_title"] = slideshowTitle
	url := s.generateUrl("upload_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return 0, false
	}
	slideshow := UploadedSlideshow{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &slideshow)
		if slideshow.UploadedID == 0 {
			return 0, false
		}
		return slideshow.UploadedID, true
	} else {
		return 0, false
	}
}

/*
func (s *Service) GetSlideshowsByGroup(groupName string, detailed bool) (Slideshows, error) {

}
*/

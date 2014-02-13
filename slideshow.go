package slideshare

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetSlideshow returns information about a slideshow, parameters:
// id int which holds the slideshow id, required.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
// return: Slideshow instance.
func (s *Service) GetSlideshow(id int, detailed ...bool) (Slideshow, error) {
	args := make(map[string]string)
	var details bool
	if detailed == nil {
		details = false
	} else {
		details = detailed[0]
	}
	args["slideshow_id"] = strconv.Itoa(id)
	args["detailed"] = Btoa(details)
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

// GetSlideshowByUser returns a Slideshows object:
// tag string required, tag name.
// detailed bool required, whether or not to include additional information for the slideshows.
// limitOffset int optinal, first int for limit, second for offset, others are ignored.
// returns SlideshowsByTag
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
	} else {
		args["limit"] = "10"
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

/*
func (s *Service) GetSlideshowsByGroup(groupName string, detailed bool) (Slideshows, error) {

}
*/

// GetSlideshowByUser returns a Slideshows object:
// username string required, username of the owner of slideshows.
// detailed bool required, whether or not to include additional information for the slideshows.
// limitOffset int optinal, first int for limit, second for offset, others are ignored.
// returns SlideshowsByUser
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
	} else {
		args["limit"] = "10"
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

// Search for a slideshow:
// queryString required, search keyword
// page optional, the page number of the results (works in conjunction with items_per_page), default is 1
// items_per_page optional, number of results to return per page, default is 12
// lang optional, Language of slideshows (default is English, 'en') ('**':All,'es':Spanish,'pt':Portuguese,'fr':French,'it':Italian,'nl':Dutch, 'de':German,'zh':Chinese,'ja':Japanese,'ko':Korean,'ro':Romanian, '!!':Other)
// sort optional, Sort order, default is relevance.
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
		}
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

// Edit a Slideshow, you can change only the title, tags and privacy of the slideshow.
// username required, owner username of the slideshow which is being edited
// password required, owner password of the slideshow which is being edited
// slideshowID required, id of slideshow which is being edited.
// slideshow_title optional, Title of the slideshow
// slideshow_ tags optional, Comma separated list of tags
// make_slideshow_private optional, Should be Y if you want to make the slideshow private.
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

// Delete a slideshow all parameters are required
// username required,owner username of the slideshow which is being deleted
// password required,owner password of the slideshow which is being deleted
// slideshowID required,Id of slideshow which is being deleted
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

// Upload a slideshow
// username required, username of the  requesting user
// password required, password of the  requesting user
// uploadURL required, string containing an url pointing to the power point file: ex: http://domain.tld/directory/my_power_point.ppt
// slideshow_title required, Title of the slideshow
// slideshow_description optional, slideshow description
// slideshow_tags optional, Comma separated list of tags.
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

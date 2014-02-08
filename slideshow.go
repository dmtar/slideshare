package slideshare

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type EditedSlideshow struct {
	EditedID int `xml:"SlideShowID"`
}

// SlideshowsByTag struct keeps the Tag we are looking for
// Count of all found slideshow and array with them.
type SlideshowsByTag struct {
	TagName    string      `xml:"Name"`
	Count      uint32      `xml:"Count"`
	Slideshows []Slideshow `xml:"Slideshow"`
}

// SlideshowsByUser struct keeps the username we are looking for
// Count of all found slideshow and array with them.
type SlideshowsByUser struct {
	UserName   string      `xml:"Name"`
	Count      uint32      `xml:"Count"`
	Slideshows []Slideshow `xml:"Slideshow"`
}

// SlideshowsByUser struct keeps the query string.
// offset of the result,results per page and total results
// and array with slideshows.
type SlideshowsSearch struct {
	QueryString  string      `xml:"Meta>Query"`
	ResultOffset uint16      `xml:"Meta>ResultOffset"`
	NumResults   uint32      `xml:"Meta>NumResults"`
	TotalResults uint64      `xml:"Meta>TotalResults"`
	Slideshows   []Slideshow `xml:"Slideshow"`
}

// Type which holds deleted slideshow ID.
type SlideShowDeleteted struct {
	XMLName xml.Name `xml:"SlideShowDeleted"`
	ID      uint64   `xml:SlideShowDeleted>SlideshowID`
}

// Slideshow type, which holds all the information about a slideshow
// properties bellow InContest are detailed, they will have reliabe information
// if detailed flag is set to true.
type Slideshow struct {
	XMLName           xml.Name `xml:"Slideshow"`
	ID                uint64   `xml:"ID"`
	Title             string   `xml:"Title"`
	Description       string   `xml:"Description"`
	Username          string   `xml:"Username"`
	Status            uint8    `xml:"Status"`
	Url               string   `xml:"URL"`
	ThumbnailUrl      string   `xml:"ThumbnailURL"`
	ThumbnailSize     string   `xml:"ThumbnailSize"`
	ThumbnailSmallUrl string   `xml:"ThumbnailSmallURL"`
	Embed             string   `xml:"Embed"`
	Created           string   `xml:"Created"`
	Updated           string   `xml:"Updated"`
	Language          string   `xml:"Language"`
	Format            string   `xml:"Format"`
	Download          bool     `xml:"Download"`
	DownloadUrl       string   `xml:"DownloadUrl"`
	SlideshowType     uint8    `xml:"SlideshowType"`
	InContest         bool     `xml:"InContest"`
	UserID            uint64   `xml:"UserID"`
	PPTLocation       string   `xml:"PPTLocation"`
	StrippedTitle     string   `xml:"StrippedTitle"`
	Tags              []string `xml:"Tags>Tag"`
	Audio             bool     `xml:"Audio"`
	NumDownloads      uint32   `xml:"NumDownloads"`
	NumViews          uint32   `xml:"NumViews"`
	NumComments       uint32   `xml:"NumComments"`
	NumFavorites      uint32   `xml:"NumFavorites"`
	NumSlides         uint16   `xml:"NumSlides"`
	RelatedSlideshows []uint64 `xml:"RelatedSlideshows>RelatedSlideshowID"`
	PrivacyLevel      bool     `xml:"PrivacyLevel"`
	FlagVisible       bool     `xml:"FlagVisible"`
	ShowOnSS          bool     `xml:"ShowOnSS"`
	SecretURL         bool     `xml:"SecretUrl"`
	AllowEmbed        bool     `xml:"AllowEmbed"`
	ShareWithContacs  bool     `xml:"ShareWithContacs"`
}

var apiUrl = "https://www.slideshare.net/api/2"

func (s *Service) generateUrl(apiMethod string, arguments map[string]string) string {
	values := url.Values{}
	for key, value := range arguments {
		values.Set(key, value)
	}
	values.Set("api_key", s.ApiKey)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	values.Set("ts", timestamp)
	hash := sha1.New()
	io.WriteString(hash, s.SharedSecret+timestamp)
	values.Set("hash", fmt.Sprintf("%x", hash.Sum(nil)))
	return apiUrl + "/" + apiMethod + "?" + values.Encode()
}
func Btoa(input bool) string {
	if input {
		return "1"
	} else {
		return "0"
	}
}

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
			break
		case 2:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			break
		case 3:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["lang"] = additionalParams[2]
			break
		case 4:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["lang"] = additionalParams[2]
			args["sort"] = additionalParams[3]
			break
		case 5:
			args["page"] = additionalParams[0]
			args["items_per_page"] = additionalParams[1]
			args["lang"] = additionalParams[2]
			args["sort"] = additionalParams[3]
			args["upload_date"] = additionalParams[4]
			break
		default:
		}
	} else {
		args["page"] = "1"
		args["items_per_page"] = "16"
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

/*
// Delete a slideshow all parameters are required
// username required,owner username of the slideshow which is being deleted
// password required,owner password of the slideshow which is being deleted
// slideshowID required,Id of slideshow which is being deleted
func (s *Service) DeleteSlideshow(username string, password string, slideshowID string) bool {}

// Upload a slideshow
// username required, username of the  requesting user
// password required, password of the  requesting user
// uploadURL required, string containing an url pointing to the power point file: ex: http://domain.tld/directory/my_power_point.ppt
// slideshow_title required, Title of the slideshow
// slideshow_description optional, slideshow description
// slideshow_tags optional, Comma separated list of tags.
func (s *Service) UploadSlideshow(username string, password string, uploadURL string, slideshow_title string, slideshow_description string, slideshow_tags string, make_src_public string) bool {
}
*/

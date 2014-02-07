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

// Slideshows type holds ... an array with Slideshows.
type SlideshowsByTag struct {
	TagName    string      `xml:"Name"`
	Count      uint32      `xml:"Count"`
	Slideshows []Slideshow `xml:"Slideshow"`
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

// GetSlideshowsByTag returns a Slideshows object:
// tag string required, holds the tag name.
// limit int optional, specify number of items to return.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
// return: Slideshows instance.
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
func (s *Service) GetSlideshowsByGroup(groupName string, detailed bool) (Slideshows, error)    {}

// GetSlideshowByUser returns a Slideshows object:
// username string required, username of the owner of slideshows.
// limit int optional, specify number of items to return.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
func (s *Service) GetSlideshowsByUser(username string, limit int, detailed bool) (Slideshows, error) {}

// Search for a slideshow:
// queryString required, search keyword
// page optional, the page number of the results (works in conjunction with items_per_page), default is 1
// items_per_page optional, number of results to return per page, default is 12
// lang optional, Language of slideshows (default is English, 'en') ('**':All,'es':Spanish,'pt':Portuguese,'fr':French,'it':Italian,'nl':Dutch, 'de':German,'zh':Chinese,'ja':Japanese,'ko':Korean,'ro':Romanian, '!!':Other)
// sort optional, Sort order, default is relevance.
func (s *Service) SearchSlideshows(queryString string, page string, items_per_page string, lang string, sort string) (Slideshows, error) {
}

// Edit a Slideshow, you can change only the title, tags and privacy of the slideshow.
// username required, owner username of the slideshow which is being edited
// password required, owner password of the slideshow which is being edited
// slideshowID required, id of slideshow which is being edited.
// slideshow_title optional, Title of the slideshow
// slideshow_ tags optional, Comma separated list of tags
// make_slideshow_private optional, Should be Y if you want to make the slideshow private.
func (s *Service) EditSlideshow(username string, password string, slideshowID string, slideshow_title string, slideshow_tags string, make_slideshow_private string) bool {
}

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

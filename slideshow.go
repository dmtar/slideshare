package slideshare

// Tags keeps an array with Tag IDs.
type Tags struct {
	Tags []Tag `xml:"Tags"`
}

// Tag type keeps tag IDs.
type Tag struct {
	ID uint64 `xml:"Tag"`
}

// RelatedSlideshowIDs ... holds an array with related slideshow IDs.
type RelatedSlideshowIDs struct {
	RelatedSlideshowIDs []RelatedSlideshowID
}

// RelatedSlideshowID holds info about the ID for a related slideshow.
type RelatedSlideshowID struct {
	ID uint64 `xml:"RelatedSlideshowID"`
}

// Slideshows type holds ... an array with Slideshows.
type Slideshows struct {
	Slideshows []Slideshow
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
	XMLName           xml.Name            `xml:"Slideshow"`
	ID                uint64              `xml:"ID"`
	Title             string              `xml:"Title"`
	Description       string              `xml:"Description"`
	Username          string              `xml:"Username"`
	Status            uint8               `xml:"Status"`
	Url               string              `xml:"URL"`
	ThumbnailUrl      string              `xml:"ThumbnailURL"`
	ThumbnailSize     string              `xml:"ThumbnailSize"`
	ThumbnailSmallUrl string              `xml:"ThumbnailSmallURL"`
	Embed             string              `xml:"Embed"`
	Created           string              `xml:"Created"`
	Updated           string              `xml:"Updated"`
	Language          string              `xml:"Language"`
	Format            string              `xml:"Format"`
	Download          bool                `xml:"Download"`
	DownloadUrl       string              `xml:"DownloadUrl"`
	SlideshowType     uint8               `xml:"SlideshowType"`
	InContest         bool                `xml:"InContest"`
	UserID            uint64              `xml:"UserID"`
	PPTLocation       string              `xml:"PPTLocation"`
	StrippedTitle     string              `xml:"StrippedTitle"`
	Tags              Tags                `xml:"Tags"`
	Audio             bool                `xml:"Audio"`
	NumDownloads      uint32              `xml:"NumDownloads"`
	NumViews          uint32              `xml:"NumViews"`
	NumComments       uint32              `xml:"NumComments"`
	NumFavorites      uint32              `xml:"NumFavorites"`
	NumSlides         uint16              `xml:"NumSlides"`
	RelatedSlideshows RelatedSlideshowIDs `xml:"RelatedSlideshows"`
	PrivacyLevel      bool                `xml:"PrivacyLevel"`
	FlagVisible       bool                `xml:"FlagVisible"`
	ShowOnSS          bool                `xml:"ShowOnSS"`
	SecretURL         bool                `xml:"SecretUrl"`
	AllowEmbed        bool                `xml:"AllowEmbed"`
	ShareWithContacs  bool                `xml:"ShareWithContacs"`
}

// GetSlideshow returns information about a slideshow, parameters:
// id int which holds the slideshow id, required.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
// return: Slideshow instance.
func (s *Service) GetSlideshow(id int, detailed bool) (Slideshow, error) {}

// GetSlideshowsByTag returns a Slideshows object:
// tag string required, holds the tag name.
// limit int optional, specify number of items to return.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
// return: Slideshows instance.
func (s *Service) GetSlideshowsByTag(tag string, limit int, detailed bool) (Slideshows, error) {}
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

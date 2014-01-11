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

// GetSlideshowsByTag returns returns a Slideshows object:
// tag string required, holds the tag name.
// limit int optional, specify number of items to return.
// detailed bool Whether or not to include optional information. true to include, false (default) for basic information.
// return: Slideshows instance.
func (s *Service) GetSlideshowsByTag(tag string, limit int, detailed bool) (Slideshows, error) {}
func (s *Service) GetSlideshowsByGroup(groupName string, detailed bool) (Slideshows, error)    {}
func (s *Service) GetSlideshowsByUser(userName string, detailed bool) (Slideshows, error)      {}
func (s *Service) SearchSlideshows(queryString string) (Slideshows, error)                     {}
func (s *Service) EditSlideshow(username string, password string, slideshowID string) bool     {}
func (s *Service) DeleteSlideshow(username string, password string, slideshowID string) bool   {}
func (s *Service) UploadSlideshow(username string, password string, uploadURL string, slideshowTitle string) bool {
}

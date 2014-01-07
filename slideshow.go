package slideshare

type Tags struct {
	Tags []Tag `xml:"Tags"`
}
type Tag struct {
	ID uint64 `xml:"Tag"`
}
type RelatedSlideshowIDs struct {
	RelatedSlideshowIDs []RelatedSlideshowID
}
type RelatedSlideshowID struct {
	ID uint64 `xml:"RelatedSlideshowID"`
}
type Slideshows struct {
	Slideshows []Slideshow
}
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

func (s *Service) GetSlideshow(id int, detailed bool) (Slideshow, error)                       {}
func (s *Service) GetSlideshowsByTag(tag string, limit int, detailed bool) (Slideshows, error) {}
func (s *Service) GetSlideshowsByGroup(groupName string, detailed bool) (Slideshows, error)    {}
func (s *Service) GetSlideshowsByUser(userName string, detailed bool) (Slideshows, error)      {}
func (s *Service) SearchSlideshows(queryString string) (Slideshows, error)                     {}
func (s *Service) EditSlideshow(username string, password string, slideshowID string) bool     {}
func (s *Service) DeleteSlideshow(username string, password string, slideshowID string) bool   {}
func (s *Service) UploadSlideshow(username string, password string, uploadURL string, slideshowTitle string) bool {
}

package slideshare

type EditedSlideshow struct {
	EditedID int `xml:"SlideShowID"`
}
type DeletedSlideshow struct {
	DeletedID int `xml:"SlideshowID"`
}
type UploadedSlideshow struct {
	UploadedID int `xml:"SlideShowID"`
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
	NumResults   uint32      `xml:"Meta>NumResults"`
	TotalResults uint64      `xml:"Meta>TotalResults"`
	Slideshows   []Slideshow `xml:"Slideshow"`
}

// Slideshow type, which holds all the information about a slideshow
// properties bellow InContest are detailed, they will have reliabe information
// if detailed flag is set to true.
type Slideshow struct {
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

// Package slideshare provides API wrapper for SlideShare.
package slideshare

const (
	APIURL = "https://www.slideshare.net/api/2/"
)

// Type Service defines the service, it needs Api key and Shared Secret
// which, the one who uses the service, should request from the SlideShare website.
type Service struct {
	Api_key      string
	SharedSecret string
}

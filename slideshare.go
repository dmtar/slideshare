// Package slideshare provides API wrapper for SlideShare.
package slideshare

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/url"
	"time"
)

var apiUrl = "https://www.slideshare.net/api/2"

// Type Service defines the service, it needs Api key and Shared Secret
// which, the one who uses the service, should request from the SlideShare website.
type Service struct {
	ApiKey       string
	SharedSecret string
}

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

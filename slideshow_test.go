package slideshare

import (
	"fmt"
	"testing"
)

const (
	ApiKey       = "5Pl6RFlI"
	SharedSecret = "X1lMfjPo"
)

func TestGetSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(29905515, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(slideshow.ID)
	fmt.Println(slideshow.Title)
	fmt.Println(slideshow.Description)
	fmt.Println(slideshow.Username)
	fmt.Println(slideshow.Status)
	fmt.Println(slideshow.Url)
	fmt.Println(slideshow.Created)
	fmt.Println(slideshow.NumDownloads)
	fmt.Println(slideshow.NumComments)
	fmt.Println(slideshow.NumSlides)
}

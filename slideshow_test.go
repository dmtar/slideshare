package slideshare

import (
	"testing"
)

const (
	ApiKey       = "5Pl6RFlI"
	SharedSecret = "X1lMfjPo"
	Password     = "golangtestproject"
	Username     = "ddishev"
)

// Test basic getting of slideshow
func TestGetSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30975136)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30975136 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
}

// Test getting a slideshow with detailed info
func TestGetSlideshowDetailed(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshow, err := service.GetSlideshow(30975136, true)
	if err != nil {
		t.Fatal(err)
	}
	if slideshow.ID != 30975136 {
		t.Fail()
	}
	if slideshow.Username != "ddishev" {
		t.Fail()
	}
	if slideshow.UserID != 59394728 {
		t.Fail()
	}
	if slideshow.Format != "pptx" {
		t.Fail()
	}
	if slideshow.NumSlides != 63 {
		t.Fail()
	}
	if slideshow.PrivacyLevel != false {
		t.Fail()
	}
	if slideshow.FlagVisible != true {
		t.Fail()
	}
	if slideshow.RelatedSlideshows[0] != 30098657 {
		t.Fail()
	}
	if slideshow.StrippedTitle != "databases-30975136" {
		t.Fail()
	}
}
func TestGetSlideshowsByTag(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.GetSlideshowsByTag("dishev", false, 1)
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.TagName != "dishev" {
		t.Fail()
	}
	if slideshows.Slideshows[0].Title != "Do not delete!" {
		t.Fail()
	}
	if slideshows.Slideshows[0].ID != 30975136 {
		t.Fail()
	}
	if slideshows.Slideshows[0].Format != "pptx" {
		t.Fail()
	}
}
func TestGetSlideshowsByUser(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.GetSlideshowsByUser("ddishev", true, 10)
	lastSlideshowIndex := slideshows.Count - 1
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.UserName != "ddishev" {
		t.Fail()
	}
	if slideshows.Slideshows[lastSlideshowIndex].ID != 30975136 {
		t.Fail()
	}
}

func TestEditSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isEdited := service.EditSlideshow(Username, Password, 30975136, "Do not delete!")
	if !isEdited {
		t.Fail()
	}
}

func TestSearchSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	slideshows, err := service.SearchSlideshows("golang", false, "1", "12")
	if err != nil {
		t.Fatal(err)
	}
	if slideshows.Slideshows[0].ID != 23464107 {
		t.Fail()
	}
}

var uploadedSlideshowID = 0

func TestUploadSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	id, isUploaded := service.UploadSlideshow(Username, Password, "http://www.fmi.uni-sofia.bg/Members/marian/42143844144243543c438-43e44143d43e43243043d438-43d430-43743d43043d43844f-44143f43544643843043b43d43e441442-41843d44443e44043c43044643843e43d43d438-44143844144243543c438-2011-2012-44344743543143d430-43343e43443843d430/KBS_Lecture4_1314.pdf", "Test Title", "Uploaded by API upload method")
	if !isUploaded || id == 0 {
		t.Fail()
	}
	uploadedSlideshowID = id
}

func TestDeleteSlideshow(t *testing.T) {
	service := Service{ApiKey, SharedSecret}
	isDeleted := service.DeleteSlideshow(Username, Password, uploadedSlideshowID)
	if !isDeleted {
		t.Fail()
	}
}

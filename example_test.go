package slideshare

import (
	"fmt"
	"github.com/dmtar/slideshare"
)

// You have to get API_KEY & SHARED_SECRET from Slideshare
// Here we are creating our service by providing the API_KEY & SHARED_SECRET
// After we call the method GetSlideshow with the Slideshow ID and without detailed information
// if you want detailed information for the slideshow, second argument has to be true, if everything is ok
// we have slideshow object with all the properties.
func ExampleService_GetSlideshow() {
	service := slideshare.Service{"5Pl6RFlI", "X1lMfjPo"}
	slideshow, err := service.GetSlideshow(29551397, false)

	fmt.Println("ID: ", slideshow.ID)
	fmt.Println("Title: ", slideshow.Title)
	fmt.Println("Description: ", slideshow.Description)
	fmt.Println("Username: ", slideshow.Username)
	// Output:
	// ID: 29551397
	// Title: Databases
	// Description: Representing Data Elements, presentation from FMI Databases course.
	// Username: ddishev
}

// Getting slideshows which have "sql" tag.
func ExampleService_GetSlideshowsByTag() {
	service := slideshare.Service{"5Pl6RFlI", "X1lMfjPo"}
	slideshows, err := service.GetSlideshowsByTag("dishev", false)

	fmt.Println("ID: ", slideshows.Slideshows[0].ID)
	fmt.Println("Title: ", slideshows.Slideshows[0].Title)
	fmt.Println("Username: ", slideshows.Slideshows[0].Username)
	// Output:
	// ID: 30975136
	// Title: Do not delete!
	// Username: ddishev
}
func ExampleService_GetSlideshowsByUser() {
	service := slideshare.Service{"5Pl6RFlI", "X1lMfjPo"}
	slideshows, err := service.GetSlideshowsByUser("ddishev", false)
	lastSlideshowIndex := slideshows.Count - 1

	fmt.Println("ID: ", slideshows.Slideshows[lastSlideshowIndex].ID)
	fmt.Println("Title: ", slideshows.Slideshows[lastSlideshowIndex].Title)
	fmt.Println("Username: ", slideshows.Slideshows[lastSlideshowIndex].Username)
	// Output:
	// ID: 30975136
	// Title: Do not delete!
	// Username: ddishev
}

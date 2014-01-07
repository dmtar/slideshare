package slideshare

import (
	"fmt"
	"github.com/dishev/go-slideshare"
)

// You have to get API_KEY & SHARED_SECRET from Slideshare
// Here we are creating our service by providing the API_KEY & SHARED_SECRET
// After we call the method GetSlideshow with the Slideshow ID and without detailed information
// if you want detailed information for the slideshow, second argument has to be true, if everything is ok
// we have slideshow object with all the properties.
func ExampleGetSlideshow() {
	service := slideshare.Service{"API_KEY", "SHARED_SECRET"}
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

// Same test but with detailed information
func ExampleGetSlideshow2() {
	service := slideshare.Service{"API_KEY", "SHARED_SECRET"}
	slideshow, err := service.GetSlideshow(29551397, true)

	fmt.Println("Number of downloads: ", slideshow.NumDownloads)
	fmt.Println("Number of views: ", slideshow.NumViews)
	fmt.Println("Number of comments: ", slideshow.NumComments)
	fmt.Println("Number of slides: ", slideshow.NumSlides)
	// Output:
	// Number of downloads: 1
	// Number of views: 48
	// Number of comments: 1
	// Number of slides: 63
}

// Getting slideshows which have "sql" tag.
func ExampleGetSlideshowsByTag() {
	service := slideshare.Service{"API_KEY", "SHARED_SECRET"}
	slideshows, err := service.GetSlideshowsByTag("sql", 10, false)

	fmt.Println("ID: ", slideshows.Slideshows[0].ID)
	fmt.Println("Title: ", slideshows.Slideshows[0].Title)
	fmt.Println("Description: ", slideshows.Slideshows[0].Description)
	fmt.Println("Username: ", slideshows.Slideshows[0].Username)
	// Output:
	// ID: 29748976
	// Title: Adding Value to HBase with IBM InfoSphere BigInsights and BigSQL
	// Description: This is the extended deck I used for my presentation at the Information On Demand 2013 conference for Session Number 1687 - Adding Value to HBase with IBM InfoSphere BigInsights and BigSQL.
	// This presentation covers accessing HBase using Big SQL. It starts by going over general HBase concepts, than delves into how Big SQL adds an SQL layer on top of HBase (via HBase storage handler), secondary index support, queries, etc.
	// Username: ppruski
}

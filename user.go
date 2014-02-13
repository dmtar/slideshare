package slideshare

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetUserFavorites needs username of user whose Favorites are being requested.
func (s *Service) GetUserFavorites(username_for string) (UserFavorites, error) {
	args := make(map[string]string)
	args["username_for"] = username_for
	url := s.generateUrl("get_user_favorites", args)
	resp, err := http.Get(url)
	if err != nil {
		return UserFavorites{}, err
	}
	favorites := UserFavorites{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &favorites)
	}
	return favorites, err
}

// GetUserContacts needs username of user whose Contacts are being requested.
// You can 1 additional parameter int to specify number of items to return.
func (s *Service) GetUserContacts(username_for string, limitOffset ...int) (UserContacts, error) {
	args := make(map[string]string)
	if limitOffset != nil {
		switch len(limitOffset) {
		case 1:
			args["limit"] = strconv.Itoa(limitOffset[0])
			break
		case 2:
			args["limit"] = strconv.Itoa(limitOffset[0])
			args["offset"] = strconv.Itoa(limitOffset[1])
			break
		default:
		}
	}
	args["username_for"] = username_for
	url := s.generateUrl("get_user_contacts", args)
	resp, err := http.Get(url)
	if err != nil {
		return UserContacts{}, err
	}
	contacts := UserContacts{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &contacts)
	}
	return contacts, err
}

// GetUserTags needs username and password of the requesting user.
func (s *Service) GetUserTags(username string, password string) (Tags, error) {
	args := make(map[string]string)
	args["username"] = username
	args["password"] = password
	url := s.generateUrl("get_user_tags", args)
	resp, err := http.Get(url)
	if err != nil {
		return Tags{}, err
	}
	tags := Tags{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		xml.Unmarshal([]byte(responseBody), &tags)
	}
	return tags, err
}

/*
// Returns user groups
// username_for required, username of user whose groups are being requested
func (s *Service) GetUserGroups(username_for string, additionalParams ...string) (Groups, error) {
	args := make(map[string]string)
	args["username_for"] = username_for
	url := s.generateUrl("get_user_groups", args)
	resp, err := http.Get(url)
	if err != nil {
		return Groups{}, err
	}
	groups := Groups{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		fmt.Println(responseBody)
		//xml.Unmarshal([]byte(responseBody), &groups)
	}
	return groups, err
}
*/

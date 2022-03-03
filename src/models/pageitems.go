package models

import "errors"

/*
GetPageItems return a list of Html objects data to use in the template and form rendering.
*/
func GetPageItems(pagename string) (Items, error) {
	if pagename == "" {
		return *new(Items), errors.New("page name is empty")
	}

	items, err := items(pagename)

	return items, err
}

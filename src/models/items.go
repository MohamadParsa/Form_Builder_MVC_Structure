package models

import (
	"errors"
)

/*
publicItem is used for polymorphism to make a map containing different struct types but with common functions.
*/
type publicItem interface {
	/*
		getData is a method of elements structures and get HTML element id and return required data to create an HTML object
		in the page.
	*/
	getData(id string) interface{}
}

/*
Items contain a list of HTML objects data.
a page has a list of HTML objects. for templates should send an instance of this object.
*/
type Items struct {
	Item []Item
}

/*
Item contain HTML object data. if the Type equals "textbox", ItemData is an instance of textbox struct.
so each item can have a different type for  ItemData.
*/
type Item struct {
	//The ID of HTML object on the page
	ElementID string
	//represents type of HTML object such as "textbox","br"...
	Type string
	//contain information of HTML object to show on the page such as title, value...
	ItemData interface{}
}

/*
items function gets a page name and returns a dynamic list of defined HTML objects for the page.
*/
func items(pagename string) (Items, error) {

	items, err := getitems(pagename)
	if err != nil {
		return items, err
	}

	itemsConstructors := makeItemsMap()

	for index, item := range items.Item {
		items.Item[index].ItemData = itemsConstructors[item.Type].getData(item.ElementID)
	}
	return items, nil
}

/*
getitems function gets a page name and connects to the storage and returns a dynamic list of defined HTML objects for the page.
*/
func getitems(pagename string) (Items, error) {
	var items Items
	if pagename == "" {
		return items, errors.New("page name is empty")
	}

	/*
		in this step, we should connect to a database or any other storage and send "pagename" then get items information.
		fetching data don't implement so just a piece of constant information has been returned.
		you need to develop this section as your logic.
	*/

	if pagename == "index.html" {
		items = Items{
			[]Item{
				{ElementID: "txtLogin", Type: "textbox"},
				{ElementID: "br", Type: "br"},
				{ElementID: "btnLogin", Type: "button"},
			}}
		return items, nil

	}
	return items, errors.New("the page hasn't any items")
}

/*
makeItemsMap function return a map contain implemented HTML objects. this map created to call getData function for each item
and getting item data from databases or other storage.
*/
func makeItemsMap() map[string]publicItem {

	itemmap := make(map[string]publicItem)

	itemmap["br"] = br{}
	itemmap["button"] = button{}
	itemmap["textbox"] = textbox{}

	return (itemmap)

}

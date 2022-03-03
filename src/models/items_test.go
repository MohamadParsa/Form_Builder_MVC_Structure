package models

import "testing"

func Test_makeItemsMap(t *testing.T) {
	if itemMap := makeItemsMap(); len(itemMap) == 0 {
		t.Log("makeItemsMap func doesn't work.")
		t.Fail()
	}
}
func Test_getItemsParmeterValidator(t *testing.T) {
	if _, err := getitems(""); err == nil {
		t.Log("getitems shouldn't accept empty as pagename param. ")
		t.Fail()
	}
}
func Test_getItemsInvalidPageName(t *testing.T) {
	if _, err := getitems("an invalid Page Name"); err == nil {
		t.Log("getitems shouldn't accept an invalid pagename. ")
		t.Fail()
	}
}
func Test_itemsParmeterValidator(t *testing.T) {
	if _, err := items(""); err == nil {
		t.Log("items shouldn't accept empty as pagename param. ")
		t.Fail()
	}
}
func Test_itemsInvalidPageName(t *testing.T) {
	if _, err := items("an invalid Page Name"); err == nil {
		t.Log("items shouldn't accept an invalid pagename. ")
		t.Fail()
	}
}

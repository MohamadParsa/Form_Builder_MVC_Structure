package models

import "testing"

func Test_getPageItemsParmeterValidator(t *testing.T) {
	if _, err := GetPageItems(""); err == nil {
		t.Log("GetPageItems shouldn't accept empty as pagename param. ")
		t.Fail()
	}
}

func Test_getPageItemsInvalidPageName(t *testing.T) {
	if _, err := GetPageItems("an invalid Page Name"); err == nil {
		t.Log("GetPageItems shouldn't accept an invalid pagename. ")
		t.Fail()
	}
}

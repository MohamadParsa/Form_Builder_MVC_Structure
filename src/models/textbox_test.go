package models

import (
	"testing"
)

func Test_textboxCreate(t *testing.T) {
	var textboxItem textbox
	result := textboxItem.getData("test")
	_, ok := result.(textbox)
	if !ok {
		t.Log("getData doesn't return an instance of textbox ")
		t.Fail()
	}
}

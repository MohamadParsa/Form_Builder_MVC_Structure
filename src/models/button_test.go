package models

import (
	"testing"
)

func Test_buttonCreate(t *testing.T) {
	var buttonItem button
	result := buttonItem.getData("test")
	_, ok := result.(button)
	if !ok {
		t.Log("getData doesn't return an instance of button ")
		t.Fail()
	}
}

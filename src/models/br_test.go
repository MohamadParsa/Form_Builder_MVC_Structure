package models

import (
	"testing"
)

func Test_brCreate(t *testing.T) {
	var brItem br
	result := brItem.getData("test")
	_, ok := result.(br)
	if !ok {
		t.Log("getData doesn't return an instance of br ")
		t.Fail()
	}
}

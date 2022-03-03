package controllers

import "testing"

func Test_getContentType(t *testing.T) {
	if ContentType, err := getContentType(""); err == nil || ContentType != "" {
		t.Log("getContentType func shouldn't accept empty as param.")
		t.Fail()
	}
	if ContentType, err := getContentType("url\file"); err != nil || ContentType != "text/plain" {
		t.Log("The unknow file is not supplied.")
		t.Fail()
	}

	if ContentType, err := getContentType("url\file.css"); err != nil || ContentType != "text/css" {
		t.Log("The CSS file is not supplied.")
		t.Fail()
	}
	if ContentType, err := getContentType("url\file.html"); err != nil || ContentType != "text/html" {
		t.Log("The HTML file is not supplied.")
		t.Fail()
	}
	if ContentType, err := getContentType("url\file.js"); err != nil || ContentType != "application/javascript" {
		t.Log("The JS file is not supplied.")
		t.Fail()
	}
	if ContentType, err := getContentType("url\file.png"); err != nil || ContentType != "image/png" {
		t.Log("The PNG file is not supplied.")
		t.Fail()
	}
	if ContentType, err := getContentType("url\file.mp4"); err != nil || ContentType != "video/mp4" {
		t.Log("The MP4 file is not supplied.")
		t.Fail()
	}
}

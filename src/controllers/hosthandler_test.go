package controllers

import (
	"testing"
)

func Test_HostHandlerParametersValidation(t *testing.T) {
	setting := Settings{
		RootWebPath:      "",
		DefaultPageName:  "/index.html",
		NotFindePagePath: "/ErrorPages/PageNotFound.html"}

	if _, err := HostHandler("", 8099, setting); err == nil {
		t.Log("HostHandler shouldn't accept empty as routing param. ")
		t.Fail()
	}
	if _, err := HostHandler("/", 0, setting); err == nil {
		t.Log("HostHandler shouldn't accept 0 as port param. ")
		t.Fail()
	}
	if _, err := HostHandler("/", 8099, setting); err == nil {
		t.Log("HostHandler shouldn't accept empty as setting.RootWebPath param. ")
		t.Fail()
	}
}

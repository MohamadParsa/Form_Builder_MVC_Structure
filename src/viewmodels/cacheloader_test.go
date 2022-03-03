package viewmodels

import (
	"fmt"
	"testing"
	"text/template"
)

func Test_ObjectsLoadeParameterValidation(t *testing.T) {
	if err := CacheFiles("", template.New("template")); err == nil {
		t.Log("tmplatesFolderName shouldn't accept empty as routing param. ")
		t.Fail()
	}

}
func Test_FolderNotFind(t *testing.T) {
	if err := CacheFiles("other folder name", template.New("template")); err == nil {
		t.Log("checking exists folder does not work.")
		t.Fail()
	}
}
func Test_ObjectsLoade(t *testing.T) {
	if err := CacheFiles("htmlobjects", template.New("template")); err != nil {
		fmt.Println(err)
		t.Log("objectsLoade func doesn't work.")
		t.Fail()
	}
}

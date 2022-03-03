package viewmodels

import (
	"errors"
	"os"
	"text/template"
)

/*
CacheFiles function read all files of tmplatesFolderName address, then parse and load them into the
template.Template variable and finally function return error to exception handling.
*/
func CacheFiles(tmplatesFolderName string, template *template.Template) error {

	//open the folder that contain our Html Pages and objects like *.html, buttom , textbox ,...
	if tmplatesFolderName == "" {
		return errors.New("folder path is empty")
	}

	//open the folder
	htmlObjectsFolder, errOpen := os.Open(tmplatesFolderName)
	if errOpen != nil {
		return errors.New("folder not found")
	}
	defer htmlObjectsFolder.Close()

	//find all files path and store in templateObjectsPaths slice
	htmlObjectsPaths, errDir := htmlObjectsFolder.ReadDir(-1)
	if errDir != nil {
		return errors.New("files not found")
	}
	templateObjectsPaths := new([]string)
	for _, templatePath := range htmlObjectsPaths {
		if !templatePath.IsDir() {
			*templateObjectsPaths = append(*templateObjectsPaths, tmplatesFolderName+"/"+templatePath.Name())
		}
	}

	//parse and add all files as templates
	_, errParseFiles := template.ParseFiles(*templateObjectsPaths...)

	return errParseFiles
}

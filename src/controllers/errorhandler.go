package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
errorPage writes errors in http.ResponseWriter.
it does with error handling as a not found page or a simple text.
*/
func errorPage(w http.ResponseWriter, errorPageAddress string, errorNumber int, errorDetail error) {
	fmt.Println(errorDetail)
	notFoundPageData, errorInReadingNotFoundPage := ioutil.ReadFile(errorPageAddress)

	if errorInReadingNotFoundPage == nil {
		w.Write(notFoundPageData)
	} else {
		w.WriteHeader(errorNumber)
		w.Write([]byte(strconv.Itoa(errorNumber) + " - " + http.StatusText(errorNumber)))
	}
}

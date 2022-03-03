package controllers

import (
	"errors"
	"strings"
)

//getContentType function return content type acording requested file
func getContentType(reqURLPah string) (string, error) {
	if reqURLPah == "" {
		return "", errors.New("url is invalid. ")
	}
	contentType := "text/plain"
	switch {
	case strings.HasSuffix(reqURLPah, ".css"):
		contentType = "text/css"
	case strings.HasSuffix(reqURLPah, ".html"):
		contentType = "text/html"
	case strings.HasSuffix(reqURLPah, ".js"):
		contentType = "application/javascript"
	case strings.HasSuffix(reqURLPah, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(reqURLPah, ".mp4"):
		contentType = "video/mp4"
	}
	return contentType, nil
}

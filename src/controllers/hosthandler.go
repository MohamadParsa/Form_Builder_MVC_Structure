package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"viewmodels"
)

/*
http.Handler is an interface from net/http, It enables us to overwrite the ServeHTTP function.
so we can define a method with a receiver such as ( this *handler ) and overwrite the ServeHTTP function.
*/
type handler struct {
	http.Handler
	handlerSetting Settings
}

/*
Settings struct contain server running configure. such as physical web pages addresses and error pages addresses.
*/
type Settings struct {
	//RootWebPath web pages physical address. based on the Main package.
	RootWebPath string
	//				Custom Pages Address				//
	//NotFindePagePath contain custom 404 error page address. if didn't set, the server returns a simple 404 message.
	NotFindePagePath string
	//DefaultPageName contains your default page, like index.html.
	DefaultPageName string
	//ErrorPage contains a custom 500 error page address. if didn't set, the server returns a simple 500 message.
	ErrorPage string
}

//ServeHTTP function is a overwrite of ServeHTTP in net/http to get custom settings.
func (handlerInstance *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//if url doesn't contain page name , we can use default page name from setting.
	reqURLPath := req.URL.Path
	if reqURLPath == "" || reqURLPath == "/" {
		reqURLPath = handlerInstance.handlerSetting.DefaultPageName
	}
	fmt.Println(reqURLPath)

	/*
		in this step pages and HTML objects are read from the disk and loaded on the memory.
		cacheLoader has a method called ObjectsLoade that read all files in the folder
		and add into the template. remember template variable is a pointer.
		if any error in the CacheFiles function exists, the server should return the error number 500.
	*/
	templates := template.New("templates")

	errCache := viewmodels.CacheFiles("viewmodels/webpages", templates)
	if errCache != nil {
		errorPage(w, handlerInstance.handlerSetting.ErrorPage, 500, errCache)
		return
	}

	errCache = viewmodels.CacheFiles("viewmodels/htmlobjects", templates)
	if errCache != nil {
		errorPage(w, handlerInstance.handlerSetting.ErrorPage, 500, errCache)
		return
	}
	/*
		if the user types index instead of index.html the "index.html" page should open. so we first remove suffix and then add ".html"
		in the templates variable, there are all files as templates and actually, the file name is
		template name.
	*/
	fileName := strings.TrimSuffix(reqURLPath[1:], ".html") + ".html"

	/*
		extract requested template with applied subtemplate
	*/
	requestedTemplate := templates.Lookup(fileName)

	/*
		if the template exists, return the template as a response,
		else check for the existing physical files.
		if any error in the GetPageItems function exists, the server should return the error number 500.
	*/
	if requestedTemplate != nil {

		context, errGetPageItems := models.GetPageItems(fileName)

		if errGetPageItems != nil {
			errorPage(w, handlerInstance.handlerSetting.ErrorPage, 500, errGetPageItems)
			return
		}
		requestedTemplate.Execute(w, context)

	} else {

		pageData, errorInReadingPage := os.Open(handlerInstance.handlerSetting.RootWebPath + reqURLPath)

		/*
			if the desired file exists, return the file as a response.
			we use bufferReader to avoid the problem of large file loading.
		*/
		if errorInReadingPage == nil {

			bufferReader := bufio.NewReader(pageData)
			contentType, err := getContentType(reqURLPath)
			if err == nil {
				w.Header().Add("Content Type", contentType)
				bufferReader.WriteTo(w)
			}

		} else {
			/*
				if the URL path wasn't in templates or web page files, the server should return the error number 404.
			*/
			errorPage(w, handlerInstance.handlerSetting.NotFindePagePath, 404, errors.New("page not found"))
		}
	}
}

/*
HostHandler function runs an instance of the webserver.
routing contains routing address for example "/" for all routs or "/cars" for handle requests of domain/cars
port determine the webserver listening port
*/
func HostHandler(routing string, port int, setting Settings) (status string, err error) {
	status = "initialize"
	fmt.Println(status)

	if routing == "" || strings.Index(routing, "/") != 0 {
		return status, errors.New("routing is invalid")
	} else if port == 0 {
		return status, errors.New("port is invalid")
	} else if setting.RootWebPath == "" {
		return status, errors.New("RootWebPath can't be empty")
	}

	status = "running"
	fmt.Println(status)

	http.Handle(routing, &handler{handlerSetting: setting})
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

	status = "stop"
	fmt.Println(status)

	return status, nil

}

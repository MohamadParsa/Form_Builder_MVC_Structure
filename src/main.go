//Main Package run provide a selfe host app .
package main

import (
	"controllers"
	"fmt"
)

func main() {
	setting := controllers.Settings{
		RootWebPath:      "viewmodels/webpages",
		DefaultPageName:  "/index.html",
		NotFindePagePath: "viewmodels/errorpages/pageNotFound.html",
		ErrorPage:        "viewmodels/errorpages/servererror.html",
	}
	status, err := controllers.HostHandler("/", 8099, setting)
	fmt.Println(status, err)
}

package app

import "github.com/gin-gonic/gin"

/*
* Every request that our application receives
* will be handled by this router
 */
var (
	router = gin.Default()
)

//StartApplication - Starts with Caps for export purpose
func StartApplication() {
	//Handle the URLs
	mapUrls()

	//Run the server in port 8080
	router.Run(":8080")
}

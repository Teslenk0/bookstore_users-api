package app

import (
	"github.com/Teslenk0/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

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

	logger.Info("about to start the application")

	//Run the server in port 8080
	router.Run(":8082")
}

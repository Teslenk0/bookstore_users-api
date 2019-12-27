package app

import (
	"github.com/Teslenk0/bookstore_users-api/controllers/ping"
	"github.com/Teslenk0/bookstore_users-api/controllers/users"
)

func mapUrls() {

	//---------------------- GET INFORMATION -----------------------------
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.GET("/internal/users/search", users.Search)

	//---------------------- SAVES INFORMATION --------------------------------

	router.POST("/users", users.Create)

	//------------------- UPDATE INFORMATION ----------------------------------

	//Complete Update
	router.PUT("/users/:user_id", users.Update)

	//Partial Update
	router.PATCH("/users/:user_id", users.Update)

	//----------------------- DELETE INFORMATION -----------------------------------

	router.DELETE("/users/:user_id", users.Delete)

	//--------------------------------------------------------------------------------

}

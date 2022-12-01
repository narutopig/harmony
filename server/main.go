package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgradeHandler = websocket.Upgrader{}
//default upgrade settings

func connectHandler(w http.ResponseWriter, r *http.Request) { //value not reference 

	connection, err = upgradeHandler.Upgrade(w, r, nil)
	if (err != nil) {
		log.Println("CONNECTION ISSUE")
		log.panic(err)

	}

}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API is up!",
			})
		})
	}

	router.Run(":3000")

	// api will impleemnt get and post alter
}

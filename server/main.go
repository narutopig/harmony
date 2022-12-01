package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgradeHandler = websocket.Upgrader{}

//default upgrade settings

//doing the funny to make a func for a simple error check

func checkErr(err error) {
	if err != nil {
		log.Println("ERROR THROWN: ")
		log.Panic(err)
		return
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request) { //value not reference

	connection, err := upgradeHandler.Upgrade(w, r, nil)
	checkErr(err)
	defer connection.Close()

	log.Println("Connection Established")

	connectReader(connection)

}

func connectReader(co *websocket.Conn) {
	for {
		messageType, content, err := co.ReadMessage()
		checkErr(err)

		fmt.Println("Read message: " + string(content))

		if err := co.WriteMessage(messageType, content); err != nil {
			//retursne rror if comes across one, will execute writemssage func and will nly retirn error if there is one
			log.Panic(err)
		}
	}
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			connectHandler(c.Writer, c.Request) //on req establish connection or just in general
		})
	}

	router.Run(":3000")

	// api will impleemnt get and post alter
}

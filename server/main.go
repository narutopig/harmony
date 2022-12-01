package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgradeHandler = websocket.Upgrader{}
//default upgrade settings

//doing the funny to make a func for a simple error check

func checkErr(err any) {
	if (err != nil) {
		log.Println("ERROR THROWN: ")
		log.panic(err)
		return
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request) { //value not reference 

	connection, err = upgradeHandler.Upgrade(w, r, nil)
	checkErr(err)
	defer connection.Close()

	log.Println("Connection Established")

	connectReader(connection)


}

func connectReader(co *websocket.Conn) {
	for {
		type, content, err = co.ReadMessage()
		checkErr(err)
		
		fmt.Println("Read message: " + content)

		if er := co.WriteMessage(messageType, p); er != nil {
			//retursne rror if comes across one, will execute writemssage func and will nly retirn error if there is one 
			log.panic(er)
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
			c.JSON(http.StatusOK, gin.H{
				"message": "API is up!",
			})
		})
	}

	router.Run(":3000")

	// api will impleemnt get and post alter
}

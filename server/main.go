package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgradeHandler = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Method == "GET" && r.URL.Path == "/api/ws"
	},
}

// default upgrade settings

// doing the funny to make a func for a simple error check
func checkErr(err error) {
	if err != nil {
		log.Println("ERROR THROWN: ")
		log.Panic(err)
		return
	}
}

func sendPacket(co *websocket.Conn, content []byte) {
	// even tho cotnent is techncially a byte array its lkay of io use any
	// experimental forever loop

	msgT, _, err := co.ReadMessage()
	// assuming that error is not existent
	checkErr(err) // will break out with panic log if error
	co.WriteMessage(msgT, content)
}

func connectHandler(w http.ResponseWriter, r *http.Request) { // value not reference

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

		log.Printf("Read message: %s\n", content)

		if err := co.WriteMessage(messageType, content); err != nil {
			// retursne rror if comes across one, will execute writemssage func and will nly retirn error if there is one
			log.Panic(err)
		}

		go sendPacket(co, []byte("connected"))
	}
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	api := router.Group("/api")
	{
		api.GET("/ws", func(c *gin.Context) {
			connectHandler(c.Writer, c.Request) // on req establish connection or just in general
		})
	}

	router.Run(":3000")

	// api will impleemnt get and post alter
}

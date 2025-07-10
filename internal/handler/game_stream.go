package handler

import (
	"go_learn/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var SocketConnection *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ WARNING: This allows all origins. Use proper validation in production.
	},
}

func Socket(context *gin.Context) {
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	defer conn.Close()
	SocketConnection = conn

	for {

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Println(messageType, string(message))

	}

}

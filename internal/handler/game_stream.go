package handler

import (
	"go_learn/internal/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
		}
	}()
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte("hello world"))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
		time.Sleep(5 * time.Second)
	}
}

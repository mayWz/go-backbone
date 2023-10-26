package handler

import (
	"go-chat/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler interface {
	Chat(c *gin.Context)
}

type chatHandler struct {
}

func NewChatHandler() *chatHandler {
	return &chatHandler{}
}

var users = make(map[*model.ChatUser]bool)

func (ch *chatHandler) Chat(c *gin.Context) {
	roomID := c.Param("roomId")
	userID := c.Param("userID")

	//TODO: check chatroom is exist or not. if not create a new room
	//TODO: check user is already in the chat room or ont

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	chatUser := &model.ChatUser{Conn: conn, UserID: userID, RoomID: roomID}
	users[chatUser] = true

	for {
		// Reads and handles WebSocket messages
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			delete(users, chatUser)
			break
		}

		// Broadcast the message to all connected users in the same chat room
		for user := range users {
			if user.Conn != conn && user.RoomID == roomID {
				err := user.Conn.WriteMessage(messageType, p)
				//TODO: store message data to database
				if err != nil {
					delete(users, user)
					break
				}
			}
		}
	}
}

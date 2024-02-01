package server

import (
	// "log"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

func init() {
	// Initialize AllRooms when the package is imported
	AllRooms.Init()
}

// /this will create room and return roomID
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID := AllRooms.CreateRoom()
	type resp struct {
		RoomID string `json:"room_id`
	}
	fmt.Println("ddcdfdc")

	log.Println(AllRooms)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

type broadcastingMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("jnhjihbkj")

	roomId, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("roomID missing in the URL Parameteer")
		return
	}
	fmt.Println(roomId)
	return
}

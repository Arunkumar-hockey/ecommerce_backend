package websocket

import (
	"TFP/service"
	json2 "encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("false...")
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Println(t)
			result, err := service.ProductsService.GetAll()
			if err != nil {
				fmt.Println("Unable to fetch products")
			}

			json, Err := json2.Marshal(result)
			if Err != nil {
				fmt.Println(Err)
			}

			err1 := conn.WriteMessage(websocket.TextMessage, []byte(json))
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			fmt.Println(result)
		}
	}
}

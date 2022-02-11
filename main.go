package main

import (
	"fmt"
	"log"
	"net/http"
	"poker/mechanics"

	"github.com/gorilla/websocket"
)

var messages = make(chan []byte)
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {

	// hands := [...]string{
	// 	"2h 2d 2c kc qd",
	// 	"2h 5h 7d 8c 9s",
	// 	"ah 2d 3c 4c 5d",
	// 	"2h 3h 2d 3c 3d",
	// 	"2h 7h 2d 3c 3d",
	// 	"2h 7h 7d 7c 7s",
	// 	"th jh qh kh ah",
	// 	"4h 4s ks 5d ts",
	// 	"qc tc 7c 6c 4c",
	// 	"ah ah 7c 6c 4c",
	// }

	hand := []string{"2h", "2d", "2c", "kc", "qd"}
	fmt.Println(mechanics.AnalyzeHand(hand))

	//routes()
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func reverseByte(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reverseByte(input[1:]), input[0])
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func socket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	clients[ws] = true
	log.Println("Client connected...")
	go reader(ws)
}

func reader(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		for k, _ := range clients {
			if k == conn {
				continue
			} else {
				if err := k.WriteMessage(messageType, p); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func routes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", socket)
}

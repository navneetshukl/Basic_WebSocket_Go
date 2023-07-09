package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	mux := chi.NewRouter()

	mux.Get("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		for {

			//The ReadMessage method is used to read a message from the WebSocket connection.
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			s:=msg;
			n:=len(s)
			msg = []byte(fmt.Sprintf("The Length of given string is %d", n))

			
			//The WriteMessage method is used to write a message to the WebSocket connection.
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	mux.Get("/",func(w http.ResponseWriter,r *http.Request) {
		http.ServeFile(w,r,"websocket.html")
	})



	http.ListenAndServe(":8080", mux)
}

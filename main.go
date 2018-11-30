package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin:	func(r *http.Request) bool { return true},
}

func main(){
	http.HandleFunc("/", wsHandler)
	http.ListenAndServe(":8080", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request){
	sock, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for {
		mt, msg, err := sock.ReadMessage()
		if err != nil{
			fmt.Printf("%v", err)
			return
		}
		fmt.Println(string(msg))
		if err = sock.WriteMessage(mt, msg); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}
}
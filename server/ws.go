package server

import (
	"net/http"
	"log"

	"github.com/spf13/viper"
	
	"local.com/sai0556/Go-ChatwithRabbitMQ/client"
)

type WebSocket struct{}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (ws *WebSocket) Start() {
	hub := client.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(viper.GetString("ws.addr"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// func (ws *WebSocket) Start() {
// 	hub := client.NewHub()
// 	go hub.Run()
	
// 	http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		client.ServeWs(hub, w, r)
// 	})
// 	err := http.ListenAndServe(viper.GetString("ws.addr"), nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }


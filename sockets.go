package main

import (
	"log"
	"time"

	"github.com/googollee/go-socket.io"
	"github.com/owenso/crypto-portfolio-api/utils"
)

// InitializeSockets
func InitializeSockets() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("Client Connected")
		so.Join("polo")
		utils.DoEvery(10*time.Second, func(t time.Time) {
			result := CallCMC(t)
			so.Emit("polo", result)
		})
		so.On("disconnection", func() {
			log.Println("SOCKET DISCONNECTED")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}

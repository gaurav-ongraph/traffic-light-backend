package main

import (
	"time"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
	// common "github.com/frame-lang/frame-demos/persistenttrafficlight/common"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		trafficlight.SocketConn = conn
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if string(p) == "start" {
			res := trafficlight.CreateResponse("begin", "", "Creating...", true)
			trafficlight.SendResponse(res)
			time.Sleep(1 * time.Second)
			trafficlight.MOM.Start()
		} else if string(p) == "error" {
			trafficlight.MOM.SystemError()
		} else if string(p) == "restart" {
			trafficlight.MOM.SystemRestart()
		} else if string(p) == "end" {
			log.Println("END")
			trafficlight.MOM.Stop()
			res := trafficlight.CreateResponse("end", "", "Create Traffic Light", false)
			trafficlight.SendResponse(res)
		}
	}
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected...")
	reader(ws)
	
}


func main() {
	// stop := make(chan bool)
	// finished := make(chan bool)
	trafficlight.CreateNewTrafficLight()

	// Create End Point for Web Socket
	http.HandleFunc("/ws", wsEndPoint)

	// Create HTTP server
	log.Fatal(http.ListenAndServe(":8000", nil))

	// if err != nil {
	// 	log.Fatal(err)	// }
	// ticker := time.NewTicker(1000 * time.Millisecond)
	// 
	// go func() {
	// 	for {
	// 		select {
	// 		case <-stop:
	// 			ticker.Stop()
	// 			mom.Stop()
	// 			finished <- true
	// 			return
	// 		case <-ticker.C:
	// 			fmt.Println("tick")
	// 			mom.Tick()
	// 		}

	// 	}
	// }()

	// stop <- true
	// <-finished
	// fmt.Println("finished")
}

package trafficlight

import (
	"log"
	"time"
	"reflect"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
)

type commandStruct struct {
	Name   string   `json:"name"`
	Message string `json:"message"`
	Text string `json:"text"`
	Loading bool `json:"loading"`
}

var Stopper chan<- bool
var SocketConn *websocket.Conn
var MOM = trafficlight.NewTrafficLightMom()

func CreateResponse(state string, message string, text string, loading bool) string {
	command := &commandStruct{
		Name: state,
		Message: message,
		Text: text,
		Loading: loading,
	}

	json, _ := json.Marshal(command)
	return string(json)
}

func SendResponse(data string) {
	if err := SocketConn.WriteJSON(data); err != nil {
		log.Println(err)
		return
	}
}

func SetInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				log.Println("stop setInterval")
				return
			case <-ticker.C:
				reflect.ValueOf(p).Call([]reflect.Value{})
			}
		}

	}()

  	// return the bool channel to use it as a stopper
	return stopIt
}
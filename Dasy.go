package Dasy

import (
	"strconv"
	"fmt"
	"log"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

/* Types */
type Server struct {
	Port int
	Clients []Client
}

type Client struct {
	ID uuid.UUID
}

/* Server functions */

func CreateServer(port int) Server {
	return Server{Port: port, Clients: []Client{}}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (s Server) Start() error {
	fmt.Println("Server started on port " + strconv.Itoa(s.Port))

	http.HandleFunc("/echo", serverHandler)

	return nil 	
}

/* Client functions */

func CreateClient() Client {
	id := uuid.New()
	return Client{ID: id}
}

func (c Client) ConnectToServer(address string) error {
	return errors.New("Could not connect to server " + address)
}
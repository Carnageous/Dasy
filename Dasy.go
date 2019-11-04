package Dasy

import (
	"strconv"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
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

func (s Server) Start() {
	
	http.HandleFunc("/echo", serverHandler)
	
	port := strconv.Itoa(s.Port)
	
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe("localhost:" + port, nil))
}

/* Client functions */

func CreateClient() Client {
	id := uuid.New()
	return Client{ID: id}
}

func (c Client) ConnectToServer(address string) error {
	// Needs to be rewritten using WS over HTTP (not TCP, obviously)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: address, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	wc, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer wc.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := wc.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return nil
		case t := <-ticker.C:
			err := wc.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return nil
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := wc.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}
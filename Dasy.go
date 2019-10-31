package Dasy

import (
	"fmt"
	"errors"
	"strconv"
	"github.com/google/uuid"
)

/* Types */

type Client struct {
	ID uuid.UUID
}

type Server struct {
	Port int
	Clients []Client
}

/* Server functions */

func CreateServer(port int) Server {
	return Server{Port: port, Clients: []Client{}}
}

func (s Server) Start() error {
	fmt.Println("Server started on port " + strconv.Itoa(s.Port))

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
package main

import (
	"fmt"
	Dasy "github.com/Carnageous/dasy"
)

func main() {
	c := Dasy.CreateClient()

	fmt.Println(c.ID)

	err := c.ConnectToServer("127.0.0.1")

	if err != nil {
		fmt.Println(err)
	}
}
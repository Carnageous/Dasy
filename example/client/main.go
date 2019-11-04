package main

import (
	"fmt"
	Dasy "github.com/Carnageous/dasy"
)

func main() {
	c := Dasy.CreateClient()

	fmt.Println(c.ID)

	err := c.ConnectToServer("localhost:3004")

	if err != nil {
		fmt.Println(err)
	}
}
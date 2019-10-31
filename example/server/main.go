package main

import (
	Dasy "github.com/Carnageous/Dasy"
)

func main() {
	s := Dasy.CreateServer(3004)

	s.Start()
}
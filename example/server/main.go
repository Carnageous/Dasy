package main

import (
	Dasy "github.com/Carnageous/dasy"
)

func main() {
	s := Dasy.CreateServer(3004)

	s.Start()
}
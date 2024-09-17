package main

import (
	"fmt"
	"lemin/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	filename := os.Args[1]
	RoomStruct := functions.RoomStruct{}
	correctfile := RoomStruct.CheckLemin(filename)
	if !correctfile {
		return
	}
	fmt.Println(RoomStruct)
}

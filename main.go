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
	fmt.Println("Starting room: ", RoomStruct.StartingRoom)
	fmt.Println("Ending room: ", RoomStruct.EndingRoom)
	fmt.Println("Number of Ants: ", RoomStruct.Ants)
	for i := 0; i < len(RoomStruct.AllRooms); i++ {
		fmt.Print("Room name: ", RoomStruct.AllRooms[i].Name, ", ")
		fmt.Print("Links: ")
		fmt.Print(RoomStruct.AllRooms[i].Links)
		fmt.Println()
	}
}

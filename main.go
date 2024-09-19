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
	// fmt.Println("Starting room: ", RoomStruct.StartingRoom)
	// fmt.Println("Ending room: ", RoomStruct.EndingRoom)
	// fmt.Println("Number of Ants: ", RoomStruct.Ants)
	// for i := 0; i < len(RoomStruct.AllRooms); i++ {
	// 	fmt.Print("Room name: ", RoomStruct.AllRooms[i].Name, ", ")
	// 	fmt.Print("Links: ")
	// 	fmt.Print(RoomStruct.AllRooms[i].Links)
	// 	fmt.Println()
	// }
	RoomMap := make(map[string]*functions.Rooms)
	for i := 0; i < len(RoomStruct.AllRooms); i++ {
		RoomMap[RoomStruct.AllRooms[i].Name] = &RoomStruct.AllRooms[i]
	}
	Paths := functions.FindAllPathsBFS(RoomMap, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)
	fmt.Println(Paths)
	BestPath := functions.FilterPath(Paths, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)
	fmt.Println("Best paths: ", BestPath)

	ant := functions.AntInPath(BestPath, RoomStruct.Ants)
	for i := 0; i < len(ant); i++ {
		fmt.Println("Path numero:", i+1, "|| nombre de fourmi:", ant[i])
	}
	// test := [][]string{}
	// t1 := []string{"e", "q", "a", "d"}
	// t2 := []string{"s"}
	// test = append(test, t1)

	// fmt.Println(functions.CheckPath(test, t2))
}

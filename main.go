package main

import (
	"fmt"
	"lemin/functions"
	"os"
)

func main() {
	//Checking if we have the good lenght of argument
	if len(os.Args) != 2 {
		fmt.Println("Invalid Input, Usage: go run . [filename]")
		return
	}

	//Storing the filename into a variable
	filename := os.Args[1]
	RoomStruct := functions.RoomStruct{}

	//Checking if the file is valid or not
	correctfile := RoomStruct.CheckLemin(filename)
	if !correctfile {
		//If not we just return
		return
	}

	//Storing the Rooms info on a struct with their Name as key to that value
	RoomMap := make(map[string]*functions.Rooms)
	for i := 0; i < len(RoomStruct.AllRooms); i++ {
		RoomMap[RoomStruct.AllRooms[i].Name] = &RoomStruct.AllRooms[i]
	}

	//Getting all the possible path using the BFS Algorithm
	Paths := functions.FindAllPathsBFS(RoomMap, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)
	// fmt.Println(Paths)

	//Filtering the solution to get the most path possibles that doesn't use similar rooms
	BestPath := functions.FilterPath(Paths, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)
	fmt.Println("Best paths: ", BestPath)
	fmt.Println("-------------------------------------------")

	//Distribution of the Ants into the room they are going to use
	ant := functions.DistributeAnts(BestPath, RoomStruct.Ants)
	for i := 0; i < len(ant); i++ {
		fmt.Println("Path number:", i+1, "|| Ants in this path:", ant[i])
	}

	//Moving the Ants on the path and printing their movements
	fmt.Println("-------------------------------------------")
	functions.SimulateAntMovement(BestPath, ant)
	fmt.Println("-------------------------------------------")
}

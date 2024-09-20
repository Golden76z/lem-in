package main

import (
	"fmt"
	"lemin/functions"
	"os"
)

func main() {
	//ebiten.SetRunningSlowMessage("")
	if len(os.Args) != 2 {
		fmt.Println("Invalid Input, Usage: go run . [filename]")
		return
	}

	filename := os.Args[1]
	RoomStruct := functions.RoomStruct{}

	correctfile := RoomStruct.CheckLemin(filename)
	if !correctfile {
		return
	}

	RoomMap := make(map[string]*functions.Rooms)
	for i := 0; i < len(RoomStruct.AllRooms); i++ {
		RoomMap[RoomStruct.AllRooms[i].Name] = &RoomStruct.AllRooms[i]
	}

	Paths := functions.FindAllPathsBFS(RoomMap, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)

	BestPath := functions.FilterPath(Paths, RoomStruct.StartingRoom.Name, RoomStruct.EndingRoom.Name)
	fmt.Println("Best paths: ", BestPath)
	fmt.Println("-------------------------------------------")

	antDistribution := functions.DistributeAnts(BestPath, RoomStruct.Ants)
	for i := 0; i < len(antDistribution); i++ {
		fmt.Println("Path number:", i+1, "|| Ants in this path:", antDistribution[i])
	}

	functions.SimulateAntMovement(BestPath, antDistribution)
	fmt.Println("-------------------------------------------")
	fmt.Println("Starting visualization...")

	RunVisualizer(&RoomStruct, BestPath, antDistribution)
}

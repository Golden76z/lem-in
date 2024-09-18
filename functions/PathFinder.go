package functions

import (
	"container/list"
)

func FindAllPathsBFS(rooms map[string]*Rooms, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		// Print the current state of the queue
		// fmt.Println("Current queue:")
		for e := queue.Front(); e != nil; e = e.Next() {
			// fmt.Printf("%v ", e.Value)
		}
		// fmt.Println()

		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]
		// fmt.Println("Processing room:", lastRoom)

		if lastRoom == end {
			paths = append(paths, path)
			continue
		}

		// Explore other rooms as before
		for _, nextRoom := range rooms[lastRoom].Links {
			if !contains(path, nextRoom) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}

	return paths
}
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// func (RoomArray *RoomStruct) PathFinder() {
// 	path := []string{}
// 	count := 0
// 	if len(RoomArray.StartingRoom.Links) == 0 {
// 		return
// 	}
// 	RoomArray.StartingRoom.Visited = true
// 	for i := 0; i < len(RoomArray.StartingRoom.Links); i++ {
// 		index1 := 0
// 		// fmt.Println(RoomArray.AllRooms[i].Name)
// 		// fmt.Println(RoomArray.StartingRoom.Links[i])
// 		// fmt.Println("----------------------------")
// 		for j := 0; j < len(RoomArray.AllRooms); j++ {
// 			if RoomArray.AllRooms[j].Name == RoomArray.StartingRoom.Links[i] {
// 				index1 = j
// 			}
// 		}
// 		currentposition := RoomArray.AllRooms[index1]
// 		fmt.Println("Current position: ", currentposition.Name)
// 		for currentposition.Name != RoomArray.EndingRoom.Name {
// 			currentposition.Visited = true
// 			path = append(path, currentposition.Name)
// 			// RoomArray.AllRooms[i].Visited = true
// 			index := 0
// 			for j := 0; j < len(currentposition.Links); j++ {
// 				for k := 0; k < len(RoomArray.AllRooms); k++ {
// 					if currentposition.Links[j] == RoomArray.AllRooms[k].Name {
// 						index = k
// 					}
// 				}
// 				fmt.Println(RoomArray.AllRooms[index].Visited)
// 				fmt.Println(RoomArray.AllRooms[index].Name)
// 				fmt.Println("--------------------------------")
// 				if !RoomArray.AllRooms[index].Visited {
// 					count++
// 				}
// 			}
// 			fmt.Println(count)
// 			fmt.Println("================================")
// 			currentposition = RoomArray.AllRooms[index]
// 			count = 0
// 			break
// 		}
// 		RoomArray.StartingRoom.Visited = false
// 	}
// 	RoomArray.AllPath = append(RoomArray.AllPath, path)
// }

package functions

import (
	"container/list"
)

// ?Function that will get us all the possible paths with the BFS Algorithm
func FindAllPathsBFS(rooms map[string]*Rooms, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		// Print the current state of the queue
		// fmt.Println("Current queue:")
		// for e := queue.Front(); e != nil; e = e.Next() {
		// 	fmt.Printf("%v ", e.Value)
		// }
		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]

		//Checking if the path is currently on the end position
		if lastRoom == end {
			//If that's the case we append this path and go to the next one
			paths = append(paths, path)
			continue
		}

		//Explore other rooms as before
		for _, nextRoom := range rooms[lastRoom].Links {
			//If the link we are heading isn't contained in our path, then we move
			if !contains(path, nextRoom) {
				//Creating a new path that will have the same lenght as the current path
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
	return paths
}

// ?Function that will check that the path we are heading isn't already used
// ?to avoid getting into an infinite loop
func contains(slice []string, element string) bool {
	for _, current := range slice {
		if current == element {
			//If the element is already there we return true
			return true
		}
	}
	//Otherwise we return false
	return false
}

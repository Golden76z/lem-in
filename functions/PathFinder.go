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

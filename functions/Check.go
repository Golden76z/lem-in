package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ? Function that will all kinds of potentials errors in the file
func (RoomArray *RoomStruct) CheckLemin(filename string) bool {
	//Opening the file
	file, err := os.Open("./examples/" + filename + ".txt")
	if err != nil {
		return false
	}

	//Boolean values to check if we have start and end rooms
	startingroom := false
	endingroom := false
	startcount := 0
	endcount := 0
	first := true

	//Going throught the file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if first {
			numberofants, err := strconv.Atoi(fileScanner.Text())
			if err != nil || numberofants < 1 {
				fmt.Println("------------------------------------")
				fmt.Println("   Error: Invalid number of Ants")
				fmt.Println("------------------------------------")
				return false
			} else {
				RoomArray.Ants = numberofants
			}
			first = false
		}
		//Creation of 2 arrays that will store coordinates and links values
		temparray := strings.Split(fileScanner.Text(), " ")
		linkarray := strings.Split(fileScanner.Text(), "-")
		//Checking if the line is a room
		if len(temparray) == 3 && startingroom {
			coordinate_x, err1 := strconv.Atoi(temparray[1])
			coordinate_y, err2 := strconv.Atoi(temparray[2])
			//Checking if the coordinates are int values
			if err1 != nil || err2 != nil {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid coordinates")
				fmt.Println("------------------------------------")
				return false

				//Checking if the room name starts with a F or a # (Error)
			} else if temparray[0][0] == '#' || temparray[0][0] == 'L' {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid room name")
				fmt.Println("------------------------------------")
				return false
			} else {
				//Storing the room values into a struct
				SingleRoom := Rooms{
					Name:    temparray[0],
					x_value: coordinate_x,
					y_value: coordinate_y,
				}
				//Storing that room into the room array
				RoomArray.StartingRoom = SingleRoom
				RoomArray.AllRooms = append(RoomArray.AllRooms, SingleRoom)
				startingroom = false
			}
			//Checking if the line is a link
		} else if len(temparray) == 3 && endingroom {
			coordinate_x, err1 := strconv.Atoi(temparray[1])
			coordinate_y, err2 := strconv.Atoi(temparray[2])
			//Checking if the coordinates are int values
			if err1 != nil || err2 != nil {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid coordinates")
				fmt.Println("------------------------------------")
				return false

				//Checking if the room name starts with a F or a # (Error)
			} else if temparray[0][0] == '#' || temparray[0][0] == 'L' {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid room name")
				fmt.Println("------------------------------------")
				return false
			} else {
				//Storing the room values into a struct
				SingleRoom := Rooms{
					Name:    temparray[0],
					x_value: coordinate_x,
					y_value: coordinate_y,
				}
				//Storing that room into the room array
				RoomArray.EndingRoom = SingleRoom
				RoomArray.AllRooms = append(RoomArray.AllRooms, SingleRoom)
				endingroom = false
			}
			//Checking if the line is a link
		} else if len(temparray) == 3 {
			coordinate_x, err1 := strconv.Atoi(temparray[1])
			coordinate_y, err2 := strconv.Atoi(temparray[2])
			//Checking if the coordinates are int values
			if err1 != nil || err2 != nil {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid coordinates")
				fmt.Println("------------------------------------")
				return false

				//Checking if the room name starts with a F or a # (Error)
			} else if temparray[0][0] == '#' || temparray[0][0] == 'L' {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid room name")
				fmt.Println("------------------------------------")
				return false
			} else {
				//Storing the room values into a struct
				SingleRoom := Rooms{
					Name:    temparray[0],
					x_value: coordinate_x,
					y_value: coordinate_y,
				}
				//Storing that room into the room array
				RoomArray.AllRooms = append(RoomArray.AllRooms, SingleRoom)
			}
			//Checking if the line is a link
		} else if len(linkarray) == 2 {
			//Checking if a room is linked to itself
			if linkarray[0] == linkarray[1] {
				fmt.Println("------------------------------------")
				fmt.Println("Error: Cannot link a room to itself")
				fmt.Println("------------------------------------")
				return false
			}
			exist1 := false
			exist2 := false
			index1 := 0
			index2 := 0
			//Iterating over all room names to check if the link is valid
			for i := 0; i < len(RoomArray.AllRooms); i++ {
				if RoomArray.AllRooms[i].Name == linkarray[0] {
					exist1 = true
					index1 = i
				} else if RoomArray.AllRooms[i].Name == linkarray[1] {
					exist2 = true
					index2 = i
				}
			}
			//If one of the 2 rooms that are linked doesn't exist, display an error
			if !exist1 || !exist2 {
				fmt.Println("------------------------------------")
				fmt.Println("     Error: Invalid Room Name")
				fmt.Println("------------------------------------")
				return false

				//If they both exist, we append to the first the name of the second to his links
				//And we append to the second the name of the first to his links aswell
			} else {
				RoomArray.AllRooms[index1].Links = append(RoomArray.AllRooms[index1].Links, linkarray[1])
				RoomArray.AllRooms[index2].Links = append(RoomArray.AllRooms[index2].Links, linkarray[0])
			}
		}

		//Checking if we got a start
		if fileScanner.Text() == "##start" {
			startcount++
			startingroom = true

			//Checking if we got an end
		} else if fileScanner.Text() == "##end" {
			endcount++
			endingroom = true
		}
	}
	//Checking if the start count isn't different from 1
	if startcount != 1 {
		fmt.Println("------------------------------------")
		fmt.Println("     Error: No starting point")
		fmt.Println("------------------------------------")
		return false

		//Checking if the end count isn't different from 1
	} else if endcount != 1 {
		fmt.Println("------------------------------")
		fmt.Println("   Error: No ending point")
		fmt.Println("------------------------------")
		return false
	}

	//Ranging over all the rooms to check if there's no doubles
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		for j := i + 1; j < len(RoomArray.AllRooms); j++ {
			if RoomArray.AllRooms[i].Name == RoomArray.AllRooms[j].Name {
				fmt.Println("------------------------------")
				fmt.Println(" Error: Rooms are not unique")
				fmt.Println("------------------------------")
				return false
			}
		}
	}

	//Getting the links on the starting and ending rooms
	for i := 0; i < len(RoomArray.AllRooms); i++ {
		if RoomArray.StartingRoom.Name == RoomArray.AllRooms[i].Name {
			RoomArray.StartingRoom.Links = RoomArray.AllRooms[i].Links
		} else if RoomArray.EndingRoom.Name == RoomArray.AllRooms[i].Name {
			RoomArray.EndingRoom.Links = RoomArray.AllRooms[i].Links
		}
	}

	//If we didn't encounter an error until now we return true
	return true
}

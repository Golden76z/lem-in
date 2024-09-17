package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (RoomArray *RoomStruct) CheckLemin(filename string) bool {
	//Opening the file
	file, err := os.Open("./examples/" + filename + ".txt")
	if err != nil {
		return false
	}

	//Boolean values to check if we have start and end rooms
	startcondition := false
	endcondition := false
	first := true

	//Going throught the file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if first {
			numberofants, err := strconv.Atoi(fileScanner.Text())
			if err != nil || numberofants == 0 {
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
		if len(temparray) == 3 {
			coordinate_x, err1 := strconv.Atoi(temparray[1])
			coordinate_y, err2 := strconv.Atoi(temparray[2])
			//Checking if the coordinates are int values
			if err1 != nil || err2 != nil {
				fmt.Println("------------------------------------")
				fmt.Println("    Error: Invalid coordinates")
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
			index := 0
			//Iterating over all room names to check if the link is valid
			for i := 0; i < len(RoomArray.AllRooms); i++ {
				if RoomArray.AllRooms[i].Name == linkarray[0] {
					exist1 = true
					index = i
				} else if RoomArray.AllRooms[i].Name == linkarray[1] {
					exist2 = true
				}
			}
			if !exist1 || !exist2 {
				fmt.Println("------------------------------------")
				fmt.Println("     Error: Invalid Room Name")
				fmt.Println("------------------------------------")
				return false
			} else {
				RoomArray.AllRooms[index].Links = append(RoomArray.AllRooms[index].Links, linkarray[1])
			}
		}
		if fileScanner.Text() == "##start" {
			startcondition = true
		} else if fileScanner.Text() == "##end" {
			endcondition = true
		}
	}
	//Checking if we got a starting point
	if !startcondition {
		fmt.Println("------------------------------------")
		fmt.Println("     Error: No starting point")
		fmt.Println("------------------------------------")
		return false

		//Checking if we got an ending point
	} else if !endcondition {
		fmt.Println("------------------------------")
		fmt.Println("   Error: No ending point")
		fmt.Println("------------------------------")
		return false
	}
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
	return true
}

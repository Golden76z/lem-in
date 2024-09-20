package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// ? Function that will attribute names to our Ants
func (antTab *RoomStruct) NameAnt() {
	antsNumber := antTab.Ants

	for i := 1; i <= antsNumber; i++ {
		antTab.tabAntName = append(antTab.tabAntName, "L"+strconv.Itoa(i))
	}
}

// ? Function that will move our Ants and print the movements
func SimulateAntMovement(paths [][]string, antDistribution [][]int) {
	type AntPosition struct {
		ant  int
		path int
		step int
	}
	count := 0
	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}

	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom

				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
		antPositions = newPositions
		count++
	}
	fmt.Println("---------------------------------------------------------------------------------------")
	println("Number of move: ", count-1)
}

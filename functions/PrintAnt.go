package functions

import (
	"fmt"
	"strconv"
)

func (antTab *RoomStruct) NameAnt() {
	antsNumber := antTab.Ants

	for i := 1; i <= antsNumber; i++ {
		antTab.tabAntName = append(antTab.tabAntName, "L"+strconv.Itoa(i))
	}
	fmt.Println(antTab.tabAntName)
}

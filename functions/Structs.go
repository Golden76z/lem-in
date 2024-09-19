package functions

//Struct that will store a single room informations
type Rooms struct {
	Name    string
	x_value int
	y_value int
	Links   []string
	Visited bool
}

//Struct that will store all informations needed
type RoomStruct struct {
	Ants         int
	tabAntName   []string
	AllRooms     []Rooms
	AllPath      [][]string
	StartingRoom Rooms
	EndingRoom   Rooms
}

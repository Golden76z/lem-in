package functions

type Rooms struct {
	Name    string
	x_value int
	y_value int
	Links   []string
	Visited bool
}

type RoomStruct struct {
	Ants         int
	tabAntName   []string
	AllRooms     []Rooms
	AllPath      [][]string
	StartingRoom Rooms
	EndingRoom   Rooms
}

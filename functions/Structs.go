package functions

type Rooms struct {
	Name     string
	x_value  int
	y_value  int
	Links    []string
	Position int
	Visited  bool
}

// type StartingRoom struct {
// 	Name    string
// 	x_value int
// 	y_value int
// 	Links   []string
// }

// type EndingRoom struct {
// 	Name    string
// 	x_value int
// 	y_value int
// 	Links   []string
// }

type RoomStruct struct {
	Ants         int
	AllRooms     []Rooms
	AllPath      [][]string
	StartingRoom Rooms
	EndingRoom   Rooms
}

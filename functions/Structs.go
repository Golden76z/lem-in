package functions

type Rooms struct {
	Name    string
	x_value int
	y_value int
	Links   []string
}

type RoomStruct struct {
	Ants     int
	AllRooms []Rooms
}

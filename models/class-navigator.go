package models

type RoomNavigator struct {
	Code string
	Name string
	Description string
	Building string
	Floor int
	Long float64
	Lat float64
}

type RoomIndex struct {
	ID int
	Code string
	Name string
	Building string
	Floor int
}
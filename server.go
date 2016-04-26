package main

import (
	"fmt"

	"github.com/go-martini/martini"
)

var m *martini.Martini

func main() {
	m = martini.New()

	r := martini.NewRouter()
	r.Get("/foods", GetAllData)

	m.Action(r.Handle)
	m.Run()
}

//SomeData - An example data type
type SomeData struct {
	ID          int
	Description string
}

//GetAllData - a Sample function that demonstrates getting data and converting it to JSON
func GetAllData() []SomeData {
	fmt.Println("called GetFoods")
	sd1 := SomeData{
		ID:          1,
		Description: "The first sample food",
	}
	sd2 := SomeData{
		ID:          2,
		Description: "The second sample food",
	}
	sdl := []SomeData{sd1, sd2}
	return sdl
}

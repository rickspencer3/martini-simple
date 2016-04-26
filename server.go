package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var m *martini.Martini

func main() {
	m = martini.New()
	m.Use(render.Renderer(render.Options{}))

	r := martini.NewRouter()
	r.Get("/data", GetAllData)

	m.Action(r.Handle)
	m.Run()
}

//SomeData - An example data type
type SomeData struct {
	ID          int
	Description string
}

//GetAllData - a Sample function that demonstrates getting data and converting it to JSON
func GetAllData(r render.Render) {
	sd1 := SomeData{
		ID:          1,
		Description: "The first sample datum",
	}
	sd2 := SomeData{
		ID:          2,
		Description: "The second sample datum",
	}
	sdl := []SomeData{sd1, sd2}
	r.JSON(200, map[string]interface{}{
		"status": "OK",
		"value":  sdl})
}

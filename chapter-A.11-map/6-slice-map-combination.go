package main

import "fmt"

func main() {
	var ducks = []map[string]string{
		{"name": "Duck Blue", "gender": "Male"},
		{"name": "Duck Yellow", "gender": "Female"},
	}

	for _, duck := range ducks {
		fmt.Println(duck["gender"], duck["name"])
	}

	var datas = []map[string]string{
		{"name": "Duck", "gender": "Male", "color": "Blue"},
		{"address": "Garut", "id": "DK01"},
		{"community": "duck lovers"},
	}

	for _, data := range datas {
		fmt.Println(data)
	}
}

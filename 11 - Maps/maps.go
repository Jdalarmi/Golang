package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":      "Jefe",
		"last_name": "Dalarmi",
	}
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first_name": "Kobe",
			"last_name":  "Bryant",
		},
		"course": {
			"name": "software enginner",
		},
	}
	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)
}

package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["white1"] = "#fffff"
	colors["white2"] = "#ffff"
	colors["white3"] = "#fff"
	fmt.Println(colors)

	for key, value := range colors {
		fmt.Println("Key:", key, "Value:", value)
	}

	delete(colors, "white1")

	fmt.Println(colors)

	changeMap(colors)

	fmt.Println(colors)
}

func changeMap(m map[string]string) {
	m["cat"] = "pur"
}

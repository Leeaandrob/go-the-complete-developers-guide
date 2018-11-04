package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "how"}

	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

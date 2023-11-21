package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")
	var arrays [5]int
	arrays[0] = 1
	arrays[1] = 2
	fmt.Println(arrays)

	arrays2 := [5]string{"Guilherme", "Flamenguista", "gostoso"}
	fmt.Println(arrays2)

	arrays3 := [...]int16{1, 2, 3, 7, 6, 8}
	fmt.Println(arrays3)

	slice := []int32{1, 2, 4, 5, 7, 7, 8, 9, 80, 9, 67, 56, 7}
	fmt.Println(slice)

	slice = append(slice, 23)
	fmt.Println(slice)

	slice2 := arrays2[0:2]
	fmt.Println(slice2)

	arrays2[0] = "lucio"
	fmt.Println(slice2)

}

package main

import "fmt"

func main() {

	slice4 := make([]int16, 10, 11)

	slice4 = append(slice4, 5, 8)
	fmt.Println(slice4)
	fmt.Println(slice4)
	fmt.Print(len(slice4))
	fmt.Println(cap(slice4))

}

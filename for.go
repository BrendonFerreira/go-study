package main

import (
	"fmt"
)

func main() {
	x := []int{4,5,6,7,48}
	fmt.Println(x)
	x = append(x, 232, 23)
	fmt.Println(x)
	y := []int{2,32,12,23}
	x = append(x, y...)
	fmt.Println(x)
}



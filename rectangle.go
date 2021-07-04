package main

import "fmt"

func main() {
	var width, height uint
	fmt.Scan(&width, &height)
	perimeter := (width + height) * 2
	area := width * height
	fmt.Println("Периметр:", perimeter)
	fmt.Println("Площадь:", area)
}

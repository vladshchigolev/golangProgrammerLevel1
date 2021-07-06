package main

import (
	"fmt"
)

func main()  {
	//var userString string
	//fmt.Scan(&userString)
	//
	//if strings.Contains(userString, "чат") {
	//	fmt.Print("БОТ\n")
	//} else {
	//	fmt.Print("НЕ БОТ\n")
	//}
	//typeOfResult := reflect.TypeOf(strings.Contains(userString, "чат"))
	//fmt.Print(reflect.TypeOf(typeOfResult))
	var i int64 = 2
	for var i = 0; i < 25; i += 2{
		fmt.Println(i)
		i += 2
	}
	fmt.Println("i:", i)
}

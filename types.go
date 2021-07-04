package main

import (
	"fmt"
	"unsafe"
)

func main()  {
//	Boolean (zero-value = false)

//	Numeric (Integer)
	var intA int64 = 92
	var intB int32 = 12
	intC := intA + int64(intB) // Когда существует несоответствие типов, приводить рекомендуется менее вместительный тип к более вместительному, обратное может повлечь потерю значения
	//fmt.Println(reflect.TypeOf(a))
	latinStr := "abcd"
	cyrillicStr := "абвг"
	fmt.Println(len(latinStr), len(cyrillicStr)) // Строка в Go - последовательность байтов, функция len() возвращает их количество в переданной строке, 1 latin symbol = 1 byte, 1 cyrillic symbol = 2 byte
	fmt.Printf("Size of latin string variable: %d. Size of cyrillic string variable: %d\n", unsafe.Sizeof(latinStr), unsafe.Sizeof(cyrillicStr))
	fmt.Printf("Value of c is %d, %d bytes long\n", intC, unsafe.Sizeof(intC))
//	Важно! Если проводятся арифметические операции над int и intX, то ОБЯЗАТЕЛЬНО нужно использовать механизм приведения, так как int != int32/64 (в зависимости от платформы), это не alias к ним! Это другая сущность!
	var intD int = 56
	var intE int64 = 39
//	Хотя, казалось бы, типы d и e одного размера (int64 = 64 бита, int (в нашем случае) = 64 бита), они не являются одним и тем же, это разные сущности! Следовательно:
	intF := int64(intD) + intE // Можно так
	intG := intD + int(intE) // А можно даже и так, главное чтобы ТИПЫ БЫЛИ ОДИНАКОВЫЕ
	fmt.Println(intF, intG)

//	Numeric (Floats)
//	float32 и float64, ПРОСТО float НЕТ!
	var userResponse string
	fmt.Print("Запустить бесконечный цикл? ")
	fmt.Scan(&userResponse)
	//if strings.Compare(userResponse, "да") == 0 || strings.Compare(userResponse, "Да") == 0 || strings.Compare(userResponse, "ДА") == 0 {
	//	i := 0
	//	for i < 10 {
	//		fmt.Print(i)
	//	}
	//} else if strings.Compare(userResponse, "нет") == 0 || strings.Compare(userResponse, "Нет") == 0 || strings.Compare(userResponse, "НЕТ") == 0 {
	//	fmt.Println("Ок, как хочешь :)")
	//} else {
	//	fmt.Println("Я не понял твой ответ :(")
	//}
	compareResult := userResponse == "Да"
	fmt.Println(compareResult)
}

package main

import "fmt"

func main() {
	number := 0
	bit := 0
	val := 0
	fmt.Println("Введите число:")
	fmt.Scan(&number)
	fmt.Println("Введите номер бита:")
	fmt.Scan(&bit)
	fmt.Println("Введите значение:")
	fmt.Scan(&val)
	fmt.Println("Исходное число:")
	fmt.Printf("%b\n", number)
	if val == 1 {
		number = number | (1 << bit)
	} else {
		number = number & ^(1 << bit)
	}
	fmt.Println("Новое число:")
	fmt.Printf("%b\n", number)
}

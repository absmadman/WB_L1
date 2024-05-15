package main

var justString string // глобальная переменная

func someFunc() {
	v := createHugeString(1 << 10) // возможно не хватит памяти
	justString = v[:100]           // не можем быть уверены что в строке точно есть 100 элементов
}

func main() {
	someFunc()
}

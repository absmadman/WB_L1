package main

import "fmt"

type Human struct {
	a int
}

type Action struct {
	Human
	b int
}

// Конструктор структуры Human
func NewHuman(value int) *Human {
	return &Human{
		value,
	}
}

// Конструктор структуры Action
func NewAction(value int, human *Human) *Action {
	return &Action{
		*human,
		value,
	}
}

// Методы структуры Human
func (h Human) Speak() {
	fmt.Println("i am human")
}

func (h Human) Display() {
	fmt.Println(h.a)
}

func main() {
	human := NewHuman(1)
	human.Speak()
	human.Display()

	// При создании нового экземпляра структуры action от существующего Human мы ожидаем что полям Human останутся такими же
	// А также для экземлпяра action структуры Action можно будет вызвать методы структуры Human
	action := NewAction(2, human)
	action.Speak()
	action.Display()

	// Как видно из результата наследование выполнено успешно
}

package main

import "fmt"

type cat struct {
}

type adapter struct {
	*cat
}

func (a *adapter) PrintForAdapter() {
	a.Print()
}

func (c *cat) Print() {
	fmt.Println("meow")
}

func main() {
	c := cat{}
	c.Print()
	a := adapter{}
	a.PrintForAdapter()
}

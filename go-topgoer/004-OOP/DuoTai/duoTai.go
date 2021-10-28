// 多态：可以简单理解为以接口作为形参的函数
package main

import "fmt"

type Eater interface {
	Eat()
}

type Man struct {
}

type Woman struct {
}

func (man *Man) Eat() {
	fmt.Println("Man Eat")
}

func (woman *Woman) Eat() {
	fmt.Println("woman Eat")
}

func main() {
	var e Eater

	woman := Woman{}
	man := Man{}

	e = &woman
	e.Eat()

	e = &man
	e.Eat()
}

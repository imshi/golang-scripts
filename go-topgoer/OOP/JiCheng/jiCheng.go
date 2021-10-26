// 继承：通过结构体中的匿名字段来实现，子类继承父类的全部功能
// GO语言的继承方式采用的是匿名组合的方式
package main

import "fmt"

type Person struct {
	name string
}

type Woman struct {
	Person
	sex string
}

func main() {
	woman := Woman{Person{"wangwu"}, "女"}
	fmt.Println(woman.name)
	fmt.Println(woman.sex)
}

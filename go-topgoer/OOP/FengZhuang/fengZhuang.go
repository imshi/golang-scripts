// 封装：调用方法时，在无需知道方法内部的实现细节的前提下直接操作对象内部的数据；
package main

import "fmt"

type Person struct {
	name string
}

func (person *Person) setName(name string) {
	person.name = name
}

func (person *Person) GetInfo() {
	fmt.Println(person.name)
}

func main() {
	p := &Person{"zhangsan"}
	p.setName("lisi")
	p.GetInfo()
}

// Go 的结构体类似于其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等；
// 表示一项记录及该记录的各项属性/成员，通常选择用大写字母开头的成员名称以便于JSON导出编码；
// 结构体支持嵌套

package main

import "fmt"

// 这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

func main() {

	// 使用这个语法创建了一个新的结构体元素。
	fmt.Println(person{"Bob", 20})

	// 可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})

	// & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用点来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以对结构体指针使用 . - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变的。
	sp.age = 51
	fmt.Println(sp.age)

	// 结构体嵌套
	type base struct {
		num int
	}

	// container 结构体嵌套 base 结构体
	type container struct {
		base
		str string
	}

	co := container{
		base: base{num: 1},
		str:  "some thing",
	}

	fmt.Printf("co={num: %v, str:%v} \n", co.num, co.str)
	fmt.Println("num can also get by:", co.base.num)
}

// 临时测试文件
package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// ppp 类型判断
func ppp(s string) {
	s1 := s[0]
	// %T类型输出
	fmt.Printf("s1 type:%T\n", s1)
	// 使用反射
	fmt.Println("s1 type:", reflect.TypeOf(s1))
}

// ccc 字符串拆分
func ccc(s string) {
	arr := strings.Split(s, ".")
	fmt.Printf("slice len :%d \n", len(s))
	fmt.Printf("The first unit8 of slice is: %s \n", arr[0])
}

// slice 拷贝
func sliceCopy() {
	a := []string{"a", "b", "c"}
	// b := []string{"d", "e"}
	c := []string{"f", "g", "h", "i", "j"}
	copy(a[1:], c[:])
	fmt.Println(a)
	copy(c[1:], c[2:])
	fmt.Println(c)
}

// 编码
func sliceCode() {
	// single := '\ufe35'
	single := rune('1')
	fmt.Printf("rune等价于%T ; single的码点为：%v \n", single, single)
	fmt.Println(unicode.IsNumber(single))

	s := []byte("abc     a aaa     ccc  ddd d")
	fmt.Printf("[] byte's len: %d\n", len(s))
	// fmt.Printf()
	tORf := unicode.IsSpace(rune(s[1])) //空格判断
	fmt.Printf("is space or not: %t\n", tORf)
}

type Point struct {
	x, y int
}

// 结构体作为函数参数或返回值使用，如果考虑效率的话，结构体通常会用指针的方式传入和返回
func Scale(p *Point, factor int) Point {
	return Point{p.x * factor, p.y * factor}
}

// 返回url中最后一个"/"符号的索引
func stringIndex() {
	url := "http://www.omdbapi.com/?i=tt3896198"
	pos := strings.LastIndex(url, "/")
	fmt.Println(pos)
}

// const常量
func const0711() {
	x := 100
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Print(y)
	const (
		a = iota
		b
		c = "haha"
		d
		e = iota
	)
	fmt.Println(a, b, c, d, e)
	// 数字常量不会分配存储空间，无法寻址，故下一行的写法会报错
	// fmt.Println(&e)
}

////////////////
// switch条件语句 //
////////////////
func switch0711() {
	// 局部变量
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Println("优秀！")
	case grade == "B", grade == "C":
		fmt.Println("良好！")
	case grade == "D":
		fmt.Printf("及格！")
	case grade == "F":
		fmt.Println("不及格！")
	default:
		fmt.Println("差！")
	}
	fmt.Printf("你的等级是：%s\n", grade)
}

// select随机选择case，都不为true执行default，没有default的话阻塞
func select0713() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")

	}
}

func goto0713() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代，直接跳到指定行，这里等价与continue */
			a = a + 1
			// continue
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

// 打印9*9乘法表
func nineWithNine() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	// ppp("123abcdef")
	// ccc("tom.li")
	// sliceCopy()
	// sliceCode()

	// pp := &Point{1, 2}
	// fmt.Println(Scale(pp, 5))
	// stringIndex()
	// const0711()
	// switch0711()
	// select0713()
	// goto0713()
	nineWithNine()
}

// 临时测试文件
package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"
)

// ppp 类型判断
func ppp(s string) {
	s1 := s[0]
	// %T类型输出
	fmt.Printf("s1 type: %T\n", s1)
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
		fmt.Printf("received %d from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d to c2\n", i2)
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received %d from c3\n", i3)
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

// 字符串截取
func stringCut() {
	mapDemo := []int{1, 22, 333, 444}
	fmt.Println(fmt.Sprint(mapDemo))
	fmt.Println(strings.Trim(fmt.Sprint(mapDemo), "[]"))
	fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(mapDemo), "[]"), " ", ",", -1))
}

///////////
// 结构体指针 //
///////////
func point0812() {
	type name int8
	type first struct {
		a int
		b bool
		name
	}
	var a = first{1, false, 2}
	var b *first = &a
	fmt.Println(a.b, a.a, a.name, &a, a, b.a, &b, (*b).a)
	fmt.Printf("a's type:%T \n", &a)
	// 使用new关键字初始化结构体类型的指针
	c := new(first)
	fmt.Printf("c's type:%T, c's value:%v \n", c, c)

	// 操作指针改变变量的值
	d := 255
	e := &d
	d = d + 1 //256
	d++       //257
	*e++      //258
	fmt.Println(d, e)

}

// 使用指针传递函数参数
func change(val *int) {
	*val = 55
}
func point0813() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Printf("Type of b: %T \n", b)
	fmt.Printf("value of b: %p \n", b)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println("value of a after function call is", a)
}
func modify(sls []int) {
	sls[0] = 90
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func point0814() {
	a := [3]int{89, 90, 91}
	// 使用切片为函数传参并修改，是使用函数列表指针参数的替代方案
	modify(a[:])
	fmt.Println(a)
	var c int = 100
	var d int = 200
	swap(&c, &d)
	fmt.Printf("交换后的c:%d, d:%d \n", c, d)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type Employee struct {
	name     string
	salary   int
	currency string
}

// 结构体类型方法
func (e *Employee) displaySalary() {
	fmt.Printf("%s's salary is %s%d\n", e.name, e.currency, e.salary)
	e.salary = 6000
	fmt.Printf("%s's salary now is %s%d\n", e.name, e.currency, e.salary)
}

func printPointStruct(book *Books) {
	fmt.Printf("Book title:%s \n", book.title)
	fmt.Printf("Book author:%s \n", book.author)
	fmt.Printf("Book subject:%s \n", book.subject)
	fmt.Printf("Book id:%d \n", book.book_id)
}

func printStruct(book Books) {
	fmt.Printf("Book title:%s \n", book.title)
	fmt.Printf("Book author:%s \n", book.author)
	fmt.Printf("Book subject:%s \n", book.subject)
	fmt.Printf("Book id:%d \n", book.book_id)
}
func struct0814() {
	var Book1 Books
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	// 结构体指针传参(比较节省开销)
	printPointStruct(&Book1)
	// 结构体传参
	printStruct(Book1)
	// 结构体方法
	emp1 := Employee{
		name:     "Sam",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()
}

/////////////////////////////
// interface接口为函数传参 - 模拟多态 //
/////////////////////////////

// 定义接口
type sharp interface {
	area() float64
}

// 定义结构体 - 对象
type sqrt struct {
	l float64
}
type circle struct {
	r float64
}

// 定义结构体方法
func (s *sqrt) area() float64 {
	return s.l * s.l
}
func (c *circle) area() float64 {
	return c.r * c.r * 3.14
}

// 接口类型作为参数
func getArea(s sharp) {
	fmt.Println(s.area())
}

func interface0819() {
	s1 := &sqrt{6.5}
	c1 := &circle{2.5}
	getArea(s1)
	getArea(c1)
}

/////////////////////////////////
// 使用空接口实现各种类型的对象存储和接收任意类型作为参数 //
/////////////////////////////////
type Dog struct {
	age int
}

type Cat struct {
	weigh float64
}

type Animal1 interface {
}

// 使用空接口，接收任意类型作为参数
func info(v interface{}) {
	fmt.Println(v)
}

func interface0819_2() {
	// 使用空接口，可以实现存储各种类型的对象
	d1 := Dog{1}
	d2 := Dog{2}
	c1 := Cat{3.2}
	c2 := Cat{3.5}

	animals := []Animal1{d1, d2, c1, c2, "c3"}
	fmt.Println(animals)

	info(d1)
	info(c1)
	info("aaa")
	info(100)
}

//////////////////////////////
// 使用内建fmt.Errorf函数返回错误详细信息 //
//////////////////////////////
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("圆面积计算失败，半径：%0.2f 小于0", radius)
	}
	return math.Pi * radius * radius, nil
}

func error0821() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("圆面积为：%0.2f", area)
}

//////////////////////////////////////////////
// 并发：使用sync.WaitGroup保证gorountine执行完毕后继续执行 //
//////////////////////////////////////////////
func receiveMsg0821() {
	msg := make(chan string)
	// 创建一个新的sync.WaitGroup实例wg
	var wg sync.WaitGroup
	// 当需要使用goroutine的时候，调用wg.Add(1) （使用一次调用一次，如果知道到有N个goroutine，可以直接设置对应的N个）
	wg.Add(3)
	go func() {
		// 当goroutine执行完毕前，告诉WaitGroup执行完毕：调用对应代码defer wg.Done()
		defer wg.Done()
		time.Sleep(time.Second * 3)
		msg <- "goroutine 1"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		msg <- "goroutine 2"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		msg <- "goroutine 3"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		for i := range msg {
			fmt.Println("message :", i)
		}
	}()
	// 在需要等待所有goroutine执行完毕时，调用代码wg.Wait()
	wg.Wait()

}

//////////
// 通道遍历 //
//////////
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func channelRange0821() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func main() {
	ppp("123abcdef")
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
	// nineWithNine()
	// stringCut()
	// point0812()
	// point0813()
	// point0814()
	// struct0814()
	// interface0819()
	// interface0819_2()
	// error0821()
	// receiveMsg0821()
	// channelRange0821()
}

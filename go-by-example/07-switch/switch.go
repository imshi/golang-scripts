/*
select 和 switch 均用于分支操作，区别在于 switch 可以用于各种类型，select 只能应用于IO操作/channel的发送或者接收；
没有 default 的 select 会阻塞，没有可运行的 case 将会阻塞（死锁）；而 switch 的 default 语句可以省略，如果没有匹配的 case 就退出啥也不干并退出 switch。
select 可以有多个分支都满足条件（会随机的选取一个执行），如果没有匹配则会阻塞（可以通过 default 子句的 select 来实现非阻塞 的发送、接收）；
switch 分支是顺序执行的，从上至下逐一测试，直到匹配成功为止，匹配项后面无需加 break，自动跳出，使用 fallthrough 可以强制继续执行。
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// session 1：根据变量值 switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("true")
	}

	// session 2：多条件 switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// session 3：根据条件 switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

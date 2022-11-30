// Go 支持通过基于描述模板的时间格式化和解析
// 格式化（time.Format）：将给定的time.Time格式化为字符串
// 解析（time.Parse）：将字符串格式的时间转为 time.Time

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 一个基本的按照 RFC3339 进行格式化的例子
	// time.Format：将给定的time.Time格式化为字符串
	t := time.Now()
	p("Now: ", t.Format(time.RFC3339))

	// time.Parse：将字符串格式的时间转为 time.Time
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p("t1:", t1)

	// Format 和 Parse 使用基于例子的形式来决定日期格式：
	// 只认 2006-01-02 15:04:05（Mon Jan 2 15:04:05 MST 2006）这个常量时间 - 据说是 go 的诞生时间
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p("t2:", t2)

	// 对于纯数字表示的时间，也可以使用标准的格式化字符串来解析
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse 函数在输入的时间格式不正确是会返回一个错误
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}

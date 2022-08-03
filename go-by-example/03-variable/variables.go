// 变量需显式声明，需先定义后使用，定义后不使用也会报错，不能二次声明，
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	// 可以申明一次性声明多个变量
	var b, c uint = 1, 123
	fmt.Println(b, c)

	// 支持自动推断变量类型
	var d = true
	fmt.Println(d)

	// 声明变量未指定初始值时，将会初始化为制定类型的零值
	var e int
	fmt.Println(e)

	// 支持短声明：只能在函数体内使用，且不能用于已声明变量赋值
	f := "short"
	fmt.Println(f)

	// 浮点型
	var g float64 = 12.2
	fmt.Println(g)

	// 获取变量类型
	h := 13.3
	fmt.Printf("variable h's type: %T\n", h)

	// Go使用UTF8编码，一个英文字符占 1 byte；一个中文字符占3 byte
	str2 := "go语言"

	// 字符串是以 byte 数组形式保存的，类型是 uint8，占1个 byte；
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println("len(str2)：", len(str2))        // len(str2)：8

	// 如果字符串中包括中文并且要进行截取处理，正确的处理方式是将 string 转为 rune 数组
	runeArr := []rune(str2)

	// []rune 类型中使用int32表示每个字符，可以正确的处理中文
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) //int32

	// 打印时需要用 string 进行类型转换，否则打印的是编码值
	fmt.Println(runeArr[2], string(runeArr[2])) //35821 语
	fmt.Println("len(runeArr): ", len(runeArr)) // len(runeArr): 4

}

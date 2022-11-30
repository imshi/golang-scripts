//  regexp包：内置正则表达式功能

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 测试一个字符串是否符合一个正则表达式，返回布尔值
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 若一个正则表达式要重复多次使用，可以预先编译正则结构体，使用的时候直接匹配
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 多种预编译结构体方法
	// 匹配测试，，返回布尔值
	fmt.Println(r.MatchString("peach"))

	// 查找字符串，匹配到返回匹配的内容，未匹配到返回空值
	fmt.Println(r.FindString("peach punch"))

	// 查找第一次匹配的字符串，返回匹配内容开始和结束位置索引
	fmt.Println(r.FindStringIndex("peach punch"))

	// Submatch 返回完全匹配和局部匹配的字符串，这里会返回 p([a-z]+)ch 和 ([a-z]+) 的 peach 和 ea
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 返回全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带 All 的这个函数返回所有的匹配项，-1 表示不限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 带 All 的 Submatch 返回完全匹配和局部匹配的字符串的索引位置，-1 表示不限制匹配次数
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 提供一个正整数来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// MatchString 的等价用法： []byte参数 + 不带 String 的函数
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile： Compile 的变体，Compile 返回两个值，不能用于常量，
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("MustCompile:", r)

	// 替换部分字符串为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}

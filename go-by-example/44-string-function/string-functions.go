// 标准库的 strings 包提供了很多有用的字符串相关的函数
package main

import (
	"fmt"
	s "strings"
)

// 给 fmt.Println 一个短名字的别名，我们随后将会经常用到

var p = fmt.Println

// 一些 strings 中的函数例子（都是包中的函数，不是字符串对象自身的方法），这意味着在调用时需要传递字符作为第一个参数进行传递

func main() {
	// 字符串包含指定字符
	p("Contains:  ", s.Contains("test", "es"))
	// 字符串中指定字符的个数
	p("Count:     ", s.Count("test", "t"))
	// 前缀校验
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	// 后缀校验
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	// 字符串中指定字符首次出现的索引位置，没有的话返回-1
	p("Index:     ", s.Index("test", "e"))
	// 字符串拼接，比直接使用 + 拼接开销更低
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	// 指定次数重复
	p("Repeat:    ", s.Repeat("a", 5))
	// 指定字符全部替换
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	// 指定字符替换1次
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	// 字符串拆分：使用指定分隔符拆分字符串，返回一个string类型的列表
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	// 转化为小写
	p("ToLower:   ", s.ToLower("TEST"))
	// 转化为大写
	p("ToUpper:   ", s.ToUpper("test"))
	// 输出空白行
	p()
	// 以下这两个不是strings中的一部分，仍然值得一提
	// 获取字符串长度
	p("Len: ", len("hello"))
	// 通过索引获取一个字符（返回ASCII码值）
	p("Char:", "hello"[1])
}

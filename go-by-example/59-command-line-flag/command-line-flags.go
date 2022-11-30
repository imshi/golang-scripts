// 命令行标志：用于给命令行程序指定选项参数，如 wc -l中的 -l
// Go 提供了一个flag包用以支持命令行标志解析，第三方包 spf13/pflag 被设计用来替代标准库中的 flag 包，功能更强，兼容性更好；如果是项目，可使用三方包 spf13/viper：是一款强大的配置解析及读取工具

package main

import (
	"flag"
	"fmt"
)

/*
1、测试该程序，最好将这个程序编译成二进制文件，然后再运行这个程序
2、windows测试示例：
2.1 编译：go build .\go-by-example\59-command-line-flag\command-line-flags.go
2.2.1：省略的标志会自动设置为默认值：./command-line-flags.exe -word=opt
2.2.2：位置参数可以出现在任何标志后面：./command-line-flags.exe -word=opt a1 a2 a3
2.2.3：所有的标志必须位于位置参数之前（否则，这个标志将会被解析为位置参数）：./command-line-flags.exe -word=opt a1 a2 a3 -numb=7
2.2.4：使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本：./command-line-flags.exe -h
2.2.5：如果输入了一个程序未指定的标志，会输出一个错误信息，并再次显示帮助文本：./command-line-flags.exe -wat
*/
func main() {
	// 基本的标记声明仅支持：字符串、整数、布尔值
	// 声明一个名为 word、默认值为 foo、描述为 a string 的标志
	// flag.String 函数返回一个字符串指针（不是一个字符串值）
	wordPtr := flag.String("word", "foo", "a string")

	// 相同的方式声明俩标志，分别叫 numb 和 fork
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 也可以用程序中已有的参数来声明标志（注意在声明函数中需要使用该参数的指针）
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有的标志声明完成后，调用flag.Parse()来执行命令行解析
	flag.Parse()

	// 输出解析的选项（需要使用类似 *wordPtr语法对指针进行解引用，从而获取选项的真实值）以及后面的位置参数
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	// 位置参数（命令行中所有独立的单词，类似于shell中的 $1到$9）
	fmt.Println("tail:", flag.Args())
}

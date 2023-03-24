/*
go 本身并没有良好的错误处理机制，err != nil 的处理方式有两个问题：1、没有错误发生时的上下文信息；2、在层层的错误传递过程中，有可能已经将原始错误转化，丢失了最原始的 error。

第三方错误处理包：github.com/pkg/errors，主要使用三个方法：
1、Wrap：Wrap 返回一个错误，该错误在调用 Wrap 的点处带有堆栈跟踪的 err 注释，并提供了消息。如果 err 为 nil，则 Wrap 返回 nil；
2、Cause：返回错误的根本原因。递归拿到最里层的 error, 用于和 error 常量比较或类型断言成自定义 struct type；
3、WithMessage: WithMessage 用新消息注释 err，如果 err 为 nil，则 WithMessage 返回 nil。
*/

package main

import (
	e "errors"
	"fmt"

	"github.com/pkg/errors"
)

func pkgError01() {
	oldErr := e.New("这是底层 error")

	// 添加错误信息和堆栈信息，生成一个新的 error
	wrappedErr := errors.Wrap(oldErr, "这是封装 error")
	// 输出错误信息
	fmt.Printf("%v\n", wrappedErr)
	fmt.Println("------------pkgError01------------")

	// 输出错误信息和详细堆栈信息
	fmt.Printf("%+v\n", wrappedErr)
	fmt.Println("================================")

	//  获取并输出底层错误
	err := errors.Cause(wrappedErr)
	fmt.Println(err)
}

func pkgError02() {
	e := errors.New("make a error by errors.New")
	e = errors.WithMessage(e, "add message001")
	e = errors.WithMessage(e, "add message002")
	e = errors.WithStack(e)
	fmt.Printf("%+v", e)
}

func main() {
	pkgError01()
	fmt.Println("------------pkgError02------------")
	pkgError02()
}

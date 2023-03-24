/* Go 语言使用一个独立的、明确的返回值来传递错误信息
可以通过 errorw.New() 返回自定义的错误
推荐使用第三方错误处理包：github.com/pkg/errors
*/

package main

import (
	"errors"
	"fmt"
)

// 定义支持错误捕获的函数
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// 自定义错误类型、定制输出格式：struct + method
type argError struct {
	arg  int
	prob string
}

// 实现 Error 方法，自定义输出格式
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 按照自定义的错误类型格式输出错误
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it'"}
	}
	return arg + 3, nil
}

func main() {

	// 输出格式一
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// 输出格式二
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 要在代码中使用自定义错误类型，需要通过类型断言来得到这个错误类型的实例
	_, e := f2(42)
	// 断言 e 是不是 *argError 类型
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

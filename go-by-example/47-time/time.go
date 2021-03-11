// time包：内置包，提供时间相关操作
package main

import (
	"fmt"
	"time"
)

func main() {
	// 设置换行打印别名以便快捷调用
	p := fmt.Println

	now := time.Now()
	p(now)
	// fmt.Println(now)

	then := time.Date(2021, 03, 11, 10, 03, 20, 651387237, time.Local)

	p("自定义时间：", then)

	// 提取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出是星期几
	p(then.Weekday())
	// 两个时间的比较（精确到秒），返回布尔值
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub方法返回一个 Duration 来表示两个时间点的间隔时间
	diff := now.Sub(then)
	p(diff)
	// 计算不同单位下的时间间隔数值
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))
}

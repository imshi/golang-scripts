// 时间戳：Go 使用带 Unix 或者 UnixNano 的 time.Now 获取 Unix时间
// Unix系统起始时间：（GMT）1970年1月1日0时0分0秒
package main

import (
	"fmt"
	"time"
)

// 使用带 Unix 或者 UnixNano 的 time.Now来获取从系统起始时间起到现在的秒数或者纳秒数
func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	fmt.Println(now)

	// 没有毫秒（UnixMillis）的语法，要得到毫秒数的话，需要自己手动从纳秒转化
	millis := nanos / 1000000

	fmt.Println("自协调世界时起到距今秒：", secs)
	fmt.Println("自协调世界时起到距今毫秒：", millis)
	fmt.Println("自协调世界时起到距今纳秒：", nanos)

	// 将系统起始时间起的整数秒或者纳秒转化为相应的时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

// Go 的 math/rand 包提供了伪随机数生成器（PRNG）

// crypto/rand 包提供用于加密目的的随机数（安全性更高）

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列
	// 返回一个0到100的随机随机的整数 n： 0 <= n <= 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))

	// 返回一个64位浮点数 f，0.0 <= f <= 1.0
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 给定一个变化的种子，用于产生变化的序列
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))

	// r2 和 r3 使用相同的种子生成的随机数生成器，将会产生相同的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))

	// rad := rand.New(rand.NewSource(time.Now().Unix()))
	// for i := 0; i < rad.Intn(9)+1; i++ {
	// 	fmt.Println(rad.Intn(50))
	// }
}

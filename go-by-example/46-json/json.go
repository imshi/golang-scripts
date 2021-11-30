// encoding/json包：内置的 JSON 编解码支持，提供内置或者自定义类型与 JSON 数据之间的转化的功能
// 编码：转化为json对象；解码：将json对象转化为内置或者自定义类型
// 推荐直接使用三方包：jsoniter- 性能更高，且完全兼容原生json包；
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 自定义类型的编码和解码
type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// 基本数据类型到JSON字符串的编码
	bolB, _ := json.Marshal(true)
	fmt.Println("内置布尔类型编码：", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("内置整型编码：", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("内置浮点型编码：", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("内置字符串类型编码：", string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("内置切片类型编码：", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("内置map类型编码：", string(mapB))

	// JSON 包可以自动编码你的自定义类型，仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("自定义struct类型编码：", string(res1B))

	// 通过在结构字段声明标签来自定义编码的 JSON 数据键名称
	res2D := Response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("重定义键名：", string(res2B))

	// 以下为解码：将 JSON 数据转化为 Go 值（变量）
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	// 提供一个存储解码数据的变量
	var dat map[string]interface{}

	// 实际的解码和相关的错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("解码输出：", dat)

	// 为了使用解码 map 中的值（声明时 value为空接口），我们需要将他们进行适当的类型转换。例如这里我们将 num 的值转换成 float64类型
	num := dat["num"].(float64)
	fmt.Println("解码后内容的使用：", num)

	// 访问嵌套的值需要一系列的转化
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("解码后嵌套内容的使用：", str1)

	// 解码 JSON 值到自定义类型：好处就是可以为我们的程序带来额外的类型安全加强，并且消除在访问数据时的类型断言
	str := `{"page":1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("解码json到自定义类型：", res)
	fmt.Println("提取解析到自定义类型后的内容：", res.Fruits[0])

	// 上面的例子中，我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体
	// Encoder 主要负责将结构对象编码成 JSON 数据，我们可以调用 json.NewEncoder(io.Writer) 方法获得一个 Encoder 实例，再调用 Encode()方法将对象编码成JSON
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	// Decoder 主要负责将 JSON 数据解析成结构对象，可以调用 json.NewDecoder(io.Reader) 方法获得一个 Decoder 实例，调用Decode()方法将 JOSN 内容解析为对象类型
}

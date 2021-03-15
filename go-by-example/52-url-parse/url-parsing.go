// URL 提供了一个统一资源定位方式，go里面通过 net/url进行 url解析
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	// 该url包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)

	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 可以使用 RawQuery 函数得到字符串中的 k=v 这种格式的查询参数
	fmt.Println(u.RawQuery)
	//将查询参数解析为一个map，以查询字符串为键，对应值字符串切片为值；通过索引可提取其中的个别字段；
	m, _ := url.ParseQuery(u.RawQuery)

	fmt.Println(m)
	fmt.Println(m["k"][0])

}

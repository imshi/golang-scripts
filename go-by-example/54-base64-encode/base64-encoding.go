// Base64编码（加解密）

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 将要编码的字符串
	str := "abc123!?$*&()'-=@~"

	// StdEncoding这个代表是的标准加密，支持对URL加密：使用URLEncoding
	// 编码需要使用 []byte 类型的参数，所以要将字符串转成此类型
	sEnc := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确，就需要进行错误检查
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 URL 兼容的 base64 格式进行编解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println(uEnc)
	// 解码，也需要进行类型转化，此处转回为string
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

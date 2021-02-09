package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"unsafe"

	"crypto/md5"
	"math/rand"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func test(fieldData interface{}) func(params string) string {
	return func(params string) string {
		data := GetField(fieldData, params)
		return data.(string)
	}
}

//RenderTemplate 模板渲染
func RenderTemplate(templateStr string, fieldData interface{}) string {
	reg := regexp.MustCompile(`\$\{(.*)\}`)
	if reg != nil {
		fmt.Println("解析错误")

		result := reg.ReplaceAllStringFunc(templateStr, test(fieldData))
		fmt.Println(result)
	}

	return templateStr
}

//UpperFirst 首字母大写
func UpperFirst(str string) string {
	if len(str) <= 0 {
		return str
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])

	return string(tmp)
}

//Lowerfirst 首字母小写
func Lowerfirst(str string) string {
	if len(str) <= 0 {
		return str
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToLower(tmp[0])

	return string(tmp)
}

//IsNull 判断字符串是否为空
// func IsNull(str string) bool {
// 	if len(str) <= 0 {
// 		return true
// 	}
// 	if strings.TrimSpace(str) == "" {
// 		return true
// 	}

// 	return false
// }

//Format 数据格式化
func Format(any interface{}) string {
	str := ""
	switch any.(type) {
	case string:
		str = fmt.Sprintf(`"%v"`, any)
	case int:
		str = fmt.Sprintf(`%v`, any)
	default:
		str = fmt.Sprintf(`(%v)`, any)
	}

	return str
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes 随机字符串
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//Md5 md5摘要
func Md5(src string) string {
	sum := md5.Sum([]byte(src))
	return hex.EncodeToString(sum[:])
}

//Encode64 base64编码
func Encode64(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

//Decode64 base64解码
func Decode64(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	return string(bytes), err
}

//Str2Bytes 字符串转字节列表
func Str2Bytes(src string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&src))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

//Bytes2Str 字节列表转字符串
func Bytes2Str(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}

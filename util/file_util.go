package util

import (
	// "io/ioutil"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {

}

//WriteFile 写入文件
func WriteFile(filePath string, content string) error {
	file, frr := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if frr != nil {
		fmt.Println("打开文件失败", filePath, frr)
		return frr
	}

	defer file.Close()

	_, wrr := file.WriteString(content)
	if wrr != nil {
		fmt.Println("写入文件错误", filePath, wrr)
		return wrr
	}

	return file.Sync();
}

//StreamFile 文件流式读取文件
func StreamFile(filePath string, handle func(content string)) error {
	f, err := os.Open(filePath)
	// defer f.Close()
	if err != nil {
		return err
	}

	buffer := bufio.NewReader(f)

	for {
		line, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		content := strings.TrimSpace(Bytes2Str(line))
		handle(content)
	}
	f.Close()
	handle("")
	return nil
}

//WriteFile 写入文件
// func WriteFile(filepath string, content string) {
// 	file, fErr := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

// 	if fErr != nil {
// 		fmt.Println("打开文件失败", filepath, fErr)
// 		return
// 	}

// 	defer file.Close()

// 	_, wErr := file.WriteString(content)

// 	if wErr != nil {
// 		fmt.Println("写入文件失败", filepath, wErr)
// 	}
// }



//FileValid 文件是否有效
func FileValid(file *os.File) bool {
	_, err := file.Seek(0, 0)
	return err == nil
}

//执行检查
//上次请求的文件是否关闭
//当前请求是否

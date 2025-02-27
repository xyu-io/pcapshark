package sharkos

import (
	"io"
	"strings"
)

// 不需要fdTerminal组件
type stdoutOutput struct {
	*SharkOS
	io.Writer
}

func (o stdoutOutput) Size() (int, int) { return 120, 25 }
func (o stdoutOutput) IsTerminal() bool { return false }

// 自定义输出写入
func (o stdoutOutput) Write(p []byte) (n int, err error) {
	// 提取信息部分，不用hex和ascii部分
	if string(p) == ".[#]:" {
		return len(p), nil
	}
	if strings.Contains(string(p), "[#]") {
		str := strings.ReplaceAll(string(p), "[#]", "")
		if o.pcapChan != nil {
			o.pcapChan <- str
		}
		//fmt.Println(string(str)) //控制台输出
	}

	//fmt.Print(string(p))
	return len(p), nil
}

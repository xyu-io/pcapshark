package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"mime/multipart"
	"net/http"
)

func FileCheck(file *multipart.FileHeader, msize int) ([]byte, error) {
	if msize == 0 {
		msize = 4
	}
	fileMaxSize := msize << 20 //4M
	if int(file.Size) > fileMaxSize {
		return nil, errors.New("文件不允许大小于16M")
	}
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	b := make([]byte, file.Size)
	_, err = reader.Read(b)
	if err != nil {
		return nil, err
	}
	contentType := http.DetectContentType(b)
	if contentType != "application/octet-stream" && contentType != "application/vnd.tcpdump.pcap" {
		return nil, errors.New(fmt.Sprintf("%s:%s", "文件格式错误", contentType))
	}

	return b, nil
}

func YamlToJson(yamlStr string) (string, error) {
	var data interface{}
	// 解析YAML字符串到data变量
	err := yaml.Unmarshal([]byte(yamlStr), &data)
	if err != nil {
		return "", err
	}
	// 将解析后的数据转换为JSON字节切片
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	// 将JSON字节切片转换为字符串
	return string(jsonData), nil
}

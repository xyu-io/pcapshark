package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/xyu-io/pcapshark/format/all"
	"github.com/xyu-io/pcapshark/shark/sharkos"
	"github.com/xyu-io/pcapshark/shark/utils"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	//4M
	r.MaxMultipartMemory = 4 << 20

	r.HandleMethodNotAllowed = true
	r.Use(utils.SecurityCORS)

	r.POST("/upload", handleUpload)

	err := r.Run(":8081")
	if err != nil {
		return
	}

	// 退出信号，ctl+c
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		fmt.Println("程序退出")
		time.Sleep(5 * time.Second)
		return
	}
}

func handleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "文件上传失败")
		return
	}
	var resp = make([]string, 0)
	files := form.File["upload"]
	for _, file := range files {
		dataBys, err := utils.FileCheck(file, 4)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		var tmp = make([]byte, 0)
		var buf = bytes.NewBuffer(tmp)
		var out = make(chan string)
		go func() {
			for res := range out {
				buf.WriteString(res)
				buf.WriteByte('\n')
			}
		}()
		err = sharkos.BytesExec(c, dataBys, out)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		json, err := utils.YamlToJson(buf.String())
		if err != nil {
			fmt.Println("yaml to json: ", err)
			resp = append(resp, buf.String())
		} else {
			resp = append(resp, json)
		}
	}
	c.String(http.StatusOK, strings.Join(resp, "\n"))
}

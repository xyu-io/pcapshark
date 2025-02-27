package utils

import (
	"github.com/gin-gonic/gin"
)

func SecurityCORS(c *gin.Context) {
	//method := c.Request.Method
	//if method == "OPTIONS" {
	//	c.Request.Method = "POST"
	//}
	//// 放行所有OPTIONS方法
	//fmt.Println("method: ", c.Request.Method)
	// 获取当前请求的origin; 地址形式: scheme://domain 如  "http://localhost:8000"  注意origin的最后是没有 / 斜杠的
	origin := c.GetHeader("Origin")
	// 如果请求origin在允许的origin之中,则直接将当前请求的origin设置为允许的origin
	c.Header("Access-Control-Allow-Origin", origin)
	// 附带身份凭证的请求, 注意这里如果是true, 则 Access-Control-Allow-Origin 不允许使用 * 通配符
	c.Header("Access-Control-Allow-Credentials", "true")
	// 设定允许的请求方式
	c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,HEAD,PATCH")
	// 暴露头信息
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
	// 处理请求
	c.Next()
}

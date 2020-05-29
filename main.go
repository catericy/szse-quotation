package main

import "github.com/gin-gonic/gin"

//代理网址接口转发
const Host = "http://api.cninfo.com.cn"

//设置全局 CORS 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow", "*")
		c.Next()
	}
}

func main() {
	
	r := gin.Default()
	
	v := r.Group("/")
	
	r.Run() 
}
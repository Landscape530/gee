package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>HELLO Gee HL!</h1>")
	})
	r.GET("/branch", func(c *gee.Context) {
		c.String(http.StatusOK, "branch %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

/*	自我测试命令 for windows:
(1) $ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Content-Type: text/html
Date: Thu, 21 Mar 2024 07:36:24 GMT
Content-Length: 22

<h1>HELLO Gee HL!</h1>

(2) $ curl "http://localhost:9999/hl?name=huanglei"
branch huanglei, you are at /branch

(3) $ curl "http://localhost:9999/login" -X POST -d "username=landscape&password=2001"
{"password":"2001","username":"landscape"}

(4) $ curl "http://localhost:9999/xxxx"
404 NOT FOUND: /xxxx
 day2-router-handle
*/

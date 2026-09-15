package main

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("web/index.html", "web/other.html")

	router.GET("/", func(context *gin.Context) {
		host := context.Request.Host
		if hostname, _, error := net.SplitHostPort(host); error == nil {
			host = hostname
		}
		host = strings.ToLower(host)

		switch host {
		case "web.ip":
			context.HTML(200, "index.html", gin.H{
				"Title":   "Привет",
				"Message": "Это главная страница.",
			})
		case "other.ip":
			context.HTML(200, "other.html", gin.H{
				"Title":   "Другая страница",
				"Message": "Вы открыли other.ip.",
			})
		case "localhost":
			context.HTML(200, "index.html", gin.H{
				"Title":   "Привет",
				"Message": "Главная страница.",
			})
		default:
			if net.ParseIP(host) != nil {
				context.HTML(200, "index.html", gin.H{
					"Title":   "Привет",
					"Message": "Главная страница: " + host,
				})
				return
			}
			context.String(404, "Неизвестный домен: %s", host)
		}
	})

	router.Run(":8080")
}

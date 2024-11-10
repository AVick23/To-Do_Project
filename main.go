package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер
	router := gin.Default()

	// Настройка папок для статических файлов и шаблонов
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Определение маршрута для главной страницы
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Маршруты для различных разделов
	router.GET("/today", func(c *gin.Context) {
		c.HTML(http.StatusOK, "today.html", nil)
	})

	router.GET("/tomorrow", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tomorrow.html", nil)
	})

	router.GET("/week", func(c *gin.Context) {
		c.HTML(http.StatusOK, "week.html", nil)
	})

	// Пример маршрута для проектов
	router.GET("/project/:name", func(c *gin.Context) {
		projectName := c.Param("name")
		c.HTML(http.StatusOK, projectName+".html", nil)
	})

	// Запуск сервера на 127.0.0.1:8080
	router.Run("127.0.0.1:8080")
}

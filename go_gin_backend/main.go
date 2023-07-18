package main

import (
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginData struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/btnTest", func(c *gin.Context) {
		data := map[string]interface{}{
			"a": "value1",
			"b": "value2",
			"c": "value3",
		}
		c.JSON(200, data)
	})

	r.POST("/btnTest", func(c *gin.Context) {
		var data LoginData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		log.Println("ID:", data.ID)
		log.Println("Password:", data.Password)
		responseData := map[string]interface{}{
			"message": "Successfully received the data.",
		}
		c.JSON(200, responseData)
	})

	r.Run(":8080")
}

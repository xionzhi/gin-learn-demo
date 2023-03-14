package main

import (
	utils "apps/utils"
	login "apps/web/login"
	search "apps/web/search"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConfig := "host=localhost user=root password=123456 dbname=root port=5432 sslmode=disable"
	_, ormErr := utils.InitPostgresDB(dbConfig)
	if ormErr != nil {
		panic(ormErr)
	}
	// migrateErr := db.AutoMigrate(&user.User{})
	// if migrateErr != nil {
	// 	return
	// }

	r := gin.Default()

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"code": 1, "message": "hello world!", "data": nil}) })
	r.POST("/login", login.Login)
	r.GET("/search", search.Search)

	r.Run(":8080")
}

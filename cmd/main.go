package main

import (
	"github.com/isravm89w/gin-rest-api-example/pkg/books"
	"github.com/isravm89w/gin-rest-api-example/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	db := db.Init(dbUrl)
	books.RegisterRoutes(r, db)

	r.Run(port)
}

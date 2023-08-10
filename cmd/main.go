package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rodrigocarsin/Gastos/cmd/server/routes"
)

func main() {
	// NO MODIFICAR
	db, err := sql.Open("mysql", "root:Bauti.04@/gastos")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	//docs.SwaggerInfo.Host = "localhost:8080"
	//eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	eng.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
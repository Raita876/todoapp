package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"gorm.io/driver/postgres"

	"github.com/raita876/todoapp/internal/application/services"
	postgres2 "github.com/raita876/todoapp/internal/infrastructure/db/postgres"
	"github.com/raita876/todoapp/internal/interface/api/rest"
	"gorm.io/gorm"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version:   "TODO: version",
		Name:      "TODO: name",
		Usage:     "TODO: usage",
		UsageText: "TODO: usageText",
		Flags:     []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			e := gin.Default()
			e.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "ok",
				})
			})

			dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tokyo"
			port := ":8080"

			// TODO: DB 接続を成功させる
			gormDB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			// if err != nil {
			// 	log.Fatalf("Failed to connect to database: %v", err)
			// }

			taskRepo := postgres2.NewGormTaskRepository(gormDB)

			taskService := services.NewTaskService(taskRepo)

			rest.NewTaskController(e, taskService)

			return e.Run(port)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

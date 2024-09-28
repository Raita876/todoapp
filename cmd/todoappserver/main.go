package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raita876/todoapp/docs"
	"github.com/raita876/todoapp/internal/application/services"
	postgres2 "github.com/raita876/todoapp/internal/infrastructure/db/postgres"
	"github.com/raita876/todoapp/internal/interface/api/rest"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	version string
	name    string
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version:   version,
		Name:      name,
		Usage:     "This is a REST API for task management.",
		UsageText: "todoappserver [OPTION]...",
		Flags:     []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			e := gin.Default()
			e.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "ok",
				})
			})

			// TODO: flag or config file で制御
			dsn := "host=localhost user=postgres password=postgres dbname=todoapp_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"

			// TODO: flag で制御
			port := ":8080"

			gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				return err
			}

			taskRepo := postgres2.NewGormTaskRepository(gormDB)

			taskService := services.NewTaskService(taskRepo)

			rest.NewTaskController(e, taskService)

			docs.SwaggerInfo.BasePath = "/api/v1"
			e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

			return e.Run(port)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

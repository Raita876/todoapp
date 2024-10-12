package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raita876/todoapp/docs"
	"github.com/raita876/todoapp/internal/application/services"
	mysql2 "github.com/raita876/todoapp/internal/infrastructure/db/mysql"
	"github.com/raita876/todoapp/internal/interface/api/rest"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dsn",
				Aliases: []string{"d"},
				Value:   "mysql:mysql@tcp(127.0.0.1:3306)/todoapp_db?charset=utf8mb4&parseTime=True&loc=Local",
				Usage:   "specify data source name",
			},
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Value:   "0.0.0.0:8080",
				Usage:   "",
			},
		},
		Action: func(ctx *cli.Context) error {
			gin.DisableConsoleColor()
			r := gin.Default()

			dsn := ctx.String("dsn")
			addr := ctx.String("addr")

			gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				return err
			}

			taskRepo := mysql2.NewGormTaskRepository(gormDB)

			taskService := services.NewTaskService(taskRepo)

			rest.NewTaskController(r, taskService)

			docs.SwaggerInfo.BasePath = "/api/v1"
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

			return r.Run(addr)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

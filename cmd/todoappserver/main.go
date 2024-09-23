package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
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

			// TODO
			// taskRepo := postgres2.NewGormTaskRepository(gormDB)

			// TODO: service 実装
			// taskService := services.NewTaskService()

			// TODO
			// rest.NewTaskController(e, )

			return e.Run()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

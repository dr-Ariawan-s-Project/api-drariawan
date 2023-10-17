package main

import (
	"log"
	"os"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/app/database"
	"github.com/dr-ariawan-s-project/api-drariawan/app/router"
	migrator "github.com/dr-ariawan-s-project/api-drariawan/migration"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()

	if args := os.Args; len(args) > 1 {
		switch args[1] {
		case "help":
			log.Print(`
Usage:
go run main.go <command>
		
The commands are:
help         Show helper
migrateup    Migrate the DB to the most recent version available
migratedown  Rollback migration

WARNING : please make sure the query in rollback / down file, its may contains DROP TABLE, the data stored will lost
			`)

		case "migrateup":
			migrator.Migrate(*cfg, "up")
		case "migratedown":
			migrator.Migrate(*cfg, "down")
		}
		return
	}

	dbMysql := database.InitDBMysql(cfg)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(dbMysql, e, cfg)

	e.Logger.Fatal(e.Start(":8000"))
}

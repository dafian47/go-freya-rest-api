package main

import (
	"flag"
	cfg "github.com/dafian47/go-freya-rest-api/config"
	"github.com/dafian47/go-freya-rest-api/middleware"
	"github.com/dafian47/go-freya-rest-api/module/event/handler"
	"github.com/dafian47/go-freya-rest-api/module/event/repository"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var config cfg.Config

func init() {

	config = cfg.NewViperConfig()

	defaultPort := config.GetString("server.port")
	defaultDeployType := config.GetString("deploy.type")

	port := flag.String("port", defaultPort, "Specific port [ :5000 ]")
	deployType := flag.String("deploy_type", defaultDeployType, "Specific deploy type [ local, staging, production ]")

	flag.Parse()

	config.SetString("server.port", *port)
	config.SetString("deploy.type", *deployType)
}

func main() {

	var databaseUrl string
	var isDebug bool

	port := config.GetString("server.port")
	deployType := config.GetString("deploy.type")

	if deployType == "local" {
		databaseUrl = config.GetString("deploy.to.local.dsn")
		isDebug = config.GetBool("deploy.to.local.debug")

	} else if deployType == "staging" {
		databaseUrl = config.GetString("deploy.to.staging.dsn")
		isDebug = config.GetBool("deploy.to.staging.debug")

	} else if deployType == "production" {
		databaseUrl = config.GetString("deploy.to.production.dsn")
		isDebug = config.GetBool("deploy.to.production.debug")
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())

	db := cfg.InitDB(databaseUrl, isDebug)

	eventRepo := repository.NewEventRepo(db)
	handler.NewEventHandler(e, eventRepo)

	e.Logger.Fatal(e.Start(port))
}

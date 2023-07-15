package setup

import (
	"fmt"
	"os"

	"github.com/williamroberttv/curriculum-gen-api/src/config"
	"github.com/williamroberttv/curriculum-gen-api/src/database"
	"github.com/williamroberttv/curriculum-gen-api/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var db database.Database

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// get database config

	user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")

	db.AutoMigrateDb = true
	db.Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbhost, user, password, dbname, dbport, dbhost)
	db.DbType = os.Getenv("DB_TYPE")

	// start database
	_, err = db.Connect()

	if err != nil {
		panic(err)
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
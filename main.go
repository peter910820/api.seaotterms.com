package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"api.seaotterms.com/model"
	blogrouter "api.seaotterms.com/router/blog"
	galrouter "api.seaotterms.com/router/gal"
	teachrouter "api.seaotterms.com/router/teach"
)

var (
	// init store(session)
	blogStore = session.New(session.Config{
		Expiration:     7 * 24 * time.Hour,
		CookieSameSite: "None",
		KeyLookup:      "cookie:user-session",
		// CookieHTTPOnly: true,
	})
	// management database connect
	dbs = make(map[string]*gorm.DB)
)

func init() {
	// init logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
}

func main() {
	// init migration
	for i := 0; i <= 2; i++ {
		dbName, db := model.InitDsn(i)
		dbs[dbName] = db
		model.Migration(dbName, dbs[dbName])
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_URL"),
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: true,
	}))
	logrus.Println("Session config:", session.Config{
		CookieSecure:   false,
		CookieSameSite: "None",
	})
	// api route group
	apiGroup := app.Group("/api") // main api route group

	// site route group
	galrouter.GalRouter(apiGroup, dbs)
	blogrouter.BlogRouter(apiGroup, dbs, blogStore)
	teachrouter.TeachRouter(apiGroup, dbs)

	logrus.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PRODUCTION_PORT"))))
}

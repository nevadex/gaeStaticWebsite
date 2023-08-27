package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("GAEStaticWebsite started")

	app := fiber.New(fiber.Config{
		ReadTimeout:           time.Second * 5,
		WriteTimeout:          time.Second * 10,
		IdleTimeout:           time.Minute,
		GETOnly:               true,
		DisableStartupMessage: true,
		AppName:               "GAE Static Website",
	})

	app.Static("/", "./www", fiber.Static{
		Browse: false,
		Index:  "index.html",
		MaxAge: 600,
	})
	if port, ok := os.LookupEnv("PORT"); ok {
		fmt.Println("Listening at :" + port)
		log.Fatal(app.Listen(":" + port))
	} else {
		fmt.Println("Listening at :8000")
		log.Fatal(app.Listen(":8000"))
	}
}

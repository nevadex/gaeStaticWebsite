package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:           time.Second * 1,
		WriteTimeout:          time.Second * 10,
		DisableKeepalive:      true,
		GETOnly:               true,
		DisableStartupMessage: true,
	})

	app.Static("/", "./www", fiber.Static{
		Browse: false,
		Index:  "index.html",
		MaxAge: 60 * 60 * 24 * 7, // seconds
	})
	if port, ok := os.LookupEnv("PORT"); ok {
		log.Fatal(app.Listen(":" + port))
	} else {
		fmt.Println("No port set, defaulting to 3000")
		log.Fatal(app.Listen(":3000"))
	}
}

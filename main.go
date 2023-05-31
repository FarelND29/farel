package main

import (
	"log"

	"github.com/FarelND29/farel/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/FarelND29/farel/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/FarelND29/farel/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server

// @contact.name API Support
// @contact.url https://github.com/FarelND29
// @contact.email fareldazz123@gmail.com

// @host farel-app.herokuapp.com
// @BasePath /
// @schemes http https

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}


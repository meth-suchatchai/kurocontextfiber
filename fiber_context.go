package kuroctxfiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type defaultFiber struct {
	server  *fiber.App
	port    string
	version string
}

func (f *defaultFiber) Server() {
	//TODO implement me
	log.Println("server start port: ", f.port)
	if err := f.server.Listen(fmt.Sprintf("%v", f.port)); err != nil {
		log.Fatalf("error running server: %v\n", err)
	}
}

func New(config Config) KuroFiber {
	cfg := fiber.Config{
		AppName: config.Name,
	}

	df := &defaultFiber{}

	df.port = ":8080"
	if config.Port != "" {
		df.port = ":" + config.Port
	}

	df.server = fiber.New(cfg)

	df.EnableMonitor()
	df.EnableLog()
	df.EnableCors(CorsConfig{})

	return df
}

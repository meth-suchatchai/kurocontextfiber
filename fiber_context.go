package kuroctxfiber

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"os"
	"os/signal"
	"time"
)

var CFG = fiber.Config{}

type defaultFiber struct {
	server  *fiber.App
	port    string
	version string
}

func (f *defaultFiber) Start() {
	log.Println("server start port: ", f.port)
	if err := f.server.Listen(fmt.Sprintf("%v", f.port)); err != nil {
		log.Fatalf("error running server: %v\n", err)
	}
}

func (f *defaultFiber) Stop(timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := f.server.Shutdown(); err != nil {
		log.Fatalf("error shutting down server: %v\n", err)
	}

	select {
	case <-ctx.Done():
		log.Println("service is shutting down.")
	case <-time.After(timeout):
		log.Println("service shutdown timeout.")
	}
}

func New(config Config) KuroFiber {
	CFG.AppName = config.Name

	CFG.Views = html.New("./views", ".html")

	df := &defaultFiber{}
	df.port = ":8080"
	if config.Port != "" {
		df.port = ":" + config.Port
	}

	df.server = fiber.New(CFG)

	df.server.All("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"AppName": CFG.AppName})
	})

	return df
}

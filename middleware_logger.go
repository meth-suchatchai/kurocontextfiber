package kuroctxfiber

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func (f *defaultFiber) EnableLog() {
	f.server.Use(logger.New(logger.ConfigDefault))
}

func (f *defaultFiber) EnableLoggerFile(fileName string) {
	file, err := os.OpenFile("./log/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("enable logger file error:", err)
	}

	f.server.Use(logger.New(logger.Config{
		Output: file,
	}))
}

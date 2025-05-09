package kuroctxfiber

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

var loggerConfig = logger.ConfigDefault

func (f *defaultFiber) EnableLog(config ...logger.Config) {
	if len(config) != 0 {
		loggerConfig = config[0]
	}

	f.server.Use(logger.New(loggerConfig))
}

func (f *defaultFiber) EnableLoggerFile(fileName string) {
	file, err := os.OpenFile("./log/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("enable logger file error:", err)
	}
	defer file.Close()

	loggerConfig.Output = file

	f.server.Use(loggerConfig)
}

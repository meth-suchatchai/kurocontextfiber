package kuroctxfiber

import (
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"time"
)

type KuroFiber interface {
	CreateApiVersion(version string) Group
	Start()
	Stop(timeout time.Duration)

	EnableLog(config ...logger.Config)
	EnableLoggerFile(fileName string)

	EnableCors(cfg CorsConfig)
	EnableMonitor(config ...monitor.Config)
	EnableHealthcheck(config ...healthcheck.Config)
}

package kuroctxfiber

import "time"

type KuroFiber interface {
	CreateApiVersion(version string) Group
	Start()
	Stop(timeout time.Duration)

	EnableLog()
	EnableLoggerFile(fileName string)

	EnableCors(cfg CorsConfig)
	EnableMonitor()
}

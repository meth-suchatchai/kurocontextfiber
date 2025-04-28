package kuroctxfiber

type KuroFiber interface {
	CreateApiVersion(version string) Group
	Server()

	EnableLog()
	EnableLoggerFile(fileName string)

	EnableCors(cfg CorsConfig)
	EnableMonitor()
}

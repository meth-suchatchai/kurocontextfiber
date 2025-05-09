package kuroctxfiber

import "github.com/gofiber/fiber/v2/middleware/monitor"

var monitorConfig = monitor.Config{}

func (f *defaultFiber) EnableMonitor(config ...monitor.Config) {
	/* Get Monitor metrics */
	if len(config) != 0 {
		monitorConfig = config[0]
	}

	f.server.Get("/logs/metrics", monitor.New(monitorConfig))
}

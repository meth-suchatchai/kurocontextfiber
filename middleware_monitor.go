package kuroctxfiber

import "github.com/gofiber/fiber/v2/middleware/monitor"

func (f *defaultFiber) EnableMonitor() {
	/* Get Monitor metrics */
	f.server.Get("/logs/metrics", monitor.New())
}

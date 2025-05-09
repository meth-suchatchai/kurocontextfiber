package kuroctxfiber

import "github.com/gofiber/fiber/v2/middleware/healthcheck"

var healthConfig = healthcheck.Config{}

func (f *defaultFiber) EnableHealthcheck(config ...healthcheck.Config) {
	if len(config) != 0 {
		healthConfig = config[0]
	}

	f.server.Use(healthcheck.New(healthConfig))
}

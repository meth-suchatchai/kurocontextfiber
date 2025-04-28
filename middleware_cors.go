package kuroctxfiber

import "github.com/gofiber/fiber/v2/middleware/cors"

type CorsConfig struct {
	AllowOrigin string `json:"allow_origin"`
}

func (f *defaultFiber) EnableCors(cfg CorsConfig) {
	acceptAllOrigin := "*"
	if cfg.AllowOrigin != "" {
		acceptAllOrigin = cfg.AllowOrigin
	}

	f.server.Use(cors.New(cors.Config{
		AllowOrigins: acceptAllOrigin,
	}))
}

package kuroctxfiber

import "github.com/gofiber/fiber/v2"

type Group interface {
	Group(prefix string, handleFunc ...HandleFunc) Group
	Get(path string, handleFunc ...HandleFunc)
	Post(path string, handleFunc ...HandleFunc)
	Put(path string, handleFunc ...HandleFunc)
	Patch(path string, handleFunc ...HandleFunc)
	Delete(path string, handleFunc ...HandleFunc)

	/* Handler other api case */
	Use(middleware ...HandleFunc)
	Static(prefix, dir string)
}

type defaultGroup struct {
	g fiber.Router
}

func (d *defaultGroup) Group(prefix string, handleFunc ...HandleFunc) Group {
	fh := newHandler(handleFunc...)

	g := d.g.Group(prefix, fh...)

	return &defaultGroup{
		g: g,
	}
}

func newHandler(h ...HandleFunc) []fiber.Handler {
	var handlers []fiber.Handler
	for _, handler := range h {
		hc := func(context *fiber.Ctx) error {
			if context != nil {
				err := handler(NewContext(context))
				if err != nil {
					return err
				}
			}
			return nil
		}
		handlers = append(handlers, hc)
	}

	return handlers
}

func (d *defaultGroup) Get(path string, handleFunc ...HandleFunc) {
	fh := newHandler(handleFunc...)
	d.g.Get(path, fh...)
}

func (d *defaultGroup) Post(path string, handleFunc ...HandleFunc) {
	fh := newHandler(handleFunc...)
	d.g.Post(path, fh...)
}

func (d *defaultGroup) Put(path string, handleFunc ...HandleFunc) {
	fh := newHandler(handleFunc...)
	d.g.Put(path, fh...)
}

func (d *defaultGroup) Patch(path string, handleFunc ...HandleFunc) {
	fh := newHandler(handleFunc...)
	d.g.Patch(path, fh...)
}

func (d *defaultGroup) Delete(path string, handleFunc ...HandleFunc) {
	fh := newHandler(handleFunc...)
	d.g.Delete(path, fh...)
}

func (d *defaultGroup) Use(middleware ...HandleFunc) {
	fh := newHandler(middleware...)
	d.g.Use(fh)
}

func (d *defaultGroup) Static(prefix, dir string) {
	d.g.Static(prefix, dir)
}

// CreateApiVersion default version is /api/v1
func (f *defaultFiber) CreateApiVersion(version string) Group {
	defaultVersion := version
	if version == "" {
		defaultVersion = "/api/v1"
	}
	rg := f.server.Group(defaultVersion)
	return &defaultGroup{g: rg}
}

func (f *defaultFiber) Group(relativePath string, handleFunc ...HandleFunc) Group {
	fg := f.server.Group(relativePath)

	return &defaultGroup{g: fg}
}

package kuroctxfiber

import "github.com/gofiber/fiber/v2"

type Group interface {
	Group() Group
	Get(path string, handleFunc HandleFunc)
	Post(path string, handleFunc HandleFunc)
	Put(path string, handleFunc HandleFunc)
	Patch(path string, handleFunc HandleFunc)
	Delete(path string, handleFunc HandleFunc)

	/* Handler other api case */
	Use(middleware ...HandleFunc)
	Static()
}

type defaultGroup struct {
	fg fiber.Router
}

func (d defaultGroup) Group() Group {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Get(path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Post(path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Put(path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Patch(path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Delete(path string, handleFunc HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Use(middleware ...HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d defaultGroup) Static() {
	//TODO implement me
	panic("implement me")
}

func newHandler(handleFunc HandleFunc) fiber.Handler {
	handler := func(ctx *fiber.Ctx) error {
		return handleFunc(NewContext(ctx))
	}
	return handler
}

// CreateApiVersion default version is /api/v1
func (f *defaultFiber) CreateApiVersion(version string) Group {
	defaultVersion := version
	if version == "" {
		defaultVersion = "/api/v1"
	}
	rg := f.server.Group(defaultVersion)
	return &defaultGroup{fg: rg}
}

func (f *defaultFiber) Group(relativePath string, handleFunc ...HandleFunc) Group {
	fg := f.server.Group(relativePath)

	return &defaultGroup{fg: fg}
}

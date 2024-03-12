package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouterInterface interface {
	AddRoute(path string, controller ControllerInterface, options string)
}

type RouterGroup struct {
	*echo.Group
}

func (rg *RouterGroup) AddRoute(path string, controller ControllerInterface, options string) {
	request, method := split(options)
	handler := handler(controller, method)
	rg.Add(request, path, handlerFunc(handler, controller))
}

type Router struct {
	*echo.Echo
}

func (r *Router) RouterGroup(path string, m ...echo.MiddlewareFunc) RouterGroup {
	return RouterGroup{r.Group(path, m...)}
}

func (r *Router) AddRoute(path string, controller ControllerInterface, options string) {
	request, method := split(options)
	handler := handler(controller, method)
	r.Add(request, path, handlerFunc(handler, controller))
}

func NewRouter() Router {
	e := echo.New()
	router := Router{e}
	router.Use(initController)
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	return router
}

func initController(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Controller{c}
		return next(cc)
	}
}

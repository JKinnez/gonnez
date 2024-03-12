package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControllerInterface interface {
	Prepare() error
	Index() error
	Create() error
	Show() error
	Update() error
	Delete() error
}

type Controller struct {
	echo.Context
}

func (*Controller) Prepare() error {
	return nil
}

func (*Controller) Index() error {
	return echo.NewHTTPError(http.StatusNotImplemented, errNotImplemented)
}

func (*Controller) Create() error {
	return echo.NewHTTPError(http.StatusNotImplemented, errNotImplemented)
}

func (*Controller) Show() error {
	return echo.NewHTTPError(http.StatusNotImplemented, errNotImplemented)
}

func (*Controller) Update() error {
	return echo.NewHTTPError(http.StatusNotImplemented, errNotImplemented)
}

func (*Controller) Delete() error {
	return echo.NewHTTPError(http.StatusNotImplemented, errNotImplemented)
}

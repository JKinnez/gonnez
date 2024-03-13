package web

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	zero = 0
	one  = 1
)

var (
	errNotImplemented = "Method not implemented"
	errMethodNotFound = "Method %s does not exist in controller"
)

func methodNotFound(method string) string {
	return fmt.Sprintf(errMethodNotFound, method)
}

func split(options string) (request string, method string) {
	splited := strings.Split(options, ":")
	return splited[zero], splited[one]
}

func handler(controller ControllerInterface, method string) reflect.Value {
	handler := reflect.ValueOf(controller).MethodByName(method)
	if !handler.IsValid() {
		panic(methodNotFound(method))
	}

	return handler
}

func handlerFunc(handler reflect.Value, controller ControllerInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller.Init(c)
		err := controller.Prepare()
		if err != nil {
			return err
		}

		return result(handler.Call(nil))
	}
}

func result(method []reflect.Value) error {
	if len(method) > zero && !method[zero].IsNil() {
		if err, ok := method[zero].Interface().(error); ok {
			return err
		}
	}
	return nil
}

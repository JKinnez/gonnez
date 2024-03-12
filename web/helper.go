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
	return func(echo.Context) error {
		err := controller.Prepare()
		if err != nil {
			return err
		}
		// revive:disable:unchecked-type-assertion
		return handler.Call(nil)[0].Interface().(error)
		// revive:enable:unchecked-type-assertion
	}
}

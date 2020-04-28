package rpi

import "github.com/labstack/echo"

// Logger represents a logging interface including arguments source, msg, error, params.
type Logger interface {
	Log(echo.Context, string, string, error, map[string]interface{})
}

package middleware

import (
	"github.com/labstack/echo/v4"
	mwEcho "github.com/labstack/echo/v4/middleware"
)

const ctxKey = "dumpBody"

type dumpBody struct {
	reqBody []byte
	resBody []byte
}

var DumpBody = mwEcho.BodyDumpWithConfig(mwEcho.BodyDumpConfig{
	Skipper: func(c echo.Context) bool {
		if c.Request().Header.Get("Content-Type") != "application/json" {
			return true
		}
		if c.Request().Method == "GET" || c.Request().Method == "HEAD" || c.Request().Method == "OPTIONS" || c.Request().Method == "TRACE" || c.Request().Method == "CONNECT" {
			return true
		}

		return false
	},
	Handler: func(c echo.Context, reqBody, resBody []byte) {
		if c.Response().Status < 400 {
			return
		}

		c.Set(ctxKey, &dumpBody{
			reqBody: reqBody,
			resBody: resBody,
		})
	},
})

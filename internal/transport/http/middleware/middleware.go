package middleware

import (
	"github.com/labstack/echo"
)

type Middleware struct {
}

func NewMidleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization,Origin,Accept,datetime,signature,Content-Type")
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}

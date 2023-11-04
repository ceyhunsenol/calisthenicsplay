package cache

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func SetRequestCache(key string, value interface{}, c echo.Context) {
	c.Set(key, value)
}

func GetRequestCache(key string, c echo.Context) (interface{}, error) {
	value := c.Get(key)
	if value == nil {
		return nil, fmt.Errorf("key not found")
	}
	return value, nil
}

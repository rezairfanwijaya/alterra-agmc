package middleware

import "github.com/labstack/echo/v4"

func BasicAuth(username, password string, e echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, nil
}

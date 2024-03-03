package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
)

func EnableDarkModePost(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return cc.JSON(http.StatusOK, map[string]interface{}{
			"OK":    false,
			"Error": err.Error(),
		})
	}
	session.Values["darkMode"] = true
	err = session.Save(cc.Request(), cc.Response())
	if err != nil {
		return cc.JSON(http.StatusOK, map[string]interface{}{
			"OK":    false,
			"Error": err.Error(),
		})
	}
	return cc.JSON(http.StatusOK, map[string]interface{}{
		"OK": true,
	})
}

func DisableDarkModePost(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return cc.JSON(http.StatusOK, map[string]interface{}{
			"OK":    false,
			"Error": err.Error(),
		})
	}
	session.Values["darkMode"] = false
	err = session.Save(cc.Request(), cc.Response())
	if err != nil {
		return cc.JSON(http.StatusOK, map[string]interface{}{
			"OK":    false,
			"Error": err.Error(),
		})
	}
	return cc.JSON(http.StatusOK, map[string]interface{}{
		"OK": true,
	})
}

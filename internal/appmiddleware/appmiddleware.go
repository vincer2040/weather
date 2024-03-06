package appmiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(wctx.WCtx)
		session, err := cc.Store.Get(cc.Request(), "session")
		url := cc.Request().URL.String()
		if err != nil || session.IsNew || !session.Values["authenticated"].(bool) {
			if url == "/" || url == "/signin" || url == "/signup" {
				return next(c)
			}
			return c.Redirect(http.StatusSeeOther, "/signin")
		}
		if url == "/" || url == "/signin" || url == "signup" {
			return c.Redirect(http.StatusSeeOther, "/home")
		}
		return next(c)
	}
}

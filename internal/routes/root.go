package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
)

func RootGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	return cc.Render(http.StatusOK, "index.html", nil)
}

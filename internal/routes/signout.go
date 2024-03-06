package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
)

func SignoutPost(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return err
	}
	session.Values["authenticated"] = false
	session.Values["userID"] = -1
	err = session.Save(cc.Request(), cc.Response())
	if err != nil {
		return err
	}
	cc.Response().Header().Add("HX-Location", "/signin")
	return nil
}

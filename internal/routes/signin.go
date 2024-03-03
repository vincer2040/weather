package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
	"github.com/vincer2040/weather/internal/wsession"
)

func SigninGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return err
	}
	err = wsession.Setup(session, cc.Request(), cc.Response())
	if err != nil {
		return err
	}
	sessionData := wsession.GetSessionData(session)
	return cc.Render(http.StatusOK, "signin.html", map[string]interface{}{
		"SessionData": sessionData,
		"Route":       "signin",
	})
}

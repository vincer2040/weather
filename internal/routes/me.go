package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
	"github.com/vincer2040/weather/internal/wsession"
)

func MeGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return err
	}
	sessionData := wsession.GetSessionData(session)
	user, err := cc.DB.GetUserById(sessionData.UserId)
	if err != nil {
		return err
	}
	return cc.Render(http.StatusOK, "me.html", map[string]interface{}{
		"SessionData": sessionData,
		"Route":       "me",
		"User":        user,
	})
}

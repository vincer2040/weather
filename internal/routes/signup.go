package routes

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/types"
	"github.com/vincer2040/weather/internal/wctx"
	"github.com/vincer2040/weather/internal/wsession"
)

func SignUpGet(c echo.Context) error {
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
	return cc.Render(http.StatusOK, "signup.html", map[string]interface{}{
		"SessionData": sessionData,
		"Route":       "signup",
	})
}

func SignupPost(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return err
	}
	user := types.UserFromRequest(cc)
	repeatedPassword := cc.FormValue("repeat-password")
	if !user.IsValid() {
		return errors.New("invalid user")
	}
	if user.Password != repeatedPassword {
		return errors.New("passwords do not match")
	}
	err = user.HashPassword()
	if err != nil {
		return err
	}
	id, err := cc.DB.InsertUser(user)
	if err != nil {
		return err
	}
	session.Values["userID"] = id
	session.Values["authenticated"] = true
	err = session.Save(cc.Request(), cc.Response())
	if err != nil {
		return err
	}
	cc.Response().Header().Add("HX-Location", "/home")
	return nil
}

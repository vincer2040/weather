package wctx

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type WCtx struct {
	echo.Context
	Store *sessions.CookieStore
}

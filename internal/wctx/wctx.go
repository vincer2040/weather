package wctx

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/db"
	"github.com/vincer2040/weather/internal/search"
)

type WCtx struct {
	echo.Context
	Store         *sessions.CookieStore
	DB            *db.DB
	Search        *search.Search
	WeatherApiKey string
}

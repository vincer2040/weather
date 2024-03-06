package weather

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vincer2040/weather/internal/appmiddleware"
	"github.com/vincer2040/weather/internal/db"
	"github.com/vincer2040/weather/internal/env"
	"github.com/vincer2040/weather/internal/render"
	"github.com/vincer2040/weather/internal/routes"
	"github.com/vincer2040/weather/internal/wctx"
)

func Main() error {

	err := env.EnvInit()
	if err != nil {
		return err
	}

	dbUrl := env.GetDBURL()

	db, err := db.New(dbUrl)
	if err != nil {
		return err
	}

	err = db.CreateUserTable()
	if err != nil {
		return err
	}

	store := sessions.NewCookieStore([]byte(env.GetSessionSecret()))

	e := echo.New()

	e.Renderer = render.New()
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := wctx.WCtx{
				Context: c,
				Store:   store,
				DB:      db,
			}
			return next(cc)
		}
	})
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/assets", "public/assets")
	e.GET("/", routes.RootGet)
	e.GET("/signin", routes.SigninGet)
    e.GET("/signup", routes.SignUpGet)
    e.GET("/home", appmiddleware.AuthMiddleware(routes.HomeGet))
	e.POST("/enable-dark-mode", routes.EnableDarkModePost)
	e.POST("/disable-dark-mode", routes.DisableDarkModePost)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}

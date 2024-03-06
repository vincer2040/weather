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
	"github.com/vincer2040/weather/internal/search"
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

	// err = db.DropUserTable()
	// if err != nil {
	//     return err
	// }

	err = db.CreateUserTable()
	if err != nil {
		return err
	}

	s, err := search.New()
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
				Context:       c,
				Store:         store,
				DB:            db,
				Search:        s,
				WeatherApiKey: env.GetWeatherApiKey(),
			}
			return next(cc)
		}
	})
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/assets", "public/assets")
	e.GET("/", appmiddleware.AuthMiddleware(routes.RootGet))
	e.GET("/signin", appmiddleware.AuthMiddleware(routes.SigninGet))
	e.POST("/signin", appmiddleware.AuthMiddleware(routes.SigninPost))
	e.GET("/signup", appmiddleware.AuthMiddleware(routes.SignUpGet))
	e.POST("/signup", appmiddleware.AuthMiddleware(routes.SignupPost))
	e.POST("/signout", appmiddleware.AuthMiddleware(routes.SignoutPost))
	e.GET("/home", appmiddleware.AuthMiddleware(routes.HomeGet))
	e.GET("/city-search", appmiddleware.AuthMiddleware(routes.CitySearchGet))
	e.GET("/city-autocomplete", appmiddleware.AuthMiddleware(routes.CityAutocompleteGet))
	e.GET("/search/:zip", appmiddleware.AuthMiddleware(routes.SearchGet))
	e.GET("/me", appmiddleware.AuthMiddleware(routes.MeGet))
	e.POST("/enable-dark-mode", routes.EnableDarkModePost)
	e.POST("/disable-dark-mode", routes.DisableDarkModePost)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}

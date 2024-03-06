package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/wctx"
)

func CitySearchGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	return cc.String(http.StatusOK, "ok")
}

func CityAutocompleteGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	input := cc.FormValue("city")
	cities := cc.Search.GetTopTen(input)
	fmt.Println(cities)
	return cc.Render(http.StatusOK, "search-autocomplete.html", map[string]interface{}{
		"Len":    len(cities),
		"Cities": cities,
	})
}

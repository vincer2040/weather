package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/weather/internal/types"
	"github.com/vincer2040/weather/internal/wctx"
	"github.com/vincer2040/weather/internal/wsession"
)

func CitySearchGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	return cc.String(http.StatusOK, "ok")
}

func CityAutocompleteGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	input := cc.FormValue("city")
	cities := cc.Search.GetTopTen(input)
	return cc.Render(http.StatusOK, "search-autocomplete.html", map[string]interface{}{
		"Len":    len(cities),
		"Cities": cities,
	})
}

func SearchGet(c echo.Context) error {
	cc := c.(wctx.WCtx)
	session, err := cc.Store.Get(cc.Request(), "session")
	if err != nil {
		return err
	}
	sessionData := wsession.GetSessionData(session)
	zip := cc.Param("zip")

	queryParams := url.Values{
		"key":    {cc.WeatherApiKey},
		"q":      {zip},
		"days":   {strconv.Itoa(1)},
		"api":    {"no"},
		"alerts": {"no"},
	}

	url := "http://api.weatherapi.com/v1/forecast.json?" + queryParams.Encode()

	req, err := http.Get(url)
	if err != nil {
		return err
	}

	body := req.Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var weather types.Weather
	err = json.Unmarshal(data, &weather)
	if err != nil {
		return err
	}

	location := weather.Location
	current := weather.Current
	todaysForecast := weather.Forecast.ForecastDay[0]

	return cc.Render(http.StatusOK, "search.html", map[string]interface{}{
		"SessionData":    sessionData,
		"Route":          "search",
		"Location":       location,
		"Current":        current,
		"TodaysForecast": todaysForecast,
	})
}

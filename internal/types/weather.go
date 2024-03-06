package types

type Weather struct {
	Location WeatherLocation `json:"location"`
	Current  WeatherCurrent  `json:"current"`
	Forecast WeatherForecast `json:"forecast"`
}

type WeatherLocation struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"tz_id"`
	LocalTimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type WeatherCurrent struct {
	LastUpdatedEpoch int              `json:"last_updated_epoch"`
	LastUpdated      string           `json:"last_updated"`
	TempC            float64          `json:"temp_c"`
	TempF            float64          `json:"temp_f"`
	IsDay            int              `json:"is_day"`
	Condition        WeatherCondition `json:"condition"`
	WindMPH          float64          `json:"wind_mph"`
	WindKPH          float64          `json:"wind_kph"`
	WindDegree       float64          `json:"wind_degree"`
	WindDir          string           `json:"wind_dir"`
	PressureMB       float64          `json:"pressure_mb"`
	PressureIN       float64          `json:"pressure_in"`
	PrecipMM         float64          `json:"precip_mm"`
	PrecipIN         float64          `json:"precip_in"`
	Humidity         float64          `json:"humidity"`
	Cloud            float64          `json:"cloud"`
	FeelsLikeC       float64          `json:"feelslike_c"`
	FeelsLikeF       float64          `json:"feelslike_f"`
	VisKM            float64          `json:"vis_km"`
	VisMiles         float64          `json:"vis_miles"`
	UV               float64          `json:"uv"`
	GustMPH          float64          `json:"gust_mph"`
	GustKPH          float64          `json:"gust_kph"`
}

type WeatherCondition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type WeatherForecast struct {
	ForecastDay []WeatherDayForecast `json:"forecastday"`
}

type WeatherDayForecast struct {
	Date      string        `json:"date"`
	DateEpoch int           `json:"date_epoch"`
	Day       WeatherDay    `json:"day"`
	Astro     WeatherAstro  `json:"astro"`
	Hour      []WeatherHour `json:"hour"`
}

type WeatherDay struct {
	MaxTempC          float64          `json:"maxtemp_c"`
	MaxTempF          float64          `json:"maxtemp_f"`
	MinTempC          float64          `json:"mintemp_c"`
	MinTempF          float64          `json:"mintemp_f"`
	AverageTempC      float64          `json:"avgtemp_c"`
	AverageTempF      float64          `json:"avgtemp_f"`
	MaxWindMPH        float64          `json:"maxwind_mph"`
	MaxWindKPH        float64          `json:"maxwind_kph"`
	TotalPrecipMM     float64          `json:"totalprecip_mm"`
	TotalPrecipIN     float64          `json:"totalprecip_in"`
	TotalSnowCM       float64          `json:"totalsnow_cm"`
	AverageVisKM      float64          `json:"avgvis_km"`
	AverageVisMiles   float64          `json:"avgvis_miles"`
	AverageHumidity   float64          `json:"avghumidity"`
	DailyWillItRain   int              `json:"daily_will_it_rain"`
	DailyChanceOfRain float64          `json:"daily_chance_of_rain"`
	DailyWillItSnow   float64          `json:"daily_will_it_snow"`
	DailyChanceOfSnow float64          `json:"daily_chance_of_snow"`
	Condition         WeatherCondition `json:"condition"`
	UV                float64          `json:"uv"`
}

type WeatherAstro struct {
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	Moonrise  string `json:"moonrise"`
	Moonset   string `json:"moonset"`
	MoonPhase string `json:"moon_phase"`
	IsMoonUp  int    `json:"is_moon_up"`
	IsSunUp   int    `json:"is_sun_up"`
}

type WeatherHour struct {
	TimeEpoch    int              `json:"time_epoch"`
	Time         string           `json:"time"`
	TempC        float64          `json:"temp_c"`
	TempF        float64          `json:"temp_f"`
	IsDay        int              `json:"is_day"`
	Condition    WeatherCondition `json:"condition"`
	WindMPH      float64          `json:"wind_mph"`
	WindKPH      float64          `json:"wind_kph"`
	WindDegree   float64          `json:"wind_degree"`
	WindDir      string           `json:"wind_dir"`
	PressureMB   float64          `json:"pressure_mb"`
	PressureIN   float64          `json:"pressure_in"`
	PrecipMM     float64          `json:"precip_mm"`
	PrecipIN     float64          `json:"precip_in"`
	SnowCM       float64          `json:"snow_cm"`
	Humidity     float64          `json:"humidity"`
	Cloud        float64          `json:"cloud"`
	FeelsLikeC   float64          `json:"feelslike_c"`
	FeelsLikeF   float64          `json:"feelslike_f"`
	WindChillC   float64          `json:"windchill_c"`
	WindChillF   float64          `json:"windchill_f"`
	HeadIndexC   float64          `json:"heatindex_c"`
	HeadIndexF   float64          `json:"heatindex_f"`
	DewPointC    float64          `json:"dewpoint_c"`
	DewPointF    float64          `json:"dewpoint_f"`
	WillItRain   int              `json:"will_it_rain"`
	ChanceOfRain float64          `json:"change_of_rain"`
	WillItSnow   int              `json:"will_it_snow"`
	ChanceOfSnow float64          `json:"chance_of_snow"`
	VisKM        float64          `json:"vis_km"`
	VisMiles     float64          `json:"vis_miles"`
	GustMPH      float64          `json:"gust_mph"`
	GustKPH      float64          `json:"gust_kph"`
	UV           float64          `json:"uv"`
}

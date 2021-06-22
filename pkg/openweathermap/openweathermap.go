package openweathermap

type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   float64   `json:"pressure"`
	Humidity   float64   `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     float64   `json:"clouds"`
	Visibility float64   `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    float64   `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Minutely struct {
	Dt            int     `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   float64   `json:"pressure"`
	Humidity   float64   `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     float64   `json:"clouds"`
	Visibility float64   `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    float64   `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain,omitempty"`
}

type Rain struct {
	H float64 `json:"1h"`
}

type Daily struct {
	Dt        int       `json:"dt"`
	Sunrise   int       `json:"sunrise"`
	Sunset    int       `json:"sunset"`
	Moonrise  int       `json:"moonrise"`
	Moonset   int       `json:"moonset"`
	MoonPhase float64   `json:"moon_phase"`
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  float64   `json:"pressure"`
	Humidity  float64   `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   float64   `json:"wind_deg"`
	WindGust  float64   `json:"wind_gust"`
	Weather   []Weather `json:"weather"`
	Clouds    float64   `json:"clouds"`
	Pop       float64   `json:"pop"`
	Rain      float64   `json:"rain,omitempty"`
	Uvi       float64   `json:"uvi"`
}

type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Alert struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int      `json:"start"`
	End         int      `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type Response struct {
	Lat            float64    `json:"lat"`
	Lon            float64    `json:"lon"`
	Timezone       string     `json:"timezone"`
	TimezoneOffset int        `json:"timezone_offset"`
	Current        Current    `json:"current"`
	Minutely       []Minutely `json:"minutely,omitempty"`
	Hourly         []Hourly   `json:"hourly,omitempty"`
	Daily          []Daily    `json:"daily,omitempty"`
	Alerts         []Alert    `json:"alerts"`
}

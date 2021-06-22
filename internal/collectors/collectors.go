package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var sunriseCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_sunrise",
	Help: "openweather_sunrise",
}, []string{"type"})

var sunsetCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_sunset",
	Help: "openweather_sunset",
}, []string{"type"})

var tempCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_temp",
	Help: "openweather_temp",
}, []string{"type", "daytime"})

var feelsLikeCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_feels_like",
	Help: "openweather_feels_like",
}, []string{"type", "daytime"})

var pressureCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_pressure",
	Help: "openweather_pressure",
}, []string{"type"})

var humidityCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_humidity",
	Help: "openweather_humidity",
}, []string{"type"})

var dewPointCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_dew_point",
	Help: "openweather_dew_point",
}, []string{"type"})

var windSpeedCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_wind_speed",
	Help: "openweather_wind_speed",
}, []string{"type"})

var windDegCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_wind_deg",
	Help: "openweather_wind_deg",
}, []string{"type"})

var windGustCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_wind_gust",
	Help: "openweather_wind_gust",
}, []string{"type"})

var cloudsCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_clouds",
	Help: "openweather_clouds",
}, []string{"type"})

var moonsetCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_moonset",
	Help: "openweather_moonset",
}, []string{"type"})

var moonphaseCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_moonphase",
	Help: "openweather_moonphase",
}, []string{"type"})

var rainCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_rain",
	Help: "openweather_rain",
}, []string{"type"})

var uviCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_uvi",
	Help: "openweather_uvi",
}, []string{"type"})

var popCollector = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "openweather_pop",
	Help: "openweather_pop",
}, []string{"type"})

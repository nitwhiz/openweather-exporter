package collectors

import (
	"github.com/nitwhiz/openweathermap-exporter/pkg/openweathermap"
)

func CollectCurrent(r *openweathermap.Response) {
	c := r.Current

	sunriseCollector.WithLabelValues("current").Set(float64(c.Sunrise))
	sunsetCollector.WithLabelValues("current").Set(float64(c.Sunset))

	tempCollector.WithLabelValues("current", "current").Set(c.Temp)
	feelsLikeCollector.WithLabelValues("current", "current").Set(c.FeelsLike)

	pressureCollector.WithLabelValues("current").Set(c.Pressure)
	humidityCollector.WithLabelValues("current").Set(c.Humidity)
	dewPointCollector.WithLabelValues("current").Set(c.DewPoint)

	windSpeedCollector.WithLabelValues("current").Set(c.WindSpeed)
	windDegCollector.WithLabelValues("current").Set(c.WindDeg)
	windGustCollector.WithLabelValues("current").Set(c.WindGust)

	cloudsCollector.WithLabelValues("current").Set(c.Clouds)
	uviCollector.WithLabelValues("current").Set(c.Uvi)
}

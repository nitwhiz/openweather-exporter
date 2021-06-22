package collectors

import (
	"fmt"
	"github.com/nitwhiz/openweathermap-exporter/pkg/openweathermap"
)

func CollectDaily(r *openweathermap.Response) {
	for i, d := range r.Daily {
		day := fmt.Sprintf("daily_%d", i)

		sunriseCollector.WithLabelValues(day).Set(float64(d.Sunrise))
		sunsetCollector.WithLabelValues(day).Set(float64(d.Sunset))

		tempCollector.WithLabelValues(day, "day").Set(d.Temp.Day)
		tempCollector.WithLabelValues(day, "min").Set(d.Temp.Min)
		tempCollector.WithLabelValues(day, "max").Set(d.Temp.Max)
		tempCollector.WithLabelValues(day, "night").Set(d.Temp.Night)
		tempCollector.WithLabelValues(day, "eve").Set(d.Temp.Eve)
		tempCollector.WithLabelValues(day, "morn").Set(d.Temp.Morn)

		feelsLikeCollector.WithLabelValues(day, "day").Set(d.Temp.Day)
		feelsLikeCollector.WithLabelValues(day, "night").Set(d.Temp.Night)
		feelsLikeCollector.WithLabelValues(day, "eve").Set(d.Temp.Eve)
		feelsLikeCollector.WithLabelValues(day, "morn").Set(d.Temp.Morn)

		moonsetCollector.WithLabelValues(day).Set(float64(d.Moonset))
		moonphaseCollector.WithLabelValues(day).Set(d.MoonPhase)

		pressureCollector.WithLabelValues(day).Set(d.Pressure)
		humidityCollector.WithLabelValues(day).Set(d.Humidity)
		dewPointCollector.WithLabelValues(day).Set(d.DewPoint)

		windSpeedCollector.WithLabelValues(day).Set(d.WindSpeed)
		windDegCollector.WithLabelValues(day).Set(d.WindDeg)
		windGustCollector.WithLabelValues(day).Set(d.WindGust)

		cloudsCollector.WithLabelValues(day).Set(d.Clouds)
		rainCollector.WithLabelValues(day).Set(d.Rain)
		uviCollector.WithLabelValues(day).Set(d.Uvi)

		popCollector.WithLabelValues(day).Set(d.Pop)
	}
}

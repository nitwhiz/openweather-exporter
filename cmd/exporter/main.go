package main

import (
	"encoding/json"
	"fmt"
	"github.com/nitwhiz/openweathermap-exporter/internal/collectors"
	"github.com/nitwhiz/openweathermap-exporter/pkg/openweathermap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func getWeatherResponse() (*openweathermap.Response, error) {
	lat := os.Getenv("OPENWEATHER_LAT")
	lon := os.Getenv("OPENWEATHER_LON")
	units := os.Getenv("OPENWEATHER_UNITS")
	lang := os.Getenv("OPENWEATHER_LANG")
	appid := os.Getenv("OPENWEATHER_APPID")

	if units == "" {
		units = "metric"
	}

	if lang == "" {
		lang = "en"
	}

	if lat == "" {
		panic("please define OPENWEATHER_LAT")
	}

	if lon == "" {
		panic("please define OPENWEATHER_LON")
	}

	if appid == "" {
		panic("please define OPENWEATHER_APPID")
	}

	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%s&lon=%s&units=%s&lang=%s&exclude=minutely&appid=%s", lat, lon, units, lang, appid))

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {

		}
	}(res.Body)

	var r openweathermap.Response

	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func collector() {
	go func() {
		for {
			r, err := getWeatherResponse()

			if err != nil {
				fmt.Println(err)
			} else {
				var wg sync.WaitGroup

				wg.Add(2)

				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					collectors.CollectCurrent(r)
				}(&wg)

				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					collectors.CollectDaily(r)
				}(&wg)

				wg.Wait()
			}

			time.Sleep(time.Minute * 5)
		}
	}()
}

func main() {
	collector()

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":9924", nil)

	if err != nil {
		panic(err)
	}
}

# openweather-exporter

An exporter for [OpenWeather OneCall](https://openweathermap.org/api/one-call-api) data.

## Example Systemd Unit File

```
[Unit]
Description=OpenWeather Exporter
After=network-online.target

[Service]
Type=simple
Environment=OPENWEATHER_LAT=66.8979079
Environment=OPENWEATHER_LON=-52.8688291
Environment=OPENWEATHER_APPID=<your api key>
ExecStart=/opt/openweather_exporter/openweather_exporter

[Install]
WantedBy=multi-user.target
```

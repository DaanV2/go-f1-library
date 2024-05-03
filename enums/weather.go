package enums

type WeatherType uint8

const (
	Clear      WeatherType = 0
	LightCloud WeatherType = 1
	Overcast   WeatherType = 2
	LightRain  WeatherType = 3
	HeavyRain  WeatherType = 4
	Storm      WeatherType = 5
)
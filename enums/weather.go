package enums

// WeatherType represents the weather type.
type WeatherType uint8

const (
	WT_Clear      WeatherType = 0 // weather: clear
	WT_LightCloud WeatherType = 1 // weather: lightcloud
	WT_Overcast   WeatherType = 2 // weather: overcast
	WT_LightRain  WeatherType = 3 // weather: light rain
	WT_HeavyRain  WeatherType = 4 // weather: heavy rain
	WT_Storm      WeatherType = 5 // weather: storm
)

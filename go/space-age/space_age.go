package space

type Planet string

var conv = map[string]float64{
	"Mercury":      0.2408467,
	"Venus":        0.61519726,
	"Mars":         1.8808158,
	"Jupiter":      11.862615,
	"Saturn":       29.447498,
	"Uranus":       84.016846,
	"Neptune":      164.79132,
	"EarthDays":    365.25,
	"EarthSeconds": 31557600,
	"Earth":        1.0,
}

func Age(seconds float64, planet Planet) float64 {
	if _, ok := conv[string(planet)]; !ok {
		return -1.0
	}
	return (seconds / conv["EarthSeconds"]) / conv[string(planet)]
}

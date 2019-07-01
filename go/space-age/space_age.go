//Package space converts age from seconds to years (based on the planet's orbital cycle)
package space

//A Planet is a string
type Planet string

//EarthYears returns the number of earth years given the seconds as an argument
func EarthYears(seconds float64) float64 {
	return seconds / 31557600
}

//Age returns the number of planet years
func Age(seconds float64, planet Planet) float64 {
	planetsToEarthYears := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	earthYears := EarthYears(seconds)
	planetYears := earthYears / planetsToEarthYears[planet]
	// fmt.Printf(
	// 	"%v seconds == %v Earth Years == %v %v years\n",
	// 	seconds,
	// 	earthYears,
	// 	planetYears,
	// 	planet,
	// )
	return planetYears
}

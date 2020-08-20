package space

type Planet string

const secondsInday = 31556952

var planetsMap = map[Planet]float64{"Earth": 1, "Mercury": 0.2408467, "Venus": 0.61519726, "Mars": 1.8808158, "Jupiter": 11.862615, "Saturn": 29.447498, "Uranus": 84.016846, "Neptune": 164.79132}

func Age(sec float64, p Planet) float64 {
	secondsInEarth := sec / secondsInday

	return secondsInEarth / planetsMap[p]
}

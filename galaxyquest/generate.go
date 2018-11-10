package uni

import (
	"math"
	"math/rand"
)

func GenerateUniverseTest() (Universe, int) {
	var clusterMax int
	u := Universe{
		Distance: rand.Int63n(1000000) + 2,
		Points:   make([]Point, 0),
	}
	remaining := rand.Int() % 1000000
	allCenters := make([]Point, 0)

	for remaining > 0 {
		var center Point
		safeDistance := false
		for !safeDistance {
			center.X = int64(rand.Int() % 1000000000)
			center.Y = int64(rand.Int() % 1000000000)
			safeDistance = farEnoughAway(allCenters, center, u.Distance)
		}
		allCenters = append(allCenters, center)

		n := rand.Int() % (700000)
		if n > remaining {
			n = remaining
		}
		remaining -= n
		if n > clusterMax {
			clusterMax = n
		}

		newPoints := make([]Point, n)
		for i := 0; i < n; i++ {
			r := int64(math.Sqrt(float64(2*(u.Distance*u.Distance)))) / 2
			a := rand.Int63n(r + 1)
			b := int64(math.Sqrt(float64(r*r - a*a)))
			p := Point{
				X: center.X + a,
				Y: center.Y + b,
			}
			newPoints[i] = p
		}
		u.Points = append(u.Points, newPoints...)
	}

	if clusterMax < len(u.Points)/2 {
		clusterMax = 0
	}

	return u, clusterMax
}

func farEnoughAway(points []Point, point Point, distance int64) bool {
	for _, p := range points {
		if SameGalaxy(p, point, distance*2) {
			return false
		}
	}
	return true
}

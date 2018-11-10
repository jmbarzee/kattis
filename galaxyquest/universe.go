package uni

type (
	Universe struct {
		Distance int64
		Points   []Point
	}

	Point struct {
		X int64
		Y int64
	}
)

func FindMaxGalaxyCount(u Universe) int {
	// TODO @jmbarzee speed up by stoping after visiting more than half

	half := len(u.Points) / 2
	PointsByX := splitByX(u.Points, u.Distance, half)
	for _, plx := range PointsByX {

		pointsByY := splitByY(plx, u.Distance, half)
		for _, ply := range pointsByY {

			n := countPosibleGalaxy(ply, u.Distance, half)
			if n != 0 {
				return n
			} else {
				return 0
			}
		}
	}
	return 0
}

func countPosibleGalaxy(points []Point, distance int64, half int) int {
	stars := NewStarList()
	for _, p := range points {
		stars.Append(&Star{
			Point: p,
		})
	}

	for i := 0; i < stars.Len; i++ {
		star := stars.Head
		stars.Remove(star)
		nextStar := stars.Head
		starsFound := 1
		for j := 0; j < stars.Len; j++ {
			crntStar := nextStar
			nextStar = nextStar.Next
			if SameGalaxy(star.Point, crntStar.Point, distance) {
				stars.Remove(crntStar)
				starsFound++
				j--
			}
		}
		if starsFound > half {
			return starsFound
		}
		if stars.Len <= half {
			return 0
		}
	}
	return 0
}

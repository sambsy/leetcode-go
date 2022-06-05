package solute

import (
	"math/rand"
)

type Solution struct {
	x, y   float64
	radius float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x:      x_center,
		y:      y_center,
		radius: radius,
	}
}

func (s *Solution) RandPoint() []float64 {
	for {
		rand.Float64()
		x := rand.Float64() * s.radius
		y := rand.Float64() * s.radius
		if x*x+y*y > s.radius*s.radius {
			continue
		}
		if rand.Intn(2)%2 == 0 {
			x = x + s.x
		} else {
			x = s.x - x
		}

		if rand.Intn(2)%2 == 0 {
			y = y + s.y
		} else {
			y = s.y - y
		}
		return []float64{x, y}
	}
}

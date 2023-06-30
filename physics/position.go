package physics

import (
	"math"
	"time"
)

const g = 9.8

type Point struct {
	X, Y float64
}

func GetArrowPosition(t_ time.Duration, v, a float64, start Point) Point {
	t := float64(t_) / float64(time.Second)

	x := start.X + v*math.Cos(a*math.Pi/180)*t
	y := start.Y + v*math.Sin(a*math.Pi/180)*t - g*math.Pow(t, 2)/2

	return Point{x, y}
}

// TODO:
func GetArrowAngle(t_ time.Duration, v, a float64) float64 {
	t := float64(t_) / float64(time.Second)

	vx := v * math.Cos(a*math.Pi/180) * t
	vy := v*math.Sin(a*math.Pi/180)*t - g*math.Pow(t, 2)/2

	return math.Atan(vy / vx)
}

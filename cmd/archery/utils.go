package main

import "math"

func degToRad(deg float64) float64 { return math.Mod(deg, 360) * math.Pi / 180.0 }

func radToDegree(rad float64) float64 { return rad * 180.0 / math.Pi }

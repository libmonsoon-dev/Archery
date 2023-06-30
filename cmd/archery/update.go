package main

import (
	"archery/physics"
	"time"
)

func (g *Game) Update() error {
	g.updateArrow()

	return nil
}

func (g *Game) updateArrow() {
	//TODO:
	sinceStart := time.Duration(0) // time.Since(g.start)
	g.arrow.Position = physics.GetArrowPosition(sinceStart, g.arrow.Speed, g.arrow.InitialAngle, g.bow.Position)
	// g.arrow.Angle = physics.GetArrowAngle(sinceStart, g.arrow.Speed, g.arrow.InitialAngle)
}

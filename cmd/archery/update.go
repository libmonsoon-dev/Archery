package main

import (
	"archery/physics"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.updateAimLine()
	g.updateArrows()

	return nil
}

func (g *Game) updateAimLine() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.aimLine.IsActive = true

		x, y := ebiten.CursorPosition()
		g.aimLine.startPositionX, g.aimLine.startPositionY = x, g.screen.Bounds().Dy()-y
	}

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		x, y := ebiten.CursorPosition()
		g.aimLine.endPositionX, g.aimLine.endPositionY = x, g.screen.Bounds().Dy()-y
		g.updateBow()
		//TODO: g.bow.Angle = based on aim line
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.aimLine.IsActive = false
		g.shoot()
	}
}

func (g *Game) updateBow() {
	x, y := g.aimLine.endPositionX-g.aimLine.startPositionX, g.aimLine.endPositionY-g.aimLine.startPositionY
	length := math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
	g.bow.InitialSpeed = length
	g.bow.Angle = radToDegree(math.Acos(-float64(x) / length))
	if g.aimLine.endPositionY > g.aimLine.startPositionY {
		g.bow.Angle = 360 - g.bow.Angle
	}
	if !g.arrow.IsShooted() {
		g.arrow.Angle = g.bow.Angle
	}
}

func (g *Game) shoot() {
	g.arrow.StartTime = time.Now()
	g.arrow.Speed = g.bow.InitialSpeed
	g.arrow.Angle = g.bow.Angle
	g.arrow.InitialAngle = g.bow.Angle
}

func (g *Game) updateArrows() {
	if g.arrow.IsShooted() {
		sinceStart := time.Since(g.arrow.StartTime)
		g.arrow.Position = physics.GetArrowPosition(sinceStart, g.arrow.Speed, g.arrow.InitialAngle, g.bow.Position)
		g.arrow.Angle = radToDegree(physics.GetArrowAngle(sinceStart, g.arrow.Speed, g.arrow.InitialAngle))
		if g.aimLine.endPositionX > g.aimLine.startPositionX {
			g.arrow.Angle += 180
		}
	}

	if g.arrow.Position.Y <= 0 {
		g.arrow.StartTime = time.Time{}
		g.arrow.Position = g.bow.Position
		g.arrow.InitialAngle = g.bow.Angle
	}
}

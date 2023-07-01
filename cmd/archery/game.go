package main

import (
	"archery/physics"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type bow struct {
	Position     physics.Point
	Angle        float64
	InitialSpeed float64
}

type arrow struct {
	Position     physics.Point
	Speed        float64
	Angle        float64
	InitialAngle float64
	StartTime    time.Time
}

func (a arrow) IsShooted() bool {
	return !a.StartTime.IsZero()
}

type aimLine struct {
	IsActive                                                   bool
	startPositionX, startPositionY, endPositionX, endPositionY int
}

type Game struct {
	screen *ebiten.Image

	bow     bow
	arrow   arrow
	aimLine aimLine
}

func NewGame() *Game {
	pos := physics.Point{X: 10, Y: 10}
	return &Game{
		bow:   bow{Position: pos},
		arrow: arrow{Position: pos},
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

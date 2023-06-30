package main

import (
	"archery/physics"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type bow struct {
	Position physics.Point
	Angle    float64
}

type arrow struct {
	Position     physics.Point
	Speed        float64
	Angle        float64
	InitialAngle float64
}

type aimLine struct {
	IsActive                   bool
	startPosition, endPosition physics.Point
}

type Game struct {
	screen *ebiten.Image

	start time.Time
	bow
	arrow
	aimLine
}

func NewGame() *Game {
	pos := physics.Point{X: 10, Y: 10}
	return &Game{
		start: time.Now(),

		bow:   bow{Position: pos},
		arrow: arrow{Position: pos, Speed: 40, InitialAngle: 60},
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

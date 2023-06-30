package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.screen = screen

	g.drawBow()
	g.drawArrow()

	ebitenutil.DebugPrint(g.screen, fmt.Sprint("Angle", g.arrow.Angle))
}

func (g *Game) drawBow() {
	var geoM ebiten.GeoM
	geoM.Translate(-float64(bowImg.Bounds().Dx())/2, float64(bowImg.Bounds().Dy())/2)
	geoM.Rotate(math.Pi / 2)

	offsetX := g.bow.Position.X*meterToPixelRate + float64(bowImg.Bounds().Dx())
	offsetY := float64(g.screen.Bounds().Dy()) - g.bow.Position.Y*meterToPixelRate
	geoM.Translate(offsetX, offsetY)

	g.screen.DrawImage(bowImg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (g *Game) drawArrow() {
	var geoM ebiten.GeoM
	//TODO:
	// geoM.Rotate(float64(g.arrow.InitialAngle)*(math.Pi/180.0) - g.arrow.Angle)

	offsetX := g.arrow.Position.X*meterToPixelRate - float64(arrowImg.Bounds().Dx())/2
	offsetY := float64(g.screen.Bounds().Dy()) - g.arrow.Position.Y*meterToPixelRate - float64(arrowImg.Bounds().Dy())/2
	geoM.Translate(offsetX, offsetY)

	g.screen.DrawImage(arrowImg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

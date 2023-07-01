package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.screen = screen

	g.drawAimLine()
	g.drawBow()
	g.drawArrow()

	ebitenutil.DebugPrint(screen, fmt.Sprint(g.arrow.InitialAngle))
}

func (g *Game) drawBow() {
	var geoM ebiten.GeoM
	geoM.Translate(float64(-bowImg.Bounds().Dx())/2, float64(-bowImg.Bounds().Dy())/2)
	geoM.Rotate(2.3 * math.Pi / 3)

	offsetX := g.bow.Position.X * meterToPixelRate //+ float64(bowImg.Bounds().Dx())
	offsetY := float64(g.screen.Bounds().Dy()) - g.bow.Position.Y*meterToPixelRate
	geoM.Translate(offsetX, offsetY)

	g.screen.DrawImage(bowImg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (g *Game) drawArrow() {
	var geoM ebiten.GeoM
	// geoM.Translate(float64(-arrowImg.Bounds().Dx())/2, float64(-arrowImg.Bounds().Dy())/2)
	// TODO:
	// geoM.Rotate(degToRad(45))
	// geoM.Rotate(float64(g.arrow.InitialAngle)*(math.Pi/180.0) - g.arrow.Angle)

	offsetX := g.arrow.Position.X * meterToPixelRate //- float64(arrowImg.Bounds().Dx())/2
	offsetY := float64(g.screen.Bounds().Dy()) - g.arrow.Position.Y*meterToPixelRate - float64(arrowImg.Bounds().Dy())/2
	geoM.Translate(offsetX, offsetY)

	g.screen.DrawImage(arrowImg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (g *Game) drawAimLine() {
	if !g.aimLine.IsActive {
		return
	}

	screenDy := g.screen.Bounds().Dy()

	vector.StrokeLine(
		g.screen,
		float32(g.aimLine.startPositionX), float32(screenDy-g.aimLine.startPositionY), float32(g.aimLine.endPositionX), float32(screenDy-g.aimLine.endPositionY),
		2,
		color.White,
		true,
	)

	// ebitenutil.DebugPrint(g.screen, fmt.Sprintf("%+v", g.aimLine))
}

package main

import (
	"fmt"
	"image/color"

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

	ebitenutil.DebugPrint(screen, fmt.Sprint(g.bow.Angle, g.arrow.Angle))
}

func (g *Game) drawBow() {
	var geoM ebiten.GeoM
	geoM.Translate(float64(-bowImg.Bounds().Dx())/2, float64(-bowImg.Bounds().Dy())/2)
	geoM.Rotate(degToRad(360 - g.bow.Angle + 135))

	offsetX := g.bow.Position.X * meterToPixelRate
	offsetY := float64(g.screen.Bounds().Dy()) - g.bow.Position.Y*meterToPixelRate
	geoM.Translate(offsetX, offsetY)

	g.screen.DrawImage(bowImg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (g *Game) drawArrow() {
	var geoM ebiten.GeoM
	geoM.Translate(float64(-arrowImg.Bounds().Dx())/2, float64(-arrowImg.Bounds().Dy())/2)
	geoM.Rotate(degToRad(360 - g.arrow.Angle + 45))

	offsetX := g.arrow.Position.X * meterToPixelRate
	offsetY := float64(g.screen.Bounds().Dy()) - g.arrow.Position.Y*meterToPixelRate
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

}

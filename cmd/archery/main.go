package main

import (
	"archery/assets"
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 640
	screenHeight = 480

	meterToPixelRate = 5
)

var (
	arrowImg = decodeImg(assets.Arrow)
	bowImg   = decodeImg(assets.Bow)
)

func decodeImg(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Archery")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

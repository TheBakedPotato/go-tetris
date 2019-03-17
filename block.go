package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Block struct {
	img    *ebiten.Image
	geoM   *ebiten.GeoM
	center *Point
}

func NewBlock(img *ebiten.Image, geoM *ebiten.GeoM) *Block {
	point := &Point{X: float64(img.Bounds().Dx() / 2), Y: float64(img.Bounds().Dy() / 2)}
	return &Block{img: img, geoM: geoM, center: point}
}

func (b *Block) GetImage() *ebiten.Image {
	return b.img
}

func (b *Block) GetImageOptions() *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{GeoM: *b.geoM}
}

func (b *Block) GetSize() (width, height int) {
	return b.img.Size()
}

func (b *Block) GetHeight() (height int) {
	_, height = b.img.Size()
	return
}

func (b *Block) GetWidth() (width int) {
	width, _ = b.img.Size()
	return
}

func (b *Block) GetPosition() *Point {
	return nil
}

func (b *Block) GetCenter() *Point {
	return b.center
}

func (b *Block) Translate(x, y float64) {
	b.geoM.Translate(x, y)
	b.updateCenter(x, y)
}

func (b *Block) Draw(targetImage *ebiten.Image) {
	targetImage.DrawImage(b.img, b.GetImageOptions())
}

func (b *Block) Rotate(angle float64, focalPoint Point) {
	dX := focalPoint.X - b.center.X
	dY := focalPoint.Y - b.center.Y

	rotatedX := math.Cos(-angle)*dX - math.Sin(-angle)*dY + focalPoint.X
	rotatedY := math.Sin(-angle)*dX + math.Cos(-angle)*dY + focalPoint.Y

	dX = rotatedX - b.center.X
	dY = rotatedY - b.center.Y

	b.Translate(dX, dY)
}

func (b *Block) updateCenter(x, y float64) {
	b.center.X += x
	b.center.Y += y
}

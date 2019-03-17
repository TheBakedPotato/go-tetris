package main

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

type Line struct {
	img  *ebiten.Image
	geoM *ebiten.GeoM
}

func NewLine(width, height int, x, y float64) *Line {
	img, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	img.Fill(colornames.Black)
	geoM := &ebiten.GeoM{}
	geoM.Translate(x, y)
	return &Line{img: img, geoM: geoM}
}

func (l *Line) Draw(targetImage *ebiten.Image) {
	targetImage.DrawImage(l.img, &ebiten.DrawImageOptions{GeoM: *l.geoM})
}

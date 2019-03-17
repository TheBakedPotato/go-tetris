package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

type Grid struct {
	width     uint16
	height    uint16
	blockSize uint16

	// img        *ebiten.Image
	tetrominos []Tetromino
}

// func createGrid(g *Grid) *ebiten.Image {
// 	img, _ := ebiten.NewImage((int)(g.width*g.blockSize), (int)(g.height*g.blockSize), ebiten.FilterDefault)

// 	return nil
// }

func NewGrid(width, height, blockSize uint16) *Grid {
	return &Grid{width: width, height: height, blockSize: blockSize}
}

func (g *Grid) DrawGrid(image *ebiten.Image) {
	for _, tetromino := range g.tetrominos {
		tetromino.Draw(image)
	}
}

func generateBlocks(blockWidth int) (blocks []*Block) {
	blocks = make([]*Block, 4)
	colors := [5]color.Color{colornames.Red, colornames.Orange, colornames.Yellow, colornames.Green, colornames.Blue}

	for i := 0; i < len(blocks); i++ {
		var newImage *ebiten.Image
		newImage, _ = ebiten.NewImage(blockWidth, blockWidth, ebiten.FilterDefault)
		newImage.Fill(colors[i])
		blocks[i] = NewBlock(newImage, &ebiten.GeoM{})
	}

	return
}

func (g *Grid) AddTetromino() {
	blocks := generateBlocks((int)(g.blockSize))
	g.tetrominos = append(g.tetrominos, NewTetromino(TetrominoI, blocks))
}

func (g *Grid) MoveTetrominoRight() {
	tetromino := g.TetrominoInPlay()
	tetromino.Translate((float64)(g.blockSize), 0.0)
}

func (g *Grid) MoveTetrominoLeft() {
	tetromino := g.TetrominoInPlay()
	tetromino.Translate(-(float64)(g.blockSize), 0.0)
}

// Temp function to cause tetro to move down
func (g *Grid) MoveTetromino() {
	tetromino := g.TetrominoInPlay()
	tetromino.Translate(0.0, (float64)(g.blockSize))
}

func (g *Grid) RotateTetrominoRight() {
	tetromino := g.TetrominoInPlay()
	tetromino.Rotate(math.Pi / 2)
}

func (g *Grid) RotateTetrominoLeft() {
	tetromino := g.TetrominoInPlay()
	tetromino.Rotate(-math.Pi / 2)
}

func (g *Grid) TetrominoInPlay() Tetromino {
	return g.tetrominos[0]
}

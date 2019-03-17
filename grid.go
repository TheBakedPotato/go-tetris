package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

type Grid struct {
	position  Point
	columns   int
	rows      int
	blockSize int

	tetrominos []Tetromino

	gridLines []*Line
}

func NewGrid(xPos, yPos float64, rows, columns, blockSize int) *Grid {
	grid := &Grid{position: Point{X: xPos, Y: yPos}, rows: rows, columns: columns, blockSize: blockSize}
	grid.generateGridLines()
	return grid
}

func (g *Grid) DrawGrid(image *ebiten.Image) {
	for _, gridLine := range g.gridLines {
		gridLine.Draw(image)
	}
	for _, tetromino := range g.tetrominos {
		tetromino.Draw(image)
	}
}

func (g *Grid) generateGridLines() {
	lineWidth := 1
	width := g.columns * g.blockSize
	height := g.rows * g.blockSize

	for x := g.position.X; x <= g.position.X+float64(width); x += float64(g.blockSize) {
		g.gridLines = append(g.gridLines, NewLine(lineWidth, height, x, g.position.Y))
	}

	for y := g.position.Y; y <= g.position.Y+float64(height); y += float64(g.blockSize) {
		g.gridLines = append(g.gridLines, NewLine(width, lineWidth, g.position.X, float64(y)))
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
	blocks := generateBlocks(g.blockSize)
	tetromino := NewTetromino(TetrominoI, blocks)
	// tetromino.SetPosition(g.position.X, g.position.Y)
	g.tetrominos = append(g.tetrominos, tetromino)
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

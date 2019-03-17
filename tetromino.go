package main

import (
	"github.com/hajimehoshi/ebiten"
)

var tetrominoRegistry = make(map[TetrominoType]func([]*Block) Tetromino)

type TetrominoCreateFunc func([]*Block) Tetromino

func RegisterTetromino(name string, createFunc TetrominoCreateFunc) TetrominoType {
	tetrominoRegistry[TetrominoType(name)] = createFunc
	return TetrominoType(name)
}

type Point struct {
	X float64
	Y float64
}

type tetromino struct {
	length     uint8
	width      uint8
	blocks     []*Block
	focalPoint Point
}

type Tetromino interface {
	Draw(targetImage *ebiten.Image)
	Translate(x, y float64)
	MoveDown(y float64)
	Rotate(angle float64)
	SetPosition(x, y float64)
}

func NewTetromino(tetrominoType TetrominoType, blocks []*Block) Tetromino {
	return tetrominoRegistry[tetrominoType](blocks)
}

func (t *tetromino) Draw(targetImage *ebiten.Image) {
	for _, block := range t.blocks {
		block.Draw(targetImage)
	}
}

func (t *tetromino) Translate(x, y float64) {
	t.updateFocalPoint(x, y)
	for _, block := range t.blocks {
		block.Translate(x, y)
	}
}

func (t *tetromino) MoveDown(y float64) {
	t.Translate(0.0, y)
}

func (t *tetromino) Rotate(angle float64) {
	for _, block := range t.blocks {
		block.Rotate(angle, t.focalPoint)
	}
}

func (t *tetromino) SetPosition(x, y float64) {
	oldPosition := t.blocks[0].GetPosition()
	dX := x - oldPosition.X
	dY := y - oldPosition.Y

	for _, block := range t.blocks {
		block.Translate(dX, dY)
	}
}

func (t *tetromino) updateFocalPoint(x, y float64) {
	t.focalPoint.X += x
	t.focalPoint.Y += y
}

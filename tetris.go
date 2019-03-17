package main

import (
	"fmt"
	"log"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	ScreenWidth    = 800
	ScreenHeight   = 600
	GridBlockWidth = 20

	GridRows    = 23
	GridColumns = 10
)

func initGame() func(*ebiten.Image) error {

	grid := NewGrid(200, 20, GridRows, GridColumns, GridBlockWidth)
	grid.AddTetromino()

	return func(screen *ebiten.Image) error {

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			grid.MoveTetrominoRight()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			grid.MoveTetrominoLeft()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			grid.RotateTetrominoRight()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			grid.RotateTetrominoLeft()
		} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			grid.MoveTetromino()
		}

		if ebiten.IsDrawingSkipped() {
			// When the game is running slowly, the rendering result
			// will not be adopted.
			return nil
		}

		screen.Fill(colornames.White)
		grid.DrawGrid(screen)
		// Write your game's rendering.
		debugOutput := fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS())
		ebitenutil.DebugPrint(screen, debugOutput)

		return nil
	}
}

func main() {
	if err := ebiten.Run(initGame(), ScreenWidth, ScreenHeight, 1, "Title"); err != nil {
		log.Fatal(err)
	}
}

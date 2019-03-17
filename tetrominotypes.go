package main

type TetrominoType string

var (
	TetrominoI = TetrominoType(RegisterTetromino("I", createITetromino))
	// TetrominoT =
)

func createITetromino(blocks []*Block) Tetromino {
	var focalX float64 = float64(blocks[0].GetWidth() + 100)
	focalY := float64(blocks[0].GetHeight() * 2)
	blockY := 0
	for _, block := range blocks {
		block.Translate(100.0, (float64)(blockY))
		blockY += block.GetHeight()
	}

	focalPoint := point{X: focalX, Y: focalY}
	return &tetromino{blocks: blocks, focalPoint: focalPoint}
}

func createTTetromino(blocks []*Block) Tetromino {
	blockX := 0
	for i, block := range blocks {

		if i != 3 {
			block.Translate(float64(blockX), 0.0)
		}
	}

	return nil
}

package main

type TetrominoType string

var (
	TetrominoI = TetrominoType(RegisterTetromino("I", createITetromino))
	// TetrominoT =
)

func createITetromino(blocks []*Block) Tetromino {
	var focalX float64 = float64(blocks[0].GetWidth() * 2)
	focalY := float64(blocks[0].GetHeight())
	blockY := 0
	for _, block := range blocks {
		block.Translate((float64)(blockY), 0.0)
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

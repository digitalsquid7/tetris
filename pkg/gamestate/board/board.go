package board

import (
	"time"

	"github.com/digitalsquid7/tetris/pkg/gamestate/coordinate"
	"github.com/digitalsquid7/tetris/pkg/gamestate/tetromino"
	"github.com/digitalsquid7/tetris/pkg/tetrissprites"
)

type Block struct {
	Sprite  tetrissprites.Name
	Opacity uint8
}

func NewBlock(sprite tetrissprites.Name) *Block {
	return &Block{Sprite: sprite, Opacity: 100}
}

type Board struct {
	Blocks         [20][10]*Block
	ClearStartTime time.Time
	ClearEndTime   time.Time
	FullRows       []int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) FreeSpace(coors []coordinate.Coordinate) bool {
	for _, coor := range coors {
		if coor.Y() < 0 || coor.Y() > 19 || coor.X() < 0 || coor.X() > 9 || b.Blocks[coor.Y()][coor.X()] != nil {
			return false
		}
	}

	return true
}

func (b *Board) Collision(coors []coordinate.Coordinate) bool {
	for _, coor := range coors {
		if coor.Y() < 0 || coor.Y() > 19 || coor.X() < 0 || coor.X() > 9 {
			continue
		}

		if b.Blocks[coor.Y()][coor.X()] != nil {
			return true
		}
	}

	return false
}

func (b *Board) InSideBounds(coors []coordinate.Coordinate) bool {
	for _, coor := range coors {
		if coor.X() < 0 || coor.X() > 9 {
			return false
		}
	}

	return true
}

func (b *Board) LockInPlace(tetromino *tetromino.Tetromino) {
	for _, coor := range tetromino.Coordinates() {
		if coor.Y() > -1 {
			b.Blocks[coor.Y()][coor.X()] = NewBlock(tetromino.BlockSprite())
		}
	}
}

func (b *Board) ClearLines() int {
	b.ClearStartTime = time.Now()
	b.ClearEndTime = b.ClearStartTime.Add(time.Millisecond * 500)

nextRow:
	for row := range b.Blocks {
		for col := range b.Blocks[row] {
			if b.Blocks[row][col] == nil {
				continue nextRow
			}
		}
		b.FullRows = append(b.FullRows, row)
	}

	return len(b.FullRows)
}

func (b *Board) UpdateFullRows() bool {
	if len(b.FullRows) == 0 {
		return false
	}

	if time.Now().Before(b.ClearEndTime) {
		b.updateOpacity()
		return true
	}
	b.shiftLines()
	b.FullRows = make([]int, 0)
	return false
}

func (b *Board) updateOpacity() {
	now := time.Now()
	blockDuration := b.ClearEndTime.Sub(b.ClearStartTime) / 10

	for _, fullRow := range b.FullRows {
		for col := range b.Blocks[fullRow] {
			currEndTime := b.ClearStartTime.Add(blockDuration * time.Duration(col+1))
			if currEndTime.Before(time.Now()) {
				b.Blocks[fullRow][col].Opacity = 0
				continue
			}

			percentage := uint8((float64(currEndTime.Sub(now)) /
				float64(currEndTime.Sub(b.ClearStartTime))) * 100)
			b.Blocks[fullRow][col].Opacity = percentage
		}
	}
}

func (b *Board) shiftLines() {
	for _, fullRow := range b.FullRows {
		for row := fullRow - 1; row > -1; row-- {
			b.Blocks[row+1] = b.Blocks[row]
		}
		b.Blocks[0] = [10]*Block{}
	}
}

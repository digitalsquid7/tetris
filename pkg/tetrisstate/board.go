package tetrisstate

import (
	"time"

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

func (b *Board) FreeSpace(coors []Coordinate) bool {
	for _, coor := range coors {
		if coor.Y() < 0 || coor.Y() > 19 || coor.X() < 0 || coor.X() > 9 || b.Blocks[coor.Y()][coor.X()] != nil {
			return false
		}
	}

	return true
}

func (b *Board) InSideBounds(coors []Coordinate) bool {
	for _, coor := range coors {
		if coor.X() < 0 || coor.X() > 9 {
			return false
		}
	}

	return true
}

func (b *Board) LockInPlace(tetromino *Tetromino) {
	for _, coor := range tetromino.Coordinates() {
		b.Blocks[coor.Y()][coor.X()] = NewBlock(tetromino.BlockSprite())
	}
}

func (b *Board) ClearLines() {
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
}

func (b *Board) ClearLinesInProgress() bool {
	return len(b.FullRows) > 0
}

func (b *Board) UpdateFullRows() {
	if time.Now().Before(b.ClearEndTime) {
		b.updateOpacity()
		return
	}
	b.shiftLines()
	b.FullRows = make([]int, 0)
}

func (b *Board) updateOpacity() {
	percentage := uint8((float64(b.ClearEndTime.Sub(time.Now())) / float64(b.ClearEndTime.Sub(b.ClearStartTime))) * 100)

	for _, fullRow := range b.FullRows {
		for col := range b.Blocks[fullRow] {
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

func ref[T any](val T) *T {
	return &val
}

package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type tile int

const (
	blockTile tile = iota
	snakeHeadTile
	foodTile
	emptyTile
)

const (
	tileSize   = 20
	tileMargin = 1
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}

type Board struct {
	size  int
	board [][]tile
	snake *Snake
	food  *Food
}

func NewBoard(size int) *Board {
	board := make([][]tile, size)
	snake := NewSnake()
	for i := 0; i < size; i++ {
		board[i] = make([]tile, size)
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				board[i][j] = blockTile
			} else {
				board[i][j] = emptyTile
			}
		}
	}
	for i, s := range snake.body {
		if i == 0 {
			board[s.y][s.x] = snakeHeadTile
		} else {
			board[s.y][s.x] = blockTile
		}

	}
	food := NewRandomFood(&board)
	board[food.y][food.x] = foodTile
	return &Board{
		size,
		board,
		snake,
		food,
	}
}

func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

func (b *Board) Update() error {
	if err := b.snake.Move(KeyToDir(), b); err != nil {
		return err
	}
	size := b.size
	for i := 0; i < size; i++ {
		b.board[i] = make([]tile, size)
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				b.board[i][j] = blockTile
			} else if b.board[i][j] != snakeHeadTile {
				b.board[i][j] = emptyTile
			}
		}
	}
	b.board[b.food.y][b.food.x] = foodTile
	for i, s := range b.snake.body {
		if i == 0 && b.board[s.y][s.x] != blockTile {
			b.board[s.y][s.x] = snakeHeadTile
		} else {
			b.board[s.y][s.x] = blockTile
		}
	}

	return nil
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			op := &ebiten.DrawImageOptions{}
			x := j*tileSize + (j+1)*tileMargin
			y := i*tileSize + (i+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			if b.board[i][j] == blockTile || b.board[i][j] == snakeHeadTile {
				// log.Println("snake")
				op.ColorScale.ScaleWithColor(snakeColor)
			} else if b.board[i][j] == foodTile {
				// log.Println("food", j, i)
				op.ColorScale.ScaleWithColor(foodColor)
			} else {
				op.ColorScale.ScaleWithColor(tileColorDark)
			}
			boardImage.DrawImage(tileImage, op)
		}
	}
}

package snake

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Food struct {
	x, y int
}

func NewRandomFood(b *[][]tile) *Food {
	availableTiles := []Food{}
	for i := 0; i < len(*b); i++ {
		for j := 0; j < len(*b); j++ {
			if (*b)[i][j] == emptyTile {
				availableTiles = append(availableTiles, Food{j, i})
			}
		}
	}

	result := availableTiles[random.Intn(len(availableTiles))]
	return &result
}

func RefreshFood(b *Board) *Food {
	prevFood := b.food
	b.board[prevFood.y][prevFood.x] = emptyTile
	result := NewRandomFood(&b.board)
	b.food = result
	return result
}

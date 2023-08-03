package snake

import (
	"errors"
)

type snakeBody struct {
	x, y int
}

type Snake struct {
	body []snakeBody
	dir  Dir
}

func NewSnake() *Snake {
	return &Snake{
		[]snakeBody{{4, 4}, {3, 4}},
		DirRight,
	}
}

var errSnakeCollision = errors.New("snake collided")

// Return error if fail to move snake
func (s *Snake) Move(dir Dir, board *Board) error {
	pos := s.body[0]

	switch dir {
	case DirUp:
		if board.food.x == pos.x && board.food.y == pos.y {
			RefreshFood(board)
			s.body = append([]snakeBody{{pos.x, pos.y - 1}}, s.body[:len(s.body)]...)
		} else if board.board[pos.y][pos.x] != blockTile {
			// log.Println(board.board[pos.y][pos.x])
			s.body = append([]snakeBody{{pos.x, pos.y - 1}}, s.body[:len(s.body)-1]...)
		} else {
			return errSnakeCollision
		}
	case DirDown:
		if board.food.x == pos.x && board.food.y == pos.y {
			s.body = append([]snakeBody{{pos.x, pos.y + 1}}, s.body[:len(s.body)]...)
			RefreshFood(board)
		} else if board.board[pos.y][pos.x] != blockTile {
			s.body = append([]snakeBody{{pos.x, pos.y + 1}}, s.body[:len(s.body)-1]...)
		} else {
			return errSnakeCollision
		}
	case DirLeft:
		if board.food.x == pos.x && board.food.y == pos.y {
			s.body = append([]snakeBody{{pos.x - 1, pos.y}}, s.body[:len(s.body)]...)
			RefreshFood(board)
		} else if board.board[pos.y][pos.x] != blockTile {
			s.body = append([]snakeBody{{pos.x - 1, pos.y}}, s.body[:len(s.body)-1]...)
		} else {
			return errSnakeCollision
		}
	case DirRight:
		if board.food.x == pos.x && board.food.y == pos.y {
			s.body = append([]snakeBody{{pos.x + 1, pos.y}}, s.body[:len(s.body)]...)
			RefreshFood(board)
		} else if board.board[pos.y][pos.x] != blockTile {
			s.body = append([]snakeBody{{pos.x + 1, pos.y}}, s.body[:len(s.body)-1]...)
		} else {
			return errSnakeCollision
		}
	}
	return nil

}

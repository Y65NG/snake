package snake

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	boardSize    = 20
)

type State int

const (
	StartMenuState State = iota
	GameState
	PausedState
	EndState
)

type Game struct {
	boardImage *ebiten.Image
	board      *Board
	state      State
}

func NewGame() *Game {
	return &Game{nil, NewBoard(boardSize), StartMenuState}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if err := g.board.Update(); err != nil {
		return err
	}
	// log.Println(g.board.snake.body, g.board.food)
	dt, _ := time.ParseDuration("0.1s")
	time.Sleep(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

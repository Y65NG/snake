package snake

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

var (
	Score        int
	InitialSpeed int = 100
)

type Game struct {
	boardImage *ebiten.Image
	board      *Board
	state      State
}

func NewGame() *Game {
	return &Game{nil, NewBoard(boardSize), StartMenuState}
}

func (g *Game) nextState(state State) {
	g.state = state
}

func (g *Game) restart() {
	g.board = NewBoard(boardSize)
	g.state = StartMenuState
	lastPressed = DirRight
	Score = 0
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	switch g.state {
	case StartMenuState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(GameState)
		}
	case GameState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(PausedState)
		}
		if err := g.board.Update(); err != nil {
			g.nextState(EndState)
		}

		var tArg string
		if mul := Score / 15; mul < 10 {
			tArg = fmt.Sprintf("%vms", InitialSpeed+mul*10)
		} else {
			tArg = fmt.Sprintf("%vms", InitialSpeed+100)
		}

		dt, _ := time.ParseDuration(tArg)
		time.Sleep(dt)
	case PausedState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(GameState)
		}
	case EndState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.restart()
		}
	}

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

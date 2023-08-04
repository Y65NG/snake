package snake

import (
	"fmt"
	"image/color"
	"os"
	"strconv"
	"time"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 800
	boardSize    = 20
)

type State int

const (
	StartMenuState State = iota
	GameState
	PausedState
	EndState
)

var hintImage *ebiten.Image

type Game struct {
	boardImage *ebiten.Image
	scoreImage *ebiten.Image
	board      *Board
	score      int
	opts       GameOpts
	state      State
	ui         *ebitenui.UI
}

type GameOpts struct {
	initialSpeed int
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(boardSize),
		opts:  GameOpts{initialSpeed: 100},
		state: StartMenuState,
	}
}

func (g *Game) nextState(state State) {
	g.state = state
}

func (g *Game) reset() {
	g.board = NewBoard(boardSize)
	g.state = StartMenuState
	lastPressed = DirRight
	g.score = 0
	g.ui = LoadStartMenu(g)
}

func (g *Game) exit() {
	os.Exit(0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	switch g.state {
	case StartMenuState:
		if g.ui == nil {
			g.ui = LoadStartMenu(g)
		}
		g.ui.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(GameState)
		}
	case GameState:
		g.ui = nil
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(PausedState)
		}
		if err := g.board.Update(g); err != nil {
			time.Sleep(time.Second / 2)
			g.nextState(EndState)
		}

		var tArg string
		if mul := g.score / 15; mul < 10 {
			tArg = fmt.Sprintf("%vms", g.opts.initialSpeed+mul*10)
		} else {
			tArg = fmt.Sprintf("%vms", g.opts.initialSpeed+100)
		}

		dt, _ := time.ParseDuration(tArg)
		time.Sleep(dt)
	case PausedState:
		if g.ui == nil {
			g.ui = LoadPausedMenu(g)
		}
		g.ui.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.nextState(GameState)
		}
	case EndState:
		if g.ui == nil {
			g.ui = LoadEndMenu(g)
		}
		g.ui.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	// Draw board
	if g.state == GameState {
		if g.boardImage == nil {
			g.boardImage = ebiten.NewImage(g.board.Size())
		}
		g.board.Draw(g.boardImage)
		op := &ebiten.DrawImageOptions{}
		sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
		bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
		x := (sw - bw) / 2
		y := (sh - bh) / 2
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(g.boardImage, op)

		// Draw score
		if g.scoreImage == nil {
			g.scoreImage = ebiten.NewImage(100, 50)
		}
		g.scoreImage.Fill(scoreBackgroundColor)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ScreenWidth-120, 20)
		f, _ := loadFont(30)
		str := strconv.Itoa(g.score)
		x = 40 - (len(str)-1)*8
		y = 35
		text.Draw(g.scoreImage, str, f, x, y, color.White)
		screen.DrawImage(g.scoreImage, op)

		// Draw hint
		if hintImage == nil {
			hintImage = ebiten.NewImage(500, 50)
		}
		hintImage.Fill(backgroundColor)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ScreenWidth-250, ScreenHeight-30)
		f, _ = loadFont(20)
		text.Draw(hintImage, "Press <Space> to PAUSE", f, 10, 20, color.White)
		screen.DrawImage(hintImage, op)
	}

	// Draw UI
	if g.state == StartMenuState || g.state == PausedState || g.state == EndState {
		if g.ui != nil {
			g.ui.Draw(screen)
		}
	}
}

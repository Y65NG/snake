package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

var lastPressed = DirRight

func (d Dir) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

func (d Dir) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

func KeyToDir() Dir {
	if lastPressed != DirDown && (inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW)) {

		lastPressed = DirUp
		return DirUp
	}
	if lastPressed != DirRight && (inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA)) {
		lastPressed = DirLeft
		return DirLeft
	}
	if lastPressed != DirLeft && (inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD)) {
		lastPressed = DirRight
		return DirRight
	}
	if lastPressed != DirUp && (inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS)) {
		lastPressed = DirDown
		return DirDown
	}
	return lastPressed

}

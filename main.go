package main

import (
	"log"
	"snake/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := snake.NewGame()

	ebiten.SetWindowSize(720, 540)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"snake/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := snake.NewGame()

	// ebiten.SetWindowSize(1280, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("贪吃蛇")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

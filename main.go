package main

import (
	"space-shooter/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *player.Player
}

func (game *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return ebiten.WindowSize()
}

func (game *Game) Update() error {
	game.player.Update()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.player.Draw(screen)
}

func main() {
	game := &Game{
		player: player.New(0, 0),
	}

	ebiten.SetWindowSize(1920, 1080)

	windowWidth, windowHeight := ebiten.WindowSize()
	game.player.Teleport(float32(windowWidth/2), float32(windowHeight/2))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}

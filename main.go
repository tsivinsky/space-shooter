package main

import (
	"bytes"
	"space-shooter/assets"
	"space-shooter/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *player.Player

	healthSprite *ebiten.Image
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
	game.drawHealth(screen)
}

func main() {
	healthSprite, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.HealthSource))
	if err != nil {
		panic(err)
	}

	game := &Game{
		player:       player.New(0, 0),
		healthSprite: healthSprite,
	}

	ebiten.SetWindowSize(1920, 1080)

	windowWidth, windowHeight := ebiten.WindowSize()
	game.player.Teleport(float32(windowWidth/2), float32(windowHeight/2))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}

package main

import (
	"bytes"
	"space-shooter/assets"
	"space-shooter/enemy"
	"space-shooter/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player  *player.Player
	enemies []*enemy.Enemy
	count   int

	healthSprite *ebiten.Image
}

func (game *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return ebiten.WindowSize()
}

func (game *Game) Update() error {
	game.count++

	game.player.Update()

	playerX, playerY := game.player.Coords()
	for _, enemy := range game.enemies {
		enemy.Update(playerX, playerY)
	}

	if game.count%240 == 0 {
		game.spawnEnemy()
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.player.Draw(screen)

	for _, enemy := range game.enemies {
		enemy.Draw(screen)
	}

	game.drawHealth(screen)
}

func (game *Game) spawnEnemy() {
	game.enemies = append(game.enemies, enemy.New())
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

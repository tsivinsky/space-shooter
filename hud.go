package main

import "github.com/hajimehoshi/ebiten/v2"

func (game *Game) drawHealth(screen *ebiten.Image) {
	_, screenHeight := ebiten.WindowSize()

	for i := 0; i < int(game.player.Health()); i++ {
		opts := ebiten.DrawImageOptions{}
		img := game.healthSprite
		opts.GeoM.Scale(0.096, 0.096) // so it would be 48x48, original size is 500x500
		opts.GeoM.Translate(float64(10+i*48), float64(screenHeight-58))
		screen.DrawImage(img, &opts)
	}
}

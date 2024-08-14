package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (player *Player) Update() {
	player.frameCount++

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.x -= float32(player.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.x += float32(player.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.y -= float32(player.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.y += float32(player.speed)
	}

	// if I want to show animation for ship while shooting, I will need to use `StartShooting` and `StopShooting` methods and cycle through sprites is isShooting == true
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.frameCount%15 == 0 {
		player.Shoot()
	}

	for _, bullet := range player.bullets {
		bullet.Update()
	}
}

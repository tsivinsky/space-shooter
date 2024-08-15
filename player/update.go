package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (player *Player) Update() {
	screenWidth, screenHeight := ebiten.WindowSize()

	player.frameCount++

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.x = max(player.x-float32(player.speed), 16)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.x = min(player.x+float32(player.speed), float32(screenWidth-16))
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.y = max(player.y-float32(player.speed), 12)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		const healthBarHeight = 48 + 10 // height + offset from border
		player.y = min(player.y+float32(player.speed), float32(screenHeight-healthBarHeight-16))
	}

	// if I want to show animation for ship while shooting, I will need to use `StartShooting` and `StopShooting` methods and cycle through sprites is isShooting == true
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.frameCount%15 == 0 {
		player.Shoot()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		player.Heal()
	}

	for _, bullet := range player.bullets {
		bullet.Update()
	}
}

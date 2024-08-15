package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (player *Player) Update() {
	screenWidth, screenHeight := ebiten.WindowSize()

	player.frameCount++

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.angle -= 4.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.angle += 4.0
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.x += float32(player.speed) * float32(math.Cos((player.angle-90)*(math.Pi/180)))
		player.y += float32(player.speed) * float32(math.Sin((player.angle-90)*(math.Pi/180)))

		player.x = max(player.x, 16)
		player.x = min(player.x, float32(screenWidth-16))
		const healthBarHeight = 48 + 10 // height + offset from border
		player.y = max(player.y, 12)
		player.y = min(player.y, float32(screenHeight-healthBarHeight-16))
	}

	// if I want to show animation for ship while shooting, I will need to use `StartShooting` and `StopShooting` methods and cycle through sprites is isShooting == true
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.frameCount%10 == 0 {
		player.Shoot()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		player.Heal()
	}

	for _, bullet := range player.bullets {
		bullet.Update()
	}
}

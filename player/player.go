package player

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x          float32
	y          float32
	speed      uint
	bullets    []*Bullet
	frameCount uint

	baseSprite         *ebiten.Image
	engineSprite       *ebiten.Image
	engineEffectSprite *ebiten.Image
	weaponSprite       *ebiten.Image
	bulletSprite       *ebiten.Image
}

func (player *Player) Teleport(x, y float32) {
	player.x = x
	player.y = y
}

func (player *Player) Shoot() {
	leftBullet := NewBullet(player.x-7, player.y-10, player.bulletSprite)
	rightBullet := NewBullet(player.x+7, player.y-10, player.bulletSprite)
	player.bullets = append(player.bullets, leftBullet, rightBullet)
}

func New(x, y float32) *Player {
	player := &Player{
		x:     x,
		y:     y,
		speed: 10,
	}

	if err := player.loadSprites(); err != nil {
		panic(err)
	}

	return player
}

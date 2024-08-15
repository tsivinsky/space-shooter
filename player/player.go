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
	health     uint
	angle      float64

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
	// TODO: fix spawn coords for bullets according to angle
	x1, y1 := player.x-7, player.y-10
	x2, y2 := player.x+7, player.y-10
	leftBullet := NewBullet(x1, y1, player.angle, player.bulletSprite)
	rightBullet := NewBullet(x2, y2, player.angle, player.bulletSprite)
	player.bullets = append(player.bullets, leftBullet, rightBullet)
}

func (player *Player) Health() uint {
	return player.health
}

func (player *Player) Heal() {
	player.health++
}

func New(x, y float32) *Player {
	player := &Player{
		x:      x,
		y:      y,
		speed:  10,
		health: 3,
	}

	if err := player.loadSprites(); err != nil {
		panic(err)
	}

	return player
}

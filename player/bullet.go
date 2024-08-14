package player

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	x          float32
	y          float32
	count      int
	frameCount int
	speed      int

	sprite *ebiten.Image
}

func (bullet *Bullet) Update() {
	bullet.count++

	bullet.y -= float32(bullet.speed)
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	i := (bullet.count / 5) % bullet.frameCount
	sx, sy := 0+i*32, 0
	img := bullet.sprite.SubImage(image.Rect(sx, sy, sx+32, sy+32)).(*ebiten.Image)
	imgSize := img.Bounds().Size()
	opts.GeoM.Translate(-float64(imgSize.X)/2, -float64(imgSize.Y)/2)
	opts.GeoM.Translate(float64(bullet.x), float64(bullet.y))

	screen.DrawImage(img, &opts)
}

func NewBullet(x, y float32, sprite *ebiten.Image) *Bullet {
	return &Bullet{
		x:          x,
		y:          y,
		sprite:     sprite,
		frameCount: 4,
		speed:      5,
	}
}

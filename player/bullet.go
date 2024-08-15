package player

import (
	"image"
	"math"
	"space-shooter/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	x          float32
	y          float32
	count      int
	frameCount int
	speed      int
	angle      float64

	sprite *ebiten.Image
}

func (bullet *Bullet) Update() {
	bullet.count++

	bullet.x += float32(bullet.speed) * float32(math.Cos((bullet.angle-90)*(math.Pi/180)))
	bullet.y += float32(bullet.speed) * float32(math.Sin((bullet.angle-90)*(math.Pi/180)))
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	i := (bullet.count / 5) % bullet.frameCount
	sx, sy := 0+i*32, 0
	img := bullet.sprite.SubImage(image.Rect(sx, sy, sx+32, sy+32)).(*ebiten.Image)
	imgSize := img.Bounds().Size()

	opts.GeoM.Translate(-float64(imgSize.X)/2, -float64(imgSize.Y)/2)
	opts.GeoM.Rotate(util.DegreesToRadians(bullet.angle))
	opts.GeoM.Translate(float64(bullet.x), float64(bullet.y))

	screen.DrawImage(img, &opts)
}

func NewBullet(x, y float32, angle float64, sprite *ebiten.Image) *Bullet {
	return &Bullet{
		x:          x,
		y:          y,
		sprite:     sprite,
		frameCount: 4,
		speed:      20,
		angle:      angle,
	}
}

package enemy

import (
	"space-shooter/util"

	"github.com/hajimehoshi/ebiten/v2"
)

func (enemy *Enemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	img := enemy.sprite
	imgSize := img.Bounds().Size()
	opts.GeoM.Translate(-float64(imgSize.X)/2, -float64(imgSize.Y)/2)
	opts.GeoM.Rotate(util.DegreesToRadians(enemy.angle))
	opts.GeoM.Translate(float64(enemy.x), float64(enemy.y))

	screen.DrawImage(img, &opts)
}

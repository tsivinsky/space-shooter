package enemy

import (
	"bytes"
	"space-shooter/assets"
	"space-shooter/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	x     float32
	y     float32
	speed uint
	angle float64

	sprite *ebiten.Image
}

func New() *Enemy {
	w, h := ebiten.WindowSize()
	x, y := util.GetRandomCoords(w, h)

	enemy := &Enemy{
		x:     float32(x),
		y:     float32(y),
		speed: 2,
	}

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.EnemyFrigateBaseSource))
	if err != nil {
		panic(err)
	}
	enemy.sprite = img

	return enemy
}

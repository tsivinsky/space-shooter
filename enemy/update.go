package enemy

import (
	"math"
	"space-shooter/util"
)

func (enemy *Enemy) Update(x, y float32) {
	xd := enemy.x - x
	yd := enemy.y - y

	rad := math.Atan2(float64(yd), float64(xd))
	angle := util.RadiansToDegress(util.RadiansToDegress(rad)) // oh no, there are 2 of them
	angle = math.Mod(angle, 360)

	enemy.angle = angle

	enemy.x += float32(enemy.speed) * float32(math.Cos((enemy.angle-90)*(math.Pi/180)))
	enemy.y += float32(enemy.speed) * float32(math.Sin((enemy.angle-90)*(math.Pi/180)))
}

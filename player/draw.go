package player

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (player *Player) Draw(screen *ebiten.Image) {
	baseOpts := ebiten.DrawImageOptions{}
	baseImage := player.baseSprite
	baseImageSize := baseImage.Bounds().Size()
	baseOpts.GeoM.Translate(-float64(baseImageSize.X)/2, -float64(baseImageSize.Y)/2)
	baseOpts.GeoM.Translate(float64(player.x), float64(player.y))

	engineOpts := ebiten.DrawImageOptions{}
	engineImage := player.engineSprite
	engineImageSize := engineImage.Bounds().Size()
	engineOpts.GeoM.Translate(-float64(engineImageSize.X)/2, -float64(engineImageSize.Y)/2)
	engineOpts.GeoM.Rotate(180 * (math.Pi / 180))
	engineOpts.GeoM.Translate(float64(player.x), float64(player.y+12))

	effectOpts := ebiten.DrawImageOptions{}
	effectImage := player.engineEffectSprite.SubImage(image.Rect(0, 0, 48, 48)).(*ebiten.Image)
	effectImageSize := effectImage.Bounds().Size()
	effectOpts.GeoM.Translate(-float64(effectImageSize.X)/2, -float64(effectImageSize.Y)/2)
	effectOpts.GeoM.Translate(float64(player.x), float64(player.y+2))

	weaponOpts := ebiten.DrawImageOptions{}
	weaponImage := player.weaponSprite.SubImage(image.Rect(0, 0, 48, 48)).(*ebiten.Image)
	weaponImageSize := weaponImage.Bounds().Size()
	weaponOpts.GeoM.Translate(-float64(weaponImageSize.X)/2, -float64(weaponImageSize.Y)/2)
	weaponOpts.GeoM.Translate(float64(player.x), float64(player.y-1))

	screen.DrawImage(effectImage, &effectOpts)
	screen.DrawImage(engineImage, &engineOpts)
	screen.DrawImage(baseImage, &baseOpts)
	screen.DrawImage(weaponImage, &weaponOpts)

	for _, bullet := range player.bullets {
		bullet.Draw(screen)
	}
}

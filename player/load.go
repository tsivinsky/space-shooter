package player

import (
	"bytes"
	"space-shooter/assets"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (player *Player) loadSprites() (err error) {
	player.baseSprite, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ShipBaseFullHealthSource))
	if err != nil {
		return
	}

	player.engineSprite, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ShipEngineSource))
	if err != nil {
		return
	}

	player.engineEffectSprite, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ShipEngineEffectSource))
	if err != nil {
		return
	}

	player.weaponSprite, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ShipWeaponSource))
	if err != nil {
		return
	}

	player.bulletSprite, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.AutocannonBulletSource))
	if err != nil {
		return
	}

	return nil
}

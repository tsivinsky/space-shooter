package assets

import _ "embed"

var (
	//go:embed sprites/ship/base_fullhealth.png
	ShipBaseFullHealthSource []byte

	//go:embed sprites/ship/engine.png
	ShipEngineSource []byte

	//go:embed sprites/ship/engine_effect.png
	ShipEngineEffectSource []byte

	//go:embed sprites/ship/weapon_autocannon.png
	ShipWeaponSource []byte

	//go:embed sprites/ship/autocannon_bullet.png
	AutocannonBulletSource []byte

	//go:embed sprites/health.png
	HealthSource []byte

	//go:embed sprites/enemies/battlecruiser_base.png
	EnemyBattleCruiserBaseSource []byte

	//go:embed sprites/enemies/frigate_base.png
	EnemyFrigateBaseSource []byte
)

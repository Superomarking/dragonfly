package item

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/world/sound"
	"math"
	"time"
)

// Trident is a melee ranged weapon that rarely drops from the Drowned mob
type Trident struct{}

// EncodeItem ...
func (Trident) EncodeItem() (name string, meta int16) {
	return "minecraft:trident", 0
}

// MaxCount ...
func (Trident) MaxCount() int {
	return 1
}

// DurabilityInfo ...
func (Trident) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability:    251,
		BrokenItem:       simpleItem(Stack{}),
		AttackDurability: 1,
		BreakDurability:  2,
		Persistent:       true,
	}
}

func (Trident) Release(releaser Releaser, duration time.Duration, ctx *UseContext) {
	trident, _ := releaser.HeldItems()
	for _, enchant := range trident.Enchantments() {
		if _, ok := enchant.Type().(interface{ OnlyUnderwater() bool }); ok {
			// Riptide shouldn't throw a projectile.
			//TODO: Riptide sounds
			return
		}
	}

	creative := releaser.GameMode().CreativeInventory()
	ticks := duration.Milliseconds() / 50
	if ticks < 5 {
		// The player must hold the trident for at least 5 ticks.
		return
	}

	d := float64(ticks) / 20
	force := math.Min((d*d+d*2)/3, 1)
	if force < 0.5 {
		// The force must be at least 0.5 for the trident to be fully charged
		return
	}

	rot := releaser.Rotation()
	rot = cube.Rotation{-rot[0], -rot[1]}
	if rot[0] > 180 {
		rot[0] = 360 - rot[0]
	}

	if !creative {
		ctx.DamageItem(1)
		ctx.Consume(trident.Grow(-trident.Count() + 1))
	}

	create := releaser.World().EntityRegistry().Config().Trident
	projectile := create(eyePosition(releaser), releaser.Rotation().Vec3().Mul(force*5), rot, releaser, true)

	projectile.Type().(entity.ThrownTridentType).SetTridentItem(trident)

	releaser.PlaySound(sound.TridentThrow{})
	releaser.World().AddEntity(projectile)
}

// Requirements ...
func (Trident) Requirements() []Stack {
	return []Stack{NewStack(Trident{}, 1)}
}

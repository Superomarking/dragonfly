package entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/internal/nbtconv"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

var tridentStack item.Stack

// NewThrownTrident creates a ThrownTrident entity. Trident is a throwable melee ranged weapon
func NewThrownTrident(pos mgl64.Vec3, rot cube.Rotation, owner world.Entity) *Ent {
	trident := Config{Behaviour: ThrownTridentConf.New(owner)}.New(ThrownTridentType{}, pos)
	trident.rot = rot
	return trident
}

var ThrownTridentConf = ProjectileBehaviourConfig{
	Gravity:               0.05,
	Drag:                  0.01,
	Damage:                8.0,
	Critical:              false,
	Hit:                   ReturnOrStrike,
	SurviveBlockCollision: true,
	DisablePickup:         true,
}

// ThrownTridentType is a world.EntityType implementation for Trident.
type ThrownTridentType struct{}

// EncodeEntity ...
func (ThrownTridentType) EncodeEntity() string {
	return "minecraft:thrown_trident"
}

// BBox ...
func (ThrownTridentType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

// ReturnOrStrike checks for the loyalty and channeling enchantment
func ReturnOrStrike(e *Ent, target trace.Result) {
	for _, enchant := range tridentStack.Enchantments() {
		if _, ok := enchant.Type().(interface{ AlwaysReturns() bool }); ok {
			// TODO: Loyalty
		}
		if _, ok := enchant.Type().(interface{ StrikesMobs() bool }); ok {
			// TODO: Channeling
		}
	}
}

// DecodeNBT ...
func (ThrownTridentType) DecodeNBT(m map[string]any) world.Entity {
	trident := nbtconv.MapItem(m, "Trident")
	tr := NewThrownTrident(nbtconv.Vec3(m, "Pos"), nbtconv.Rotation(m), nil)

	b := tr.conf.Behaviour.(*ProjectileBehaviour)
	b.conf.PickupItem = trident
	tr.vel = nbtconv.Vec3(m, "Motion")
	return tr
}

// EncodeNBT ...
func (ThrownTridentType) EncodeNBT(e world.Entity) map[string]any {
	tr := e.(*Ent)
	yaw, pitch := tr.Rotation().Elem()
	return map[string]any{
		"Pos":     nbtconv.Vec3ToFloat32Slice(tr.Position()),
		"Yaw":     float32(yaw),
		"Pitch":   float32(pitch),
		"Motion":  nbtconv.Vec3ToFloat32Slice(tr.Velocity()),
		"Trident": nbtconv.WriteItem(tridentStack, true),
		// This is used to specify the slot that the trident was held at when throwing
		"favoredSlot": 0,
	}
}

// SetTridentItem sets the trident item which was thrown for the projectile nbt
func (ThrownTridentType) SetTridentItem(t item.Stack) {
	tridentStack = t
}

// GetTridentItem returns the trident item which was thrown
func (ThrownTridentType) GetTridentItem() item.Stack {
	return tridentStack
}

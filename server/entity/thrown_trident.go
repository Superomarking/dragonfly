package entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/internal/nbtconv"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// NewThrownTrident creates a ThrownTrident entity. Trident is a throwable melee ranged weapon
func NewThrownTrident(pos mgl64.Vec3, rot cube.Rotation, owner world.Entity) *Ent {
	a := Config{Behaviour: ThrownTridentConf.New(owner)}.New(ThrownTridentType{}, pos)
	a.rot = rot
	return a
}

var ThrownTridentConf = ProjectileBehaviourConfig{
	Gravity:               0.05,
	Drag:                  0.01,
	Damage:                8.0,
	Critical:              false,
	Hit:                   hit,
	SurviveBlockCollision: true,
	PickupItem:            item.NewStack(item.Trident{}, 1),
}

// ThrownTridentType is a world.EntityType implementation for Trident.
type ThrownTridentType struct{}

// EncodeEntity ...
func (ThrownTridentType) EncodeEntity() string { return "minecraft:thrown_trident" }

// BBox ...
func (ThrownTridentType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

// hit checks for the loyalty and channeling enchantment
func hit(e *Ent, target trace.Result) {
	//TODO: loyalty
	//TODO: channeling
}

// DecodeNBT ...
func (ThrownTridentType) DecodeNBT(m map[string]any) world.Entity {
	ep := NewThrownTrident(nbtconv.Vec3(m, "Pos"), nbtconv.Rotation(m), nil)
	ep.vel = nbtconv.Vec3(m, "Motion")
	return ep
}

// EncodeNBT ...
func (ThrownTridentType) EncodeNBT(e world.Entity) map[string]any {
	ep := e.(*Ent)
	a := e.(*Ent)
	yaw, pitch := a.Rotation().Elem()
	return map[string]any{
		"Pos":    nbtconv.Vec3ToFloat32Slice(ep.Position()),
		"Yaw":    float32(yaw),
		"Pitch":  float32(pitch),
		"Motion": nbtconv.Vec3ToFloat32Slice(ep.Velocity()),
		//TODO: loyalty
		"favoredSlot": 0,
	}
}

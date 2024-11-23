package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Riptide is a trident enchantment which hurls the user in the direction the user is facing,
// but only when they are wet.
type Riptide struct{}

// Name ...
func (Riptide) Name() string {
	return "Riptide"
}

// MaxLevel ...
func (Riptide) MaxLevel() int {
	return 3
}

// Cost ...
func (Riptide) Cost(level int) (int, int) {
	min := 10 + level*7
	return min, 50
}

// Rarity ...
func (Riptide) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (Riptide) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	_, loyalty := t.(Loyalty)
	_, channeling := t.(Channeling)
	return !loyalty && !channeling
}

// CompatibleWithItem ...
func (Riptide) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Trident)
	return ok
}

// OnlyUnderwater always returns true
func (Riptide) OnlyUnderwater() bool {
	return true
}

package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Loyalty is a trident enchantment which returns to the owner once thrown
type Loyalty struct{}

// Name ...
func (Loyalty) Name() string {
	return "Loyalty"
}

// MaxLevel ...
func (Loyalty) MaxLevel() int {
	return 3
}

// Cost ...
func (Loyalty) Cost(level int) (int, int) {
	min := 5 + level*7
	return min, 50
}

// Rarity ...
func (Loyalty) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityUncommon
}

// CompatibleWithEnchantment ...
func (Loyalty) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	_, riptide := t.(Riptide)
	return !riptide
}

// CompatibleWithItem ...
func (Loyalty) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Trident)
	return ok
}

// AlwaysReturns ...
func (Loyalty) AlwaysReturns() bool {
	return true
}

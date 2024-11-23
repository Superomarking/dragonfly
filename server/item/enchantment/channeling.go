package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Channeling is a trident enchantment that produces lightning when thrown at a mob or lightning rod while a
// thunderstorm is occurring.
type Channeling struct{}

// Name ...
func (Channeling) Name() string {
	return "Channeling"
}

// MaxLevel ...
func (Channeling) MaxLevel() int {
	return 1
}

// Cost ...
func (Channeling) Cost(int) (int, int) {
	return 25, 50
}

// Rarity ...
func (Channeling) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityVeryRare
}

// CompatibleWithEnchantment ...
func (Channeling) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	_, riptide := t.(Riptide)
	return !riptide
}

// CompatibleWithItem ...
func (Channeling) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Trident)
	return ok
}

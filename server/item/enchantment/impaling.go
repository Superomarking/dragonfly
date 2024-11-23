package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Impaling is a trident enchantment which deals extra damage to mobs in water
type Impaling struct{}

// Name ...
func (Impaling) Name() string {
	return "Impaling"
}

// MaxLevel ...
func (Impaling) MaxLevel() int {
	return 5
}

// Cost ...
func (Impaling) Cost(level int) (int, int) {
	min := 1 + (level-1)*8
	return min, min + 20
}

// Rarity ...
func (Impaling) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (Impaling) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (Impaling) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Trident)
	return ok
}

// BonusDamage is the extra amount of damage the Impaling enchantment adds when
// attacking mobs that are touching water
func (Impaling) BonusDamage(level int) float64 {
	return float64(level) * 2.5
}

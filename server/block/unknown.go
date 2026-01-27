package block

import (
	"github.com/df-mc/dragonfly/server/item"
)

// Unknown is a block that does not exist but is not air
type Unknown struct {
	solid
}

// EncodeBlock ...
func (Unknown) EncodeBlock() (string, map[string]any) {
	return "minecraft:unknown", nil
}

// EncodeItem ...
func (Unknown) EncodeItem() (name string, meta int16) {
	return "minecraft:unknown", 0
}

func (u Unknown) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0,
		Harvestable: alwaysHarvestable,
		Effective:   nothingEffective,
		Drops: func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
			return nil
		},
	}
}

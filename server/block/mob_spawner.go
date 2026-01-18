package block

import (
	"strings"

	"github.com/df-mc/dragonfly/server/item"
)

// MobSpawner is a decorative block that looks like a spawner.
type MobSpawner struct {
	transparent
	solid

	// EntityIdentifier is the text ID of the mob, e.g., "minecraft:blaze"
	EntityIdentifier string
}

// NewSpawner creates a spawner with a specific mob inside.
func NewSpawner(entityType string) MobSpawner {
	return MobSpawner{EntityIdentifier: entityType}
}

// EncodeNBT tells the client how to render the block and the entity inside.
func (s MobSpawner) EncodeNBT() map[string]any {
	data := map[string]any{
		"id": "MobSpawner",
	}

	if s.EntityIdentifier != "" {
		data["EntityIdentifier"] = s.EntityIdentifier
		data["DisplayEntityWidth"] = float32(0.6)
		data["DisplayEntityHeight"] = float32(1.8)
		data["DisplayEntityScale"] = float32(0.7)

		// Client-side animation timers
		data["Delay"] = int16(20)
		data["MinSpawnDelay"] = int16(200)
		data["MaxSpawnDelay"] = int16(800)
		data["RequiredPlayerRange"] = int16(16)
	}

	return data
}

// DecodeNBT loads the mob type from saved data.
func (s MobSpawner) DecodeNBT(data map[string]any) any {
	if v, ok := data["EntityIdentifier"].(string); ok {
		s.EntityIdentifier = v
	} else if v, ok := data["EntityId"].(string); ok {
		s.EntityIdentifier = v
	}
	return s
}

// BreakInfo ensures the block drops itself with the mob data preserved.
func (s MobSpawner) BreakInfo() BreakInfo {
	st := item.NewStack(s, 1)

	// Ensure the dropped item has the correct display name
	st = st.WithCustomName(s.DisplayName())

	// You must return the stack in a slice within the drops function
	return newBreakInfo(5, alwaysHarvestable, pickaxeEffective, func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
		return []item.Stack{st}
	})
}

func (s MobSpawner) DisplayName() string {
	if s.EntityIdentifier == "" {
		return "§rEmpty Spawner"
	}

	name := strings.TrimPrefix(s.EntityIdentifier, "minecraft:")
	name = strings.ReplaceAll(name, "_", " ")

	words := strings.Split(name, " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return "§r" + strings.Join(words, " ") + " Spawner"
}

// EncodeItem returns the vanilla item ID.
func (s MobSpawner) EncodeItem() (name string, meta int16) {
	return "minecraft:mob_spawner", 0
}

// EncodeBlock returns the vanilla block ID.
func (s MobSpawner) EncodeBlock() (string, map[string]any) {
	return "minecraft:mob_spawner", nil
}

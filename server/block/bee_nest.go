package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// BeeNest is a full-sized block of bee nest.
type BeeNest struct {
	solid
	// Facing represents the direction the bee nest is facing.
	Facing cube.Direction
}

// UseOnBlock handles the placement of the bee nest.
func (b BeeNest) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, b)
	if !used {
		return false
	}

	// We set the direction to the player's looking direction so the
	// front of the nest faces the player upon placement.
	b.Facing = user.Rotation().Direction().Opposite()

	place(tx, pos, b, user, ctx)
	return placed(ctx)
}

// BreakInfo defines the hardness and harvest tool for the bee nest.
func (b BeeNest) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, axeEffective, oneOf(b)).withBlastResistance(1.5)
}

// FlammabilityInfo defines how the block burns.
func (b BeeNest) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(30, 30, true)
}

// EncodeItem ...
func (BeeNest) EncodeItem() (name string, meta int16) {
	return "minecraft:bee_nest", 0
}

// EncodeAll returns all 4 possible directional states for the Bee Nest.
func (BeeNest) EncodeAll() []world.Block {
	var all []world.Block
	// This ensures every direction is registered with honey_level 0
	for _, d := range cube.Directions() {
		all = append(all, BeeNest{Facing: d})
	}
	return all
}

// EncodeBlock encodes the block states for Bedrock Edition.
func (b BeeNest) EncodeBlock() (string, map[string]any) {
	// Map Dragonfly cube.Direction to Bedrock's specific 0-3 metadata bits
	var dir int32
	switch b.Facing {
	case cube.South:
		dir = 0
	case cube.West:
		dir = 1
	case cube.North:
		dir = 2
	case cube.East:
		dir = 3
	}

	return "minecraft:bee_nest", map[string]any{
		"direction":   dir,      // Corrected mapping
		"honey_level": int32(0), // Keeps it empty and non-functional
	}
}

func allBeeNests() []world.Block {
	return BeeNest{}.EncodeAll()
}

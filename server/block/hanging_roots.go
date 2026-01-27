package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// HangingRoots is a hanging root found under azailia trees
type HangingRoots struct {
	transparent
}

// EncodeBlock ...
func (h HangingRoots) EncodeBlock() (string, map[string]any) {
	return "minecraft:hanging_roots", nil
}

// EncodeItem ...
func (h HangingRoots) EncodeItem() (name string, meta int16) {
	return "minecraft:hanging_roots", 0
}

// BreakInfo ...
func (h HangingRoots) BreakInfo() BreakInfo {
	return newBreakInfo(0, func(t item.Tool) bool {
		return t.ToolType() == item.TypeShears
	}, shearsEffective, func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
		if t.ToolType() == item.TypeShears {
			return []item.Stack{item.NewStack(h, 1)}
		}
		return nil
	})
}

// NeighbourUpdateTick ...
func (h HangingRoots) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if !tx.Block(pos.Side(cube.FaceUp)).Model().FaceSolid(pos.Side(cube.FaceUp), cube.FaceUp.Opposite(), tx) {
		breakBlock(h, pos, tx)
	}
}

// UseOnBlock ...
func (h HangingRoots) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) bool {
	pos, _, used := firstReplaceable(tx, pos, face, h)
	if !used {
		return false
	}
	if !tx.Block(pos.Side(cube.FaceUp)).Model().FaceSolid(pos.Side(cube.FaceUp), cube.FaceUp.Opposite(), tx) {
		return false
	}

	place(tx, pos, h, user, ctx)
	return placed(ctx)
}

// Model ...
func (h HangingRoots) Model() world.BlockModel {
	return model.Skull{Direction: cube.FaceDown, Hanging: true}
}

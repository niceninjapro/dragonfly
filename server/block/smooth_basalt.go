package block

// SmoothBasalt is a full-sized block of smooth basalt.
type SmoothBasalt struct {
	solid
}

// BreakInfo ...
func (s SmoothBasalt) BreakInfo() BreakInfo {
	return newBreakInfo(1.25, alwaysHarvestable, pickaxeEffective, oneOf(s)).withBlastResistance(21)
}

// EncodeItem ...
func (SmoothBasalt) EncodeItem() (name string, meta int16) {
	return "minecraft:smooth_basalt", 0
}

// EncodeBlock ...
func (SmoothBasalt) EncodeBlock() (string, map[string]any) {
	return "minecraft:smooth_basalt", nil
}

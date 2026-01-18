package block

// SmoothBasalt is a full-sized block of smooth basalt.
type MossBlock struct {
	solid
}

// BreakInfo ...
func (m MossBlock) BreakInfo() BreakInfo {
	return newBreakInfo(0.1, alwaysHarvestable, hoeEffective, oneOf(m))
}

// EncodeItem ...
func (MossBlock) EncodeItem() (name string, meta int16) {
	return "minecraft:moss_block", 0
}

// EncodeBlock ...
func (MossBlock) EncodeBlock() (string, map[string]any) {
	return "minecraft:moss_block", nil
}

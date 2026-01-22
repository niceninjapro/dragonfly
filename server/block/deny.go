package block

// Deny is a block that is indestructible in survival and protect all the blocks above it.
type Deny struct {
	solid
}

// EncodeItem ...
func (Deny) EncodeItem() (name string, meta int16) {
	return "minecraft:deny", 0
}

// EncodeBlock ...
func (d Deny) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:deny", map[string]any{}
}

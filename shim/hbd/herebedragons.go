package hbd

type (
	ZoneCoord  float32
	WorldCoord float32
	MapID      int
	MapLevel   int
	InstanceID int
)

// WorldPosFromZone returns the world position given a zone position.
func WorldPosFromZone(zx, zy ZoneCoord, id MapID, level MapLevel) (wx, wy WorldCoord, inst InstanceID) {
	return 0, 0, 0
}

// ZonePosFromWorld returns the zone position given a world position.
func ZonePosFromWorld(wx, wy WorldCoord, id MapID, level MapLevel, allowOutOfBounds bool) (zx, zy ZoneCoord) {
	return 0, 0
}

// WorldDistance computes the distance between two positions in the same instance/continent.
func WorldDistance(inst InstanceID, srcX, srcY, dstX, dstY WorldCoord) float32 {
	return 0
}

// WorldDistanceVector is like WorldDistance, except it returns the angle and the distance.
func WorldDistanceVector(inst InstanceID, srcX, srcY, dstX, dstY WorldCoord) (distance, angle float32) {
	return 0, 0
}

// UnitWorldPosition returns the world position of the unit. If the position is not known
// or the unit is invalid, the zero value is returned.
func UnitWorldPosition(unitID string) (x, y WorldCoord, inst InstanceID) {
	return 0, 0, 0
}

// PlayerWorldPosition returns the world position of the player.
func PlayerWorldPosition() (x, y WorldCoord, inst InstanceID) {
	return 0, 0, 0
}

// PlayerZone returns the player's current zone.
func PlayerZone() (MapID, MapLevel) {
	return 0, 0
}

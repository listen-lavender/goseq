package memory

type SoftRegion = map[string]uint16

func NewSoftRegion(softRegion map[string]uint16) SoftRegion {
	return softRegion
}

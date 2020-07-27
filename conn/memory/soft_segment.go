package memory

// SoftSegment 对于seq号的跳跃段内存管理
type SoftSegment = map[uint16]uint64

func NewSoftSegment(softSegment map[uint16]uint64) SoftSegment {
	return softSegment
}

package memory

// seq号内存管理
type Seq = map[string]*uint64

func NewSeq(seq map[string]*uint64) Seq {
	return seq
}

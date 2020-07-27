package model

import (
	"context"
)

//go:generate mockgen -destination=../mock/hard_segment_mock.go -package=mock github.com/listen-lavender/goseq/model HardSegmentDao

// seq号跳跃段持久化存储
type HardSegment struct {
	HardSegmentID uint16 `bson:"_id"`
	MaxSeq        uint64 `bson:"maxSeq"`
}

func (*HardSegment) TableName() string {
	return "seq_sections"
}

type HardSegmentDao interface {
	StoreID() string
	AtomicAdd(ctx context.Context, o *HardSegment) (*HardSegment, error)
	AtomicUpdate(ctx context.Context, uType string, id uint16, o *HardSegment) (*HardSegment, error)
	Find(ctx context.Context, offset uint64, limit int, ftype string, ol []*HardSegment, filter func(*HardSegment) bool) ([]*HardSegment, error)
}

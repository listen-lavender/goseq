package model

import (
	"context"
)

//go:generate mockgen -destination=../mock/soft_segment_mock.go -package=mock github.com/listen-lavender/goseq/model SoftSegmentDao

const (
	STEP = 1
	SKIP = 10000
)

// seq号跳跃段内存管理
type SoftSegmentDao interface {
	CAS(ctx context.Context, regionID uint16, incResult uint64) (uint64, bool, error)
}

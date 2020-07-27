package model

import (
	"context"
)

//go:generate mockgen -destination=../mock/soft_region_mock.go -package=mock github.com/listen-lavender/goseq/model SoftRegionDao

const (
	MAX = 1000
)

// Region 对于namespace进行合并分区间

type SoftRegionDao interface {
	GetSet(ctx context.Context, namespace string) uint16
	Hash(ctx context.Context, namespace string) uint16
}

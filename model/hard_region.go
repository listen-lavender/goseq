package model

import (
	"context"
)

//go:generate mockgen -destination=../mock/hard_region_mock.go -package=mock github.com/listen-lavender/goseq/model HardRegionDao

// seq号的namespace区间持久化存储
type HardRegion struct {
	Namespace string `bson:"_id"` // namespace
	RegionID  uint16 `bson:"region_id"`
}

func (*HardRegion) TableName() string {
	return "ns_regions"
}

type HardRegionDao interface {
	StoreID() string
	AtomicAdd(ctx context.Context, o *HardRegion) (*HardRegion, error)
	FindByID(ctx context.Context, id string) (*HardRegion, error)
	Find(ctx context.Context, offset uint64, limit int, ftype string, ol []*HardRegion, filter func(*HardRegion) bool) ([]*HardRegion, error)
}

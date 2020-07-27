package dao

import (
	"context"
	"runtime/debug"
	"sync"

	"github.com/listen-lavender/goseq/conn/memory"
	"github.com/listen-lavender/goseq/lib"
	"github.com/listen-lavender/goseq/model"
)

type softRegionHandler struct {
	rw            sync.RWMutex
	softRegion    memory.SoftRegion
	hardRegionDao model.HardRegionDao
}

func NewSoftRegionHandler(softRegion memory.SoftRegion, hardRegionDao model.HardRegionDao) *softRegionHandler {
	ctx := context.Background()

	for namespace, regionID := range softRegion {
		if hr, err := hardRegionDao.FindByID(ctx, namespace); err == nil && hr.Namespace != "" {
			continue
		}

		hr := &model.HardRegion{
			Namespace: namespace,
			RegionID:  regionID,
		}
		hardRegionDao.AtomicAdd(ctx, hr)
	}

	hrlist, _ := hardRegionDao.Find(ctx, 0, 0, "", nil, func(*model.HardRegion) bool { return false })
	for _, hr := range hrlist {
		softRegion[hr.Namespace] = hr.RegionID
	}

	return &softRegionHandler{
		softRegion:    softRegion,
		hardRegionDao: hardRegionDao,
	}
}

func (srh *softRegionHandler) GetSet(ctx context.Context, namespace string) uint16 {
	srh.rw.RLock()
	regionID, ok := srh.softRegion[namespace]
	srh.rw.RUnlock()
	if ok {
		return regionID
	}

	regionID = uint16(lib.Murmur32([]byte(namespace)) % model.MAX)

	srh.rw.Lock()
	srh.softRegion[namespace] = regionID
	srh.rw.Unlock()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				println("update member conversation status: ", "error", err, "stack", string(debug.Stack()))
			}
		}()
		hr := &model.HardRegion{
			Namespace: namespace,
			RegionID:  regionID,
		}
		srh.hardRegionDao.AtomicAdd(ctx, hr)
	}()

	return regionID
}

func (rh *softRegionHandler) Hash(ctx context.Context, namespace string) uint16 {
	return 0
}

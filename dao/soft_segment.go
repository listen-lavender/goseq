package dao

import (
	"context"
	"errors"
	"sync"

	"github.com/listen-lavender/goseq/conn/memory"
	"github.com/listen-lavender/goseq/model"
)

type softSegmentHandler struct {
	rw             map[uint16]sync.RWMutex
	softSegment    memory.SoftSegment
	hardSegmentDao model.HardSegmentDao
}

func NewSoftSegmentHandler(softSegment memory.SoftSegment, hardSegmentDao model.HardSegmentDao) *softSegmentHandler {
	rw := make(map[uint16]sync.RWMutex, model.MAX)
	for k := 0; k < model.MAX; k++ {
		rw[uint16(k)] = sync.RWMutex{}
	}

	ctx := context.Background()
	hslist, _ := hardSegmentDao.Find(ctx, 0, 0, "", nil, func(*model.HardSegment) bool { return false })
	for _, hs := range hslist {
		softSegment[hs.HardSegmentID] = hs.MaxSeq
	}

	return &softSegmentHandler{
		rw:             rw,
		softSegment:    softSegment,
		hardSegmentDao: hardSegmentDao,
	}
}

func (ssh *softSegmentHandler) CAS(ctx context.Context, regionID uint16, incResult uint64) (uint64, bool, error) {
	var err error
	l, ok := ssh.rw[regionID]
	if !ok {
		return 0, false, errors.New("Not exist.")
	}

	l.RLock()
	max, _ := ssh.softSegment[regionID]
	l.RUnlock()

	notZero := max != 0

	if max > incResult {
		return max, true, err
	}

	max = max + model.SKIP
	hs := &model.HardSegment{
		HardSegmentID: regionID,
		MaxSeq:        max,
	}

	l.Lock()
	if notZero {
		_, err = ssh.hardSegmentDao.AtomicUpdate(ctx, "", regionID, hs)
	} else {
		_, err = ssh.hardSegmentDao.AtomicAdd(ctx, hs)
	}
	if err == nil {
		ssh.softSegment[regionID] = max
	}
	l.Unlock()

	return max, err == nil, err
}

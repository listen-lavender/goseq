package dao

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/listen-lavender/goseq/conn/memory"
	"github.com/listen-lavender/goseq/model"
)

type seqHandler struct {
	rw  sync.RWMutex
	seq memory.Seq
}

func NewSeqHandler(seq memory.Seq, softRegionHandler *softRegionHandler, softSegmentHandler *softSegmentHandler) *seqHandler {
	for namespace, regionID := range softRegionHandler.softRegion {
		id, _ := softSegmentHandler.softSegment[regionID]
		seq[namespace] = &id
	}
	return &seqHandler{
		seq: seq,
	}
}

func (sh *seqHandler) AtomicInc(ctx context.Context, namespace string) uint64 {
	var k uint64

	sh.rw.RLock()
	o, ok := sh.seq[namespace]
	sh.rw.RUnlock()

	if ok {
		k = atomic.AddUint64(o, model.STEP)
	} else {
		k = uint64(model.STEP)
		o = &k
		sh.rw.Lock()
		xo, xok := sh.seq[namespace]
		if !xok {
			sh.seq[namespace] = o
		}
		sh.rw.Unlock()
		if xok {
			k = atomic.AddUint64(xo, model.STEP)
		}
	}
	return k
}

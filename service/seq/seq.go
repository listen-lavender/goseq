package seq

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/listen-lavender/goseq/api"
	"github.com/listen-lavender/goseq/model"
)

type SeqService struct {
	seqDao         model.SeqDao
	softSegmentDao model.SoftSegmentDao
	softRegionDao  model.SoftRegionDao
}

func NewSeqService(seqDao model.SeqDao, softSegmentDao model.SoftSegmentDao, softRegionDao model.SoftRegionDao) *SeqService {
	rs := &SeqService{
		seqDao:         seqDao,
		softSegmentDao: softSegmentDao,
		softRegionDao:  softRegionDao,
	}
	return rs
}

// 获取
func (ss *SeqService) GetSeq(ctx *gin.Context) {
	req := &api.SeqReq{
		NS: ctx.Param("ns"),
	}
	regionID := ss.softRegionDao.GetSet(ctx, req.NS)
	seq := ss.seqDao.AtomicInc(ctx, req.NS)
	_, _, err := ss.softSegmentDao.CAS(ctx, regionID, seq)

	var res *api.SeqRes

	if err == nil {
		res = &api.SeqRes{
			NS:  req.NS,
			Seq: seq,
			TS:  time.Now().Unix(),
		}
	} else {
		res = &api.SeqRes{
			ErrCode: api.ErrSystemErrorCode,
			Msg:     api.ErrSystemErrorMsg,
		}
	}

	ctx.JSON(200, res)
}

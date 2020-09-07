package seq

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/listen-lavender/goseq/api"
	"github.com/listen-lavender/goseq/api/pbseq"
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
func (ss *SeqService) GetSeq(ctx context.Context, req *pbseq.GetSeqReq) (*pbseq.GetSeqResp, error) {
	regionID := ss.softRegionDao.GetSet(ctx, req.Namespace)
	seq := ss.seqDao.AtomicInc(ctx, req.Namespace)
	_, _, err := ss.softSegmentDao.CAS(ctx, regionID, seq)

	var res *pbseq.GetSeqResp

	if err == nil {
		res = &pbseq.GetSeqResp{
			Namespace: req.Namespace,
			Seq:       seq,
			Ts:        time.Now().Unix(),
		}
	}
	return res, err
}

// http获取
func (ss *SeqService) HttpGetSeq(ctx *gin.Context) {
	req := &api.SeqReq{
		NS: ctx.Param("ns"),
	}
	regionID := ss.softRegionDao.GetSet(ctx, req.NS)
	seq := ss.seqDao.AtomicInc(ctx, req.NS)
	_, _, err := ss.softSegmentDao.CAS(ctx, regionID, seq)

	var res *api.SeqResp

	if err == nil {
		res = &api.SeqResp{
			NS:  req.NS,
			Seq: seq,
			TS:  time.Now().Unix(),
		}
	} else {
		res = &api.SeqResp{
			ErrCode: api.ErrSystemErrorCode,
			Msg:     api.ErrSystemErrorMsg,
		}
	}

	ctx.JSON(200, res)
}

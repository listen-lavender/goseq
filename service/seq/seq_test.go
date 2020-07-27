package seq

import (
	"context"
	"reflect"
	"testing"

	"github.com/listen-lavender/goseq/mock"
	"github.com/agiledragon/gomonkey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSeq_HttpGetSeq(t *testing.T) {
	Convey("TestHttpGetSeq", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			namespace string = "cloudim"
			regionID  uint16 = 1
			incResult uint64 = 2
			err       error  = nil
		)

		ginCtx := &gin.Context{}

		var ctx context.Context = ginCtx

		patches := gomonkey.ApplyMethod(reflect.TypeOf(ginCtx), "Param", func(_ *gin.Context, _ string) string {
			return namespace
		})

		defer patches.Reset()
		patches.ApplyMethod(reflect.TypeOf(ginCtx), "JSON", func(_ *gin.Context, code int, obj interface{}) {
			So(code, ShouldEqual, 200)
			// t.Log(code, obj)
		})

		mockSeqDao := mock.NewMockSeqDao(ctrl)
		gomock.InOrder(
			mockSeqDao.EXPECT().AtomicInc(ctx, namespace).Return(incResult),
		)

		mockSegmentDao := mock.NewMockSoftSegmentDao(ctrl)
		gomock.InOrder(
			mockSegmentDao.EXPECT().CAS(ctx, regionID, incResult).Return(incResult, true, err),
		)

		mockSoftRegionDao := mock.NewMockSoftRegionDao(ctrl)
		gomock.InOrder(
			mockSoftRegionDao.EXPECT().GetSet(ctx, namespace).Return(regionID),
			// mockSoftRegionDao.EXPECT().Hash(ctx, namespace).Return(regionID),
		)

		ss := NewSeqService(mockSeqDao, mockSegmentDao, mockSoftRegionDao)
		ss.HttpGetSeq(ginCtx)
	})
}

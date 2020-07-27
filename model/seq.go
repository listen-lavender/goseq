package model

import (
	"context"
)

//go:generate mockgen -destination=../mock/seq_mock.go -package=mock github.com/listen-lavender/goseq/model SeqDao

// Seq 获取顺序号操作
type SeqDao interface {
	AtomicInc(ctx context.Context, namespace string) uint64
}

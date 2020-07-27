package test

import (
	"context"
	"testing"

	"github.com/listen-lavender/goseq/api/pbseq"
	"google.golang.org/grpc"
)

func BenchmarkGrpcGetSeq(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:8091", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	client := pbseq.NewSeqServiceClient(conn)
	for i := 0; i < b.N; i++ { //use b.N for looping
		req := &pbseq.GetSeqReq{
			Namespace: "haokuan",
		}
		ctx := context.Background()
		res, err := client.GetSeq(ctx, req)
		if false {
			if err != nil {
				println("err: ", err.Error())
			} else {
				println("res: ", res.Namespace, res.Seq, res.Ts, res.Msg, res.ErrorCode)
			}
		}
	}
}

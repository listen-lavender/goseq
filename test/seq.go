package main

import (
	"context"

	"github.com/listen-lavender/goseq/api/pbseq"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8091", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	client := pbseq.NewSeqServiceClient(conn)

	for k := 1; k < 100; k++ {
		req := &pbseq.GetSeqReq{
			Namespace: "haokuan",
		}
		ctx := context.Background()
		res, err := client.GetSeq(ctx, req)
		if err != nil {
			println("err: ", err.Error())
		} else {
			println("res: ", res.Namespace, res.Seq, res.Ts, res.Msg, res.ErrorCode)
		}
	}
}

package grpc

import (
	"net"

	"github.com/listen-lavender/goseq/api/pbseq"
	"github.com/listen-lavender/goseq/service/seq"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer(seqService *seq.SeqService) *GrpcServer {
	server := grpc.NewServer()
	pbseq.RegisterSeqServiceServer(server, seqService)
	return &GrpcServer{
		server: server,
	}
}

func (gs *GrpcServer) Run(host string) {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	gs.server.Serve(listen)
}

package app

import (
	"github.com/gin-gonic/gin"

	"github.com/listen-lavender/goseq/service/seq"
)

type HttpServer struct {
	engine     *gin.Engine
	seqService *seq.SeqService
}

func NewHttpServer(engine *gin.Engine,
	seqService *seq.SeqService) *HttpServer {
	return &HttpServer{
		engine:     engine,
		seqService: seqService,
	}
}

func (hs *HttpServer) Init() {
	hs.engine.GET("/seq/independent/:ns/next", hs.seqService.GetSeq)
}

func (hs *HttpServer) Run(host string) {
	hs.engine.Run(host)
}

package main

import (
	"flag"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	apphttp "github.com/listen-lavender/goseq/app/http"
	"github.com/listen-lavender/goseq/conn/db"
	"github.com/listen-lavender/goseq/conn/memory"
	"github.com/listen-lavender/goseq/dao"
	sSeq "github.com/listen-lavender/goseq/service/seq"
)

type Config struct {
	App        App
	HttpServer HttpServer
	SeqStorage Mongo `toml:"mongo_seq"`
}

type App struct {
	Name  string
	Mode  string
	Debug bool
}

type HttpServer struct {
	Host string
}

type Mongo struct {
	Addr     string
	Authdb   string
	Authuser string
	Authpass string
	Db       string
}

const (
	ConfigFlag        = "config"
	DefaultConfigFile = "config.toml"
)

var (
	IgnoreError flag.ErrorHandling = 99
	config      Config
)

func init() {
	if err := initFromFlag(); err != nil {
		panic(err)
	}
	initRuntime()
}

func initFromFlag() error {
	fs := flag.NewFlagSet("config", IgnoreError)
	fs.SetOutput(ioutil.Discard)
	fs.String(ConfigFlag, DefaultConfigFile, "default config")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	if f := fs.Lookup("v"); f != nil {
		if f.Value.String() == "" {
			return nil
		}
	}
	if f := fs.Lookup("h"); f != nil {
		if f.Value.String() == "" {
			return nil
		}
	}

	config = Config{}
	file := fs.Lookup(ConfigFlag).Value.String()
	_, err := toml.DecodeFile(file, &config)
	return err
}

func initRuntime() {
	ncpu := runtime.NumCPU()
	if ncpu > 2 {
		runtime.GOMAXPROCS(int(ncpu) / 2)
	}
}

func main() {

	//PrintAbout("api start")
	gin.SetMode("release")
	//tesing模式

	softRegionmemory := memory.NewSoftRegion(map[string]uint16{})
	seqmemory := memory.NewSeq(map[string]*uint64{})
	softSegmentmemory := memory.NewSoftSegment(map[uint16]uint64{})

	mongoClient, _ := db.NewMongo(config.SeqStorage.Addr,
		config.SeqStorage.Authdb,
		config.SeqStorage.Db,
		config.SeqStorage.Authuser,
		config.SeqStorage.Authpass)

	hardSegmentHandler := dao.NewHardSegmentHandler(mongoClient)
	hardRegionHandler := dao.NewHardRegionHandler(mongoClient)
	softSegmentHandler := dao.NewSoftSegmentHandler(softSegmentmemory, hardSegmentHandler)
	softRegionHandler := dao.NewSoftRegionHandler(softRegionmemory, hardRegionHandler)
	seqHandler := dao.NewSeqHandler(seqmemory, softRegionHandler, softSegmentHandler)

	seqService := sSeq.NewSeqService(seqHandler, softSegmentHandler, softRegionHandler)
	engine := gin.Default()

	// 性能分析 - 正式环境不要使用！！！
	//pprof.Register(engine)
	httpServer := apphttp.NewHttpServer(engine,
		seqService)

	// 设置路由
	httpServer.Init()
	httpServer.Run(config.HttpServer.Host)
}

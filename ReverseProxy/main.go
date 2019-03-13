package main

import (
	"flag"
	// "/dist/swaggerui"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	realSrvAddr = flag.String("real_srv_addr", "localhost:8000", "endpoint of real server")
	srvAddr     = flag.String("srv_addr", ":8001", "endpoint of reverse proxy server")
)

func run() error {
	var ctx context.Context
	var cc context.CancelFunc
	ctx = context.Background()
	ctx, cc = context.WithCancel(ctx)
	defer cc()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// todo gRpc server register
	// ctx gwMux, *realSrvAddr opts
	_ = opts

	mux.Handle("/", gwMux)
	log.Println("realSrvAddr:", *realSrvAddr, "srvAddr:", *srvAddr)
	return http.ListenAndServe(*srvAddr, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("../proto", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		// from dist/swaggerui
		Asset:    nil, //swaggerui.Asset,
		AssetDir: nil, //swaggerui.AssetDir,
		Prefix:   "path/swaggerui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

package http_rpc

import (
	"../http_rpc/errors"
	"encoding/json"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-os/config"
	"github.com/micro/go-web"
	"github.com/prometheus/common/log"
	"github.com/rs/cors"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//go-os config
var ConfigDefault config.Config

var Debug = false

//类似于路由
const ApiServiceName = "TODO: Write a name what you want .."

//配置Host
func APIHost() string {
	return ConfigDefault.Get("api_host").String("0.0.0.0:8080")
}

func ApiAllowedOrigins() []string {
	return ConfigDefault.Get("api_allowed_origins").StringSlice([]string{"*"})
}

//错误定义
var (
	ErrInvalidPing = errors.BadRequest(10001, "bad ...")
)

func main() {
	//启动一个web服务监听
	address := APIHost()
	service := web.NewService(
		web.Name(ApiServiceName),
		web.Address(address),
	)
	service.Init()

	//跨域处理
	c := cors.New(cors.Options{
		AllowedOrigins: ApiAllowedOrigins(),
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		Debug:          Debug,
		MaxAge:         3600,
	})

	//test
	service.Handle("/ping", c.Handler(http.HandlerFunc(Ping)))

}

// http request to rpc context
func RequestToContext(r *http.Request) context.Context {
	//定义 non-nil empty Context
	ctx := context.Background()
	//通过元数据来访问RPC
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}

func API(w http.ResponseWriter, r *http.Request) {
	//TODO: Get Args   请求参数
	//TODO: r.URL.Path 解析

	req := client.NewJsonRequest("service", "method", "args")

	ctx := RequestToContext(r)
	var rpcrsp json.RawMessage
	err := client.Call(ctx, req, &rpcrsp)
	log.Debug("req: %v, rsp: %v, err: %v", req, string(rpcrsp), err)
	err = errors.ParseFromRPCError(err)
	Response(w, rpcrsp, err)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Debug("ping read error: %v", err)
		Response(w, nil, ErrInvalidPing)
		return
	}

	var t float64
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Info("unmarshal error: %v", err)
		Response(w, nil, ErrInvalidPing)
	}

	now := float64(time.Now().UnixNano()) / 1e9
	b, err := json.Marshal([]float64{t, now})
	Response(w, b, err)
}

func Response(w http.ResponseWriter, body json.RawMessage, err error) {
	if err != nil {
		ce := errors.Parse(err.Error())
		if !Debug {
			ce.Internal = ""
		}
		switch ce.Status {
		case 0:
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Status))
		}
		w.Write([]byte(ce.Error()))
		return
	}
	b, _ := body.MarshalJSON()
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}

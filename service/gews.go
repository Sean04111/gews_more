package main

import (
	"flag"
	"fmt"

	"gews_more/service/internal/config"
	"gews_more/service/internal/handler"
	"gews_more/service/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/gews-api.yaml", "the config file")
	//Ca, _ = cache.NewCache("memory", `{"interval":60}`)
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

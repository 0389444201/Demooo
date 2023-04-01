package main

import (
	"flag"
	"fmt"

	"demo/userAccount/api/internal/config"
	"demo/userAccount/api/internal/handler"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/rediscache"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	rediscache.NewRedisCache()
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

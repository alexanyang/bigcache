package main

import (
	"log"

	"github.com/houzhongjian/bigcache/app/cache-server/handler"
	"github.com/houzhongjian/bigcache/cmd"
	"github.com/houzhongjian/bigcache/lib/conf"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	cmd := cmd.New()
	conf.Load(cmd.Conf)

	srv := handler.NewServer()
	srv.Start()
}

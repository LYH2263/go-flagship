// Command flagd 启动特性开关管理服务，默认监听 :8095。
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LYH2263/go-flagship"
	"github.com/LYH2263/go-flagship/internal/api"
)

func main() {
	addr := flag.String("addr", ":8095", "HTTP 监听地址")
	web := flag.String("web", "web", "静态管理页目录")
	persist := flag.String("persist", "", "flag 快照 JSON 路径（可选）")
	flag.Parse()

	opts := []flagship.Option{}
	if *persist != "" {
		opts = append(opts, flagship.WithPersistPath(*persist))
	}

	ship := flagship.New(opts...)
	defer ship.Close()

	srv := api.New(ship, api.Options{WebDir: *web, AllowCORS: true})
	httpSrv := &http.Server{
		Addr:              *addr,
		Handler:           srv,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("flagd 监听 %s", *addr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	_ = httpSrv.Close()
}

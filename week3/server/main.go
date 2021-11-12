package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//起一个http server服务
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "我是server服务")
	})
	eg, ctx := errgroup.WithContext(context.Background())

	server := http.Server{
		Handler: handler,
		Addr:    ":6335",
	}
	eg.Go(func() error {
		return server.ListenAndServe()
	})
	shutdownch := make(chan struct{})
	//监听退出信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	//go.Go 1
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			shutdownch <- struct{}{}
			return fmt.Errorf("ctx Done, in go.Go1")
		case <-ch:
			shutdownch <- struct{}{}
			return fmt.Errorf("sig quit, in go.Go1")
		}
	})

	//go.Go 2
	eg.Go(func() error {
		select {
		//case <-ctx.Done():
		//	log.Println("errgroup context done")
		case <-shutdownch:
			log.Println("shut chan ....ing")
		}
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		log.Println("server shutting down")
		return server.Shutdown(timeoutCtx)
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("eg.error : %+v", err)
	}
}

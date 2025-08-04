package main

import (
	"context"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/oklog/run"

	"github.com/crazyfrankie/ddd-todolist/backend/cmd"
	"github.com/crazyfrankie/ddd-todolist/backend/conf"
)

func main() {
	g := &run.Group{}

	srv, err := cmd.Init()
	if err != nil {
		panic("InitializeInfra failed, err=" + err.Error())
	}

	httpSrv := &http.Server{
		Handler: srv,
		Addr:    conf.GetConf().Server.Addr,
	}

	g.Add(func() error {
		log.Printf("Server is running at http://localhost%s\n", conf.GetConf().Server)
		return httpSrv.ListenAndServe()
	}, func(err error) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			log.Printf("failed to shutdown main server: %v", err)
		}
	})

	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		log.Printf("program interrupted, err:%s", err)
	}
}

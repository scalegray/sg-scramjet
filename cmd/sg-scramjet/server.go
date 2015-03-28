package main

import (
	log "code.google.com/p/log4go"
	"github.com/scalegray/sg-scramjet/cmd/sg-scramjet/server"
	"github.com/tsuru/config"
	"runtime"
	"time"
	//"fmt"
)

func serverRun(dry bool) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	version, _ := config.GetString("version")

	log.Info("Starting scalegray engine %s...", version)

	server, err := server.NewServer()
	log.Info("Server started---> %v.", server)
	if err != nil {
		// sleep for the log to flush
		time.Sleep(time.Second)
		panic(err)
	}

	if err := startProfiler(server); err != nil {
		panic(err)
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Error("ListenAndServe failed: ", err)
	}
}

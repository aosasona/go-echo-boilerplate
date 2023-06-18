package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"gopi/core"
	"gopi/internal/config"
	"gopi/internal/conn"
)

func main() {
	c, err := core.New()
	if err != nil {
		panic(err)
	}

	db, err := conn.InitDB()
	if err != nil {
		panic(err)
	}

	cache, err := conn.InitBolt()
	if err != nil {
		panic(err)
	}

	conf, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	c.SetCache(cache).
		SetDB(db).
		SetConfig(conf).
		InitApp()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		err = c.CloseConnections()
		if err != nil {
			log.Fatalf("failed to close connections: %v", err)
		}
		close(done)
	}()
	go handleShutDown(done, c)

	if err = c.Run(); err != nil {
		panic(err)
	}
}

func handleShutDown(done chan os.Signal, core *core.Core) {
	_ = <-done
	log.Info("Shutting down...")
	_ = core.Kill()
}

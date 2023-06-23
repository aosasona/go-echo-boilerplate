package main

import (
	"gopi/core"
	"gopi/internal/config"
	"os"
	"os/signal"
	"syscall"

	conn2 "gopi/internal/conn"

	"github.com/charmbracelet/log"
)

func init() {
}

func main() {
	c, err := core.New()
	mustNot(err)

	cache, err := conn2.InitBolt()
	mustNot(err)
	conf, err := config.LoadEnv(".")
	mustNot(err)

	db, err := conn2.InitDB(conf.IsDev)
	mustNot(err)

	mustNot(err)

	c.SetCache(cache).
		SetDB(db).
		SetConfig(conf).
		InitApp()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cleanup(c, done)
	go handleShutDown(done, c)

	mustNot(c.Run())
}

func cleanup(c *core.Core, done chan os.Signal) {
	err := c.CloseConnections()
	if err != nil {
		log.Fatalf("failed to close connections: %v", err)
	}
	close(done)
}

func handleShutDown(done chan os.Signal, core *core.Core) {
	_ = <-done
	log.Info("Shutting down...")
	_ = core.Kill()
}

func mustNot(err error) {
	if err != nil {
		panic(err)
	}
}

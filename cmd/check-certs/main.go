package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Format-C-eft/check-certs/internal/bootstrap"
	"github.com/Format-C-eft/check-certs/internal/config"
)

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancelFn()

	config.ParseServiceFlags()
	fatalIfErr(config.InitConfig(), "init config")

	if config.VersionCommand {
		fmt.Println(config.GetVersion())
		return
	}

	log.Println("Application start ...")

	store, err := bootstrap.Initialize(ctx)
	fatalIfErr(err, "bootstrap.Initialize")

	chanDone := make(chan struct{})

	go store.CheckAndReissue(ctx, chanDone)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v\n", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v\n", done)
	case <-chanDone:
		log.Println("Processing is finished")
	}

	log.Println("Application stop ...")
}

func fatalIfErr(err error, what string) {
	if err != nil {
		log.Fatalf("%s: %s", what, err)
	}
}

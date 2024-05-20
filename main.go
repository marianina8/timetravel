package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/marianina8/timetravel/cmd"
	"github.com/marianina8/timetravel/utils"
)

var Version = "dev"

func main() {

	// code to handle signal interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c
		utils.Print("Time circuits off. See you in another timeline!")
		os.Exit(1)
	}()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

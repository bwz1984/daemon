package main

import (
	"context"
	"log"
	"time"

	"github.com/bwz1984/daemon/daemon"
)

func main() {
	ctx := context.TODO()
	d := daemon.DaemonNew(ctx, func() {
		log.Println("enter process func")
		for {
			time.Sleep(time.Second)
		}
	}, time.Second*5)
	d.Run()
}

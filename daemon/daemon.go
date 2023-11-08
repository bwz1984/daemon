package daemon

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Daemon struct {
	interval  time.Duration // 执行间隔
	processFn func()
}

func DaemonNew(ctx context.Context, processFn func(), interval time.Duration) *Daemon {
	return &Daemon{processFn: processFn, interval: interval}
}

func (daemon *Daemon) Run() {
	go func() {
		daemon.runProcessFn()
	}()
	// 退出信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	<-signals
	log.Println("wait for 15s before exiting.")
	time.Sleep(time.Second * 15)
	log.Println("exit now.")
}

func (daemon *Daemon) runProcessFn() {
	if daemon.processFn == nil {
		return
	}
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("panic %v", err)
				}
			}()
			daemon.processFn()
			if daemon.interval > 0 {
				time.NewTimer(daemon.interval)
				<-time.After(daemon.interval)
			}
		}()
	}
}

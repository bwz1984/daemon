package daemon

import (
	"context"
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
	return &Daemon{processFn: processFn}
}

func (daemon *Daemon) Run() {

	daemon.listenExitSignal()
}

func (daemon *Daemon) listenExitSignal() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	for {
		select {
		case <-signals:
			return
		}
	}
}

func (daemon *Daemon) runProcessFn() {

	for {
		daemon.processFn()
	}
}

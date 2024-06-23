package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/wait/pkg/wait"
)

var listenedSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGHUP,
}

func main() {
	fs := flag.NewFlagSet("wait", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	loggerConfig := logger.Flags(fs, "logger")

	addresses := flags.New("Address", "Dial address in the form network:host:port, e.g. tcp:localhost:5432").DocPrefix("wait").StringSlice(fs, nil, nil)
	timeout := flags.New("Timeout", "Timeout of retries").DocPrefix("wait").Duration(fs, time.Second*10, nil)
	next := flags.New("Next", "Action to execute after").DocPrefix("wait").String(fs, "", nil)
	args := flags.New("NextArg", "Args for the action to execute").DocPrefix("wait").StringSlice(fs, nil, nil)

	_ = fs.Parse(os.Args[1:])

	ctx := context.Background()

	logger.Init(ctx, loggerConfig)

	var wg sync.WaitGroup
	var success atomic.Uint32

	for _, address := range *addresses {
		parts := strings.Split(strings.TrimSpace(address), ":")
		if len(parts) != 3 {
			slog.Error("address has invalid format")
			os.Exit(1)
		}

		if len(parts[0]) == 0 {
			slog.Error("network is required")
			os.Exit(1)
		}

		if len(parts[1]) == 0 {
			slog.Error("host is required")
			os.Exit(1)
		}

		if len(parts[2]) == 0 {
			slog.Error("port is required")
			os.Exit(1)
		}

		wg.Add(1)
		go func(network, address string) {
			defer wg.Done()

			if wait.Wait(network, address, *timeout) {
				success.Add(1)
			}
		}(parts[0], fmt.Sprintf("%s:%s", parts[1], parts[2]))
	}

	wg.Wait()

	if len(*addresses) != int(success.Load()) {
		os.Exit(1)
	}

	action := strings.TrimSpace(*next)
	if len(action) == 0 {
		return
	}

	command := exec.Command(action, *args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	go func() {
		signalsChan := make(chan os.Signal, 1)
		defer close(signalsChan)

		signal.Notify(signalsChan, listenedSignals...)
		defer signal.Stop(signalsChan)

		for signal := range signalsChan {
			if err := command.Process.Signal(signal); err != nil {
				slog.LogAttrs(ctx, slog.LevelError, "sending signal", slog.String("signal", signal.String()), slog.Any("error", err))
			}
		}
	}()

	if err := command.Run(); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "command", slog.Any("error", err))
		os.Exit(1)
	}
}

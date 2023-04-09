package main

import (
	"errors"
	"flag"
	"fmt"
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
	syscall.SIGCONT,
	syscall.SIGHUP,
}

func main() {
	fs := flag.NewFlagSet("wait", flag.ExitOnError)

	loggerConfig := logger.Flags(fs, "logger")

	addresses := flags.StringSlice(fs, "", "wait", "Address", "Dial address in the form network:host:port, e.g. tcp:localhost:5432", nil, nil)
	timeout := flags.Duration(fs, "", "wait", "Timeout", "Timeout of retries", time.Second*10, nil)
	next := flags.String(fs, "", "wait", "Next", "Action to execute after", "", nil)
	args := flags.StringSlice(fs, "", "wait", "NextArg", "Args for the action to execute", nil, nil)

	logger.Fatal(fs.Parse(os.Args[1:]))

	logger.Global(logger.New(loggerConfig))

	var wg sync.WaitGroup
	var success atomic.Uint32

	for _, address := range *addresses {
		parts := strings.Split(strings.TrimSpace(address), ":")
		if len(parts) != 3 {
			logger.Fatal(errors.New("address has invalid format"))
		}

		if len(parts[0]) == 0 {
			logger.Fatal(errors.New("network is required"))
		}

		if len(parts[1]) == 0 {
			logger.Fatal(errors.New("host is required"))
		}

		if len(parts[2]) == 0 {
			logger.Fatal(errors.New("port is required"))
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
				logger.Error("sending `%s` signal: %s", signal, err)
			}
		}
	}()

	logger.Fatal(command.Run())
}

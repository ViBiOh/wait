package wait

import (
	"context"
	"log/slog"
	"net"
	"time"
)

func Wait(network, addr string, timeout time.Duration) bool {
	if timeout == 0 {
		return dial(network, addr)
	}

	timeoutTimer := time.NewTimer(timeout)
	defer func() {
		timeoutTimer.Stop()

		select {
		case <-timeoutTimer.C:
		default:
		}
	}()

	for {
		select {
		case <-timeoutTimer.C:
			return false

		default:
			if dial(network, addr) {
				return true
			}

			time.Sleep(time.Second)
		}
	}
}

func dial(network, addr string) bool {
	conn, err := net.DialTimeout(network, addr, time.Second)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "dial", slog.String("addr", addr), slog.String("network", network), slog.Any("error", err))
		return false
	}

	if closeErr := conn.Close(); closeErr != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "close", slog.Any("error", err))
	}

	return true
}

package wait

import (
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
		slog.Warn("dial", "error", err, "addr", addr, "network", network)
		return false
	}

	if closeErr := conn.Close(); closeErr != nil {
		slog.Warn("close", "error", err)
	}

	return true
}

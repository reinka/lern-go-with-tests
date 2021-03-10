package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = time.Second * 10

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return a, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}

func SequentialRacer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}

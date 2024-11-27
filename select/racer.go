package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// struct{} is an empty struct, which occupies zero bytes of memory
	// Using chan struct{} is a common pattern when you only need to signal an event (completion or readiness) without passing any data
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

package racer

import "net/http"

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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

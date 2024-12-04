package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()
		// use select to effectively race to the two asynchronous processes and then we either write a response or Cancel.
		select {
		case d := <-data:
			fmt.Fprint(w, d)
			//  context has a method Done(), which returns a channel which gets sent a signal when the context is "done" or "cancelled".
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

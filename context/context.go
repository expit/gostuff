package context

import (
	"context"
	"fmt"
	"net/http"
)

//Store is interface that implements Fetch function
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

//Server starts Server
func Server(store Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprintf(w, data)
	}
}

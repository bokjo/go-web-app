package middleware

import (
	"context"
	"net/http"
	"time"
)

// TimeoutMiddleware struct
type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()

	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	//defer cancel()

	r.WithContext(ctx)

	success := make(chan struct{})

	go func() {
		tm.Next.ServeHTTP(w, r)
		success <- struct{}{}
	}()

	select {
	case <-success:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}

}

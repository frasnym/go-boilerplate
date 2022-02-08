package logging

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

type defaultLogging struct{}

func NewDefaultLogging() Logging {
	return &defaultLogging{}
}

func (*defaultLogging) LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err, debug.Stack())
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		fmt.Println(wrapped.status, r.Method, r.URL.EscapedPath(), time.Since(start))
	}

	return http.HandlerFunc(fn)
}

package logging

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

type zapLogging struct{}

func NewZapLogging() Logging {
	fmt.Println("Zap logging initilized")
	return &zapLogging{}
}

func (*zapLogging) LoggingMiddleware(next http.Handler) http.Handler {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				sugar.Errorw("Log err",
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		sugar.Infow(http.StatusText(wrapped.status),
			"status", wrapped.status,
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
	}

	return http.HandlerFunc(fn)
}

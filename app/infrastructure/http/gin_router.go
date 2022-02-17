package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NextRequestHandler struct {
	c *gin.Context
}

// Run the next request in the middleware chain and return
// See: https://godoc.org/github.com/gin-gonic/gin#Context.Next
func (h *NextRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.c.Next()
}

type ginRouter struct{}

var ginDispatcher = gin.New()

func NewGinRouter(loggingMiddleware func(next http.Handler) http.Handler) Router {
	WrapHH := func(hh func(h http.Handler) http.Handler) gin.HandlerFunc {
		// Steps:
		// - create an http handler to pass `hh`
		// - call `hh` with the http handler, which returns a function
		// - call the ServeHTTP method of the resulting function to run the rest of the middleware chain

		return func(c *gin.Context) {
			hh(&NextRequestHandler{c}).ServeHTTP(c.Writer, c.Request)
		}
	}

	ginDispatcher.Use(WrapHH(loggingMiddleware))
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.GET(uri, gin.WrapF(f))
}

func (*ginRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.POST(uri, gin.WrapF(f))
}

func (*ginRouter) SERVE(port string) {
	fmt.Printf("Gin HTTP service is running on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), ginDispatcher)
}

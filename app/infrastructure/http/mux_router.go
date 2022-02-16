package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter(loggingMiddleware func(next http.Handler) http.Handler) Router {
	muxDispatcher.Use(loggingMiddleware)
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP service is running on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), muxDispatcher)
}

package main

import (
	"http2/http/cmd/app"
	"http2/http/pkg/banners"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", hi)
	mux.HandleFunc("/", notFound)

	//myHandler := handler{}
	httpServer := http.Server{
		Addr:    "0.0.0.0:9999",
		Handler: mux,
	}
	log.Println("server start")
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("http server: ", err)
	}

}
func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}
	return srv.ListenAndServe()
}

//type handler struct{}
//
//func (handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	//log.Println("handler ServeHTTP")
//	//writer.Write([]byte("hello, ====> handler ServeHTTP"))
//	path := request.URL.Path
//	switch path {
//	case "/hi":
//		writer.WriteHeader(http.StatusOK)
//		writer.Write([]byte("hello, ====> handler ServeHTTP"))
//	default:
//		writer.WriteHeader(http.StatusNotFound)
//		writer.Write([]byte("Page not found"))
//	}
//}
func hi(writer http.ResponseWriter, request *http.Request) {
	log.Println("hi====> test ", request.URL.Path)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("hello, ====> handler ServeHTTP"))
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	log.Println("hi====> test ", request.URL.Path)
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Page not found ++++=====>"))
}

//package main
//
//import (
//	"log"
//	"net"
//	"net/http"
//	"os"
//	"sync"
//)
//
//type handler struct {
//	mu       *sync.RWMutex
//	handlers map[string]http.HandlerFunc
//}
//
//func (h *handler) ServerHttp(writer http.ResponseWriter, request *http.Request) {
//	_, err := writer.Write([]byte ("hello world"))
//	if err != nil {
//		log.Print(err)
//	}
//}
//func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	h.mu.RLock()
//	handler, ok := h.handlers[request.URL.Path]
//	h.mu.RUnlock()
//	if !ok {
//		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//	handler(writer,request)
//}
//
//func main() {
//	host := "0.0.0.0"
//	port := "9999"
//	if err := execute(host, port); err != nil {
//		os.Exit(1)
//	}
//}
//
//func execute(host string, port string) (err error) {
//	srv := &http.Server{
//		Addr:    net.JoinHostPort(host, port),
//		Handler: &handler{},
//	}
//	return srv.ListenAndServe()
//}
//
////func (srv *Server) ListenAndServe() error {
////if srv.shu
////}

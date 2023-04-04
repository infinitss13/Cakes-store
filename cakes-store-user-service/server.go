package cakes_store_user_service

import "net/http"

type Server struct {
	httpServer *http.Server
}

//Run : method that start server, that will work until error returns
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

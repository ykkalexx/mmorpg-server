package main

// add database once ready
type Server struct {
	listenAddr string
}

func newServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}


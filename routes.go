package main

func (s *server) routes() {
	s.router.HandleFunc("/index", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/oauth/redirect", s.handleRedirect()).Methods("GET")
}

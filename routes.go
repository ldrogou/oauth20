package main

func (s *server) routes() {
	s.router.HandleFunc("/index", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/oauth/redirect", s.handleRedirect()).Methods("GET")
	s.router.HandleFunc("/local", s.handleLocal()).Methods("POST")
	s.router.HandleFunc("/oauth20", s.handleOAuth20()).Methods("POST")
	s.router.HandleFunc("/jwt/{id}", s.handleJSONWebToken()).Methods("GET")
	s.router.HandleFunc("/jwt/refresh/{id}", s.handleRefreshToken()).Methods("POST")

}

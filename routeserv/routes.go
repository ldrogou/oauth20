package routeserv

func (s *Server) routes() {
	s.Router.HandleFunc("/index", s.handleIndex()).Methods("GET")
	s.Router.HandleFunc("/oauth/redirect", s.handleRedirect()).Methods("GET")
	s.Router.HandleFunc("/local", s.handleLocal()).Methods("POST")
	s.Router.HandleFunc("/oauth20", s.handleOAuth20()).Methods("POST")
	s.Router.HandleFunc("/jwt/{id}", s.handleJSONWebToken()).Methods("GET")
	s.Router.HandleFunc("/jwt/refresh/{id}", s.handleRefreshToken()).Methods("POST")

}

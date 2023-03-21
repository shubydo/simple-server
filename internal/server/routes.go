package server

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	//s.router.HandleFunc("/hello", s.handleHello)

}

package server
package server

func (s *server) routes() {
    s.router.HandleFunc("/", s.indexHandler)
    s.router.HandleFunc("/api/v1", s.apiHandler)
    s.router.HandleFunc("/api/v2", s.apiV2Handler)
    s.router.HandleFunc("/api/v3", s.apiV3Handler)
    s.router.HandleFunc("/api/v4", s.apiV4Handler)
    s.router.HandleFunc("/api/v5", s.apiV5Handler)
    s.router.HandleFunc("/api/v6", s.apiV6Handler)
    s.router.HandleFunc("/api/v7", s.apiV7Handler)
    s.router.HandleFunc("/api/v8", s.apiV8Handler)
}
package api

func (s *Server) Healthy() bool {
	return s != nil && s.ship != nil && !s.ship.Closed()
}

package HttpServer

func (s *Server) setupRoutes() {
	s.echo.GET("/v1/validate", s.validateHandler.Validate)
	s.echo.GET("/v2/validate", s.validateHandler.ValidateV2)
	s.echo.GET("/v3/validate", s.validateHandler.ValidateV3)
}

package server

func registerRoutes(r *Router, h *Handler, mm *MiddlewareManager) {
	r.GET("/", h.handleGetIP, mm.UABlock, mm.RateLimit)
	r.GET("/health", h.handleHealthCheck)
}

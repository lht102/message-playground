package rest

import "github.com/uptrace/bunrouter"

func (s *Server) routes() {
	r := s.router

	r.WithGroup("/api", func(g *bunrouter.CompatGroup) {
		g.WithGroup("/job", func(g *bunrouter.CompatGroup) {
			g.GET("/:uuid", getJobHandler(s.jobService))
			g.POST("", createJobHandler(s.jobService))
		})
	})
}

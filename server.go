package Advertising

import (
	"Advertising/configs"
	"Advertising/pkg/logging"
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *configs.Config, handler http.Handler) error {
	logger := logging.GetLogger()
	logger.Info("start application")
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Listen.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const defaultReadTimeout = time.Second

type Server struct {
	http.Server
}

func NewHTTPServer(router Router, port int) Server {
	return Server{
		Server: http.Server{
			Addr:              fmt.Sprintf(":%v", port),
			Handler:           router,
			ReadHeaderTimeout: defaultReadTimeout,
		},
	}
}

func (s *Server) Run(ctx context.Context) {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Infof("http srv: run on addr %v", s.Server.Addr)

		if err := s.Server.ListenAndServe(); err != nil {
			return fmt.Errorf("failed to run server: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		if err := s.Server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Info("http srv: successfully stopped")
	}
}

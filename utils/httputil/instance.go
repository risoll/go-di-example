package httputil

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	Instance interface {
		Start() error
		ShutDown() error
	}

	instanceImpl struct {
		httpServer *http.Server
	}
)

func NewInstance(server *http.Server) Instance {
	return &instanceImpl{
		httpServer: server,
	}
}

func (s *instanceImpl) Start() error {
	if s.httpServer == nil {
		s.httpServer = &http.Server{}
	}

	logrus.Println("service running in", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe() // Blocking
	if err == http.ErrServerClosed {
		logrus.WithError(err).Error("http server is closed")
		return s.ShutDown()
	}

	return nil
}

func (s *instanceImpl) ShutDown() error {
	if s.httpServer == nil {
		logrus.Error("http server is nil")
		return errors.New("http server is nil")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to shutdown http server")
		return fmt.Errorf("failed to shutdown http server -> %s", err.Error())
	}

	return nil
}

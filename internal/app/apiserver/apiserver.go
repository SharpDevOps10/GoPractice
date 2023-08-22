package apiserver

import "github.com/sirupsen/logrus"

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting server")

	return nil
}

func (s *APIServer) configureLogger() error {
	levelOfLogging, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(levelOfLogging)

	return nil
}

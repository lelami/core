package service

import (
	"core/pkg/repository"
)

type LogService struct {
	log repository.Logging
}

func NewLogService(log repository.Logging) *LogService {
	return &LogService{log: log}
}

func (s *LogService) WriteLog(log string) error {
	return s.log.WriteLog(log)
}

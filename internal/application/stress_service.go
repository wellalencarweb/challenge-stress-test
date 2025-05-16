package application

import (
	"time"
	"challenge-stress-test/internal/domain"
)

type StressService struct {
	runner   domain.Runner
	reporter domain.Reporter
}

func NewStressService(r domain.Runner, rep domain.Reporter) *StressService {
	return &StressService{runner: r, reporter: rep}
}

func (s *StressService) Execute(url string, totalRequests, concurrency int) {
	start := time.Now()
	result := s.runner.Run(url, totalRequests, concurrency)
	totalTime := time.Since(start).Seconds()
	s.reporter.Report(totalTime, result)
}

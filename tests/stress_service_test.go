package tests

import (
	"testing"
	"challenge-stress-test/internal/application"
	"challenge-stress-test/internal/domain"
)

type MockRunner struct{}

func (m *MockRunner) Run(url string, totalRequests, concurrency int) map[int]int {
	return map[int]int{200: totalRequests}
}

type MockReporter struct {
	reportCalled bool
}

func (m *MockReporter) Report(totalTime float64, result map[int]int) {
	m.reportCalled = true
}

func TestStressService_Execute(t *testing.T) {
	runner := &MockRunner{}
	reporter := &MockReporter{}
	service := application.NewStressService(runner, reporter)
	service.Execute("http://example.com", 10, 2)
	if !reporter.reportCalled {
		t.Error("Relatório não foi chamado")
	}
}

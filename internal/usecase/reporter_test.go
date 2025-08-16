package usecase

import (
	"testing"
	"time"

	"github.com/agnaldopidev/teste_carga_cli/internal/domain"
)

type MockReporter struct {
	Output string
}

func (m *MockReporter) Print(report domain.Report) {
	m.Output = report.URL
}

func TestReporter_Output(t *testing.T) {
	report := domain.Report{URL: "http://teste.com", Duration: 1 * time.Second, TotalRequests: 3}
	mock := &MockReporter{}
	mock.Print(report)
	if mock.Output != "http://teste.com" {
		t.Errorf("Relatório incorreto")
	}
}

package usecase

import (
	"errors"
	"testing"
	"time"

	_ "github.com/agnaldopidev/teste_carga_cli/internal/domain"
)

type MockHTTPClient struct {
	Responses []struct {
		status int
		err    error
	}
	callIndex int
}

func (m *MockHTTPClient) Get(url string) (int, error) {
	if m.callIndex >= len(m.Responses) {
		return 0, errors.New("chamadas extras não previstas")
	}
	resp := m.Responses[m.callIndex]
	m.callIndex++
	return resp.status, resp.err
}

func TestLoadTester_All200(t *testing.T) {
	mock := &MockHTTPClient{
		Responses: []struct {
			status int
			err    error
		}{
			{200, nil}, {200, nil}, {200, nil},
		},
	}
	lt := NewLoadTester(mock)
	report := lt.Run("http://teste.com", 3, 1)
	if report.StatusCount[200] != 3 {
		t.Errorf("Esperava 3 respostas 200, obteve %d", report.StatusCount[200])
	}
}

func TestLoadTester_MixedResponses(t *testing.T) {
	mock := &MockHTTPClient{
		Responses: []struct {
			status int
			err    error
		}{
			{200, nil}, {500, nil}, {0, errors.New("erro")},
		},
	}
	lt := NewLoadTester(mock)
	report := lt.Run("http://teste.com", 3, 2)
	if report.StatusCount[200] != 1 || report.StatusCount[500] != 1 || report.StatusCount[0] != 1 {
		t.Errorf("Distribuição incorreta dos status")
	}
}

func TestLoadTester_ConcurrentExecution(t *testing.T) {
	mock := &MockHTTPClient{
		Responses: []struct {
			status int
			err    error
		}{
			{200, nil}, {200, nil}, {200, nil}, {500, nil},
		},
	}
	lt := NewLoadTester(mock)
	start := time.Now()
	report := lt.Run("http://teste.com", 4, 2)
	if report.TotalRequests != 4 {
		t.Errorf("Total requests incorreto")
	}
	if time.Since(start) > time.Second {
		t.Errorf("Execução demorou muito para concorrência")
	}
}

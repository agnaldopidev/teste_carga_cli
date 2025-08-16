package infra

import (
	"fmt"

	"github.com/agnaldopidev/teste_carga_cli/internal/domain"
)

type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) Print(report domain.Report) {
	fmt.Println("\n=== Relatório do Teste ===")
	fmt.Printf("URL: %s\n", report.URL)
	fmt.Printf("Tempo total: %v\n", report.Duration)
	fmt.Printf("Total de requests: %d\n", report.TotalRequests)
	for code, count := range report.StatusCount {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}

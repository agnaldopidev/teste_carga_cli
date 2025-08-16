package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/agnaldopidev/teste_carga_cli/internal/infra"
	"github.com/agnaldopidev/teste_carga_cli/internal/usecase"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("É necessário informar a URL do serviço")
		os.Exit(1)
	}

	client := infra.NewHTTPClientImpl()
	reporter := infra.NewReporter()
	loadTester := usecase.NewLoadTester(client)

	report := loadTester.Run(*url, *requests, *concurrency)
	reporter.Print(report)
}

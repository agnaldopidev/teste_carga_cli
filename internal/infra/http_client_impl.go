package infra

import (
	"net/http"
	"time"

	"github.com/agnaldopidev/teste_carga_cli/internal/domain"
)

type HTTPClientImpl struct{}

func NewHTTPClientImpl() domain.HTTPClient {
	return &HTTPClientImpl{}
}

func (c *HTTPClientImpl) Get(url string) (int, error) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

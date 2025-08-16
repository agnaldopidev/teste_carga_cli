package usecase

import (
	"sync"
	"time"

	"github.com/agnaldopidev/loadtester/internal/domain"
)

type LoadTester struct {
	client domain.HTTPClient
}

func NewLoadTester(client domain.HTTPClient) *LoadTester {
	return &LoadTester{client: client}
}

func (lt *LoadTester) Run(url string, totalRequests, concurrency int) domain.Report {
	start := time.Now()
	statusCount := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			status, _ := lt.client.Get(url)
			mu.Lock()
			statusCount[status]++
			mu.Unlock()
			<-sem
		}()
	}

	wg.Wait()
	return domain.Report{
		URL:           url,
		Duration:      time.Since(start),
		TotalRequests: totalRequests,
		StatusCount:   statusCount,
	}
}

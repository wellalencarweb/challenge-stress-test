package runner

import (
	"sync"
	"challenge-stress-test/internal/domain"
)

type ConcurrentRunner struct {
	client domain.HTTPClient
}

func NewConcurrentRunner(c domain.HTTPClient) *ConcurrentRunner {
	return &ConcurrentRunner{client: c}
}

func (r *ConcurrentRunner) Run(url string, totalRequests, concurrency int) map[int]int {
	results := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	requestsPerWorker := totalRequests / concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				status, err := r.client.DoRequest(url)
				if err == nil {
					mu.Lock()
					results[status]++
					mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()
	return results
}

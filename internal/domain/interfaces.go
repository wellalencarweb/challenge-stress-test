package domain

type HTTPClient interface {
	DoRequest(url string) (int, error)
}

type Runner interface {
	Run(url string, totalRequests, concurrency int) map[int]int
}

type Reporter interface {
	Report(totalTime float64, result map[int]int)
}

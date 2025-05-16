package main

import (
	"flag"
	"log"

	"challenge-stress-test/internal/application"
	"challenge-stress-test/internal/infra/httpclient"
	"challenge-stress-test/internal/infra/reporter"
	"challenge-stress-test/internal/infra/runner"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *totalRequests <= 0 || *concurrency <= 0 {
		log.Fatal("Parâmetros inválidos. Use --url, --requests e --concurrency")
	}

	httpClient := httpclient.New()
	runner := runner.NewConcurrentRunner(httpClient)
	report := reporter.NewConsoleReporter()

	stressService := application.NewStressService(runner, report)
	stressService.Execute(*url, *totalRequests, *concurrency)
}

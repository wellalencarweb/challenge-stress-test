package reporter

import "fmt"

type ConsoleReporter struct{}

func NewConsoleReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

func (r *ConsoleReporter) Report(totalTime float64, result map[int]int) {
	total := 0
	ok := 0
	fmt.Println("\n========= RELATÓRIO =========")
	fmt.Printf("Tempo total: %.2fs\n", totalTime)
	for code, count := range result {
		total += count
		if code == 200 {
			ok = count
		}
		fmt.Printf("Status %d: %d\n", code, count)
	}
	fmt.Printf("Total de requests: %d\n", total)
	fmt.Printf("Com sucesso (200): %d\n", ok)
	fmt.Println("=============================")
}

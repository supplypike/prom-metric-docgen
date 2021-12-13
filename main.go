package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/supplypike/prom-metric-docgen/metrics"
)

func filterMetric(m metrics.Metric) bool {
	return !strings.HasPrefix(m.Name, "go_")
}

func main() {
	m, err := metrics.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(m)
	filtered := m.Filter(filterMetric)

	fmt.Println("| NAME | TYPE | HELP |")
	fmt.Println("| --- | --- | --- |")
	for _, m := range filtered {
		fmt.Printf("| `%s` | `%s` | %s |\n", m.Name, m.Type, m.Help)
	}
}

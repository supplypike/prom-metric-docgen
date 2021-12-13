package metrics

import (
	dto "github.com/prometheus/client_model/go"
)

type Metric struct {
	Name string
	Type string
	Help string
}

type Metrics []Metric

func MetricFromDto(m dto.MetricFamily) Metric {
	return Metric{
		Name: m.GetName(),
		Type: m.GetType().String(),
		Help: m.GetHelp(),
	}
}

func (m Metrics) Len() int {
	return len(m)
}

func (m Metrics) Less(i, j int) bool {
	a := m[i].Name
	b := m[j].Name
	return a < b
}

func (m Metrics) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Metrics) Filter(predicate func(m Metric) bool) Metrics {
	output := Metrics{}
	for _, metric := range m {
		if predicate(metric) {
			output = append(output, metric)
		}
	}
	return output
}

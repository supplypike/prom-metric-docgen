package metrics_test

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supplypike/prom-metric-docgen/metrics"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("testdata/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expected := metrics.Metrics{
		{
			Name: "mf1",
			Type: "UNTYPED",
		},
		{
			Name: "mf2",
			Type: "COUNTER",
			Help: "hello world",
		},
	}

	data, err := metrics.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	sort.Sort(data)

	assert.EqualValues(t, expected, data)

}

func TestFilter(t *testing.T) {
	mock := metrics.Metrics{
		{
			Name: "mf1",
			Type: "UNTYPED",
		},
		{
			Name: "mf2",
			Type: "COUNTER",
			Help: "hello world",
		},
	}
	expected := metrics.Metrics{
		{
			Name: "mf2",
			Type: "COUNTER",
			Help: "hello world",
		},
	}

	data := mock.Filter(func(m metrics.Metric) bool {
		return m.Type == "COUNTER"
	})

	assert.EqualValues(t, expected, data)

}

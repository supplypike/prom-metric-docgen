package metrics

import (
	"errors"
	"io"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func Decode(input io.Reader) (Metrics, error) {
	decoder := expfmt.NewDecoder(input, expfmt.FmtText)
	buffer := Metrics{}

	for {
		var m dto.MetricFamily
		if err := decoder.Decode(&m); err != nil {

			if errors.Is(err, io.EOF) {
				return buffer, nil
			}

			return buffer, err
		}

		metric := MetricFromDto(m)
		buffer = append(buffer, metric)
	}
}

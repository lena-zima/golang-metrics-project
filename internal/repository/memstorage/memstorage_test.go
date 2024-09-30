package memstorage

import (
	"testing"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestPostGauge(t *testing.T) {
	stor, _ := NewMemStorage()

	tests := []struct {
		name       string
		gaugeName  string
		gaugeValue repository.Gauge
		expected   repository.Gauge
	}{
		{
			name:       "gauge 1.1",
			gaugeName:  "gauge1.1",
			gaugeValue: 1.1,
			expected:   1.1,
		},
		{
			name:       "gauge 1.2345",
			gaugeName:  "gauge1.2345",
			gaugeValue: 1.2345,
			expected:   1.2345,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stor.PostGauge(test.gaugeName, test.gaugeValue)
			got := stor.gaugeMetrics[test.gaugeName]
			assert.Equal(t, got, test.expected)
		})
	}
}

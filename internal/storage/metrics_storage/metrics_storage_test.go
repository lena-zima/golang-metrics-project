package metrics_storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGauge(t *testing.T) {

	tests := []struct {
		name       string
		gaugeName  string
		gaugeValue float64
		result     float64
	}{
		{
			name:       "Test simple gauge 1",
			gaugeName:  "test",
			gaugeValue: 1.1,
			result:     1.1,
		},
		{
			name:       "Test simple gauge 2",
			gaugeName:  "test2",
			gaugeValue: 12.12,
			result:     12.12,
		},
	}
	for _, test := range tests {

		var some_storage MemStorage

		some_storage.GaugeMetrics = make(map[string]Gauge, 0)

		some_storage.AddGauge(test.gaugeName, test.gaugeValue)
		got := float64(some_storage.GaugeMetrics[test.gaugeName])
		assert.Equal(t, got, test.result)
	}
}

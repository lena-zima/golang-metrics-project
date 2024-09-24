package memstorage

import (
	"testing"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestAddGauge(t *testing.T) {

	tests := []struct {
		name       string
		gaugeName  string
		gaugeValue repository.Gauge
		result     repository.Gauge
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

		var someStorage MemStorage

		someStorage.GaugeMetrics = make(map[string]repository.Gauge, 0)

		someStorage.PostGauge(test.gaugeName, test.gaugeValue)
		got := float64(someStorage.GaugeMetrics[test.gaugeName])
		assert.Equal(t, got, test.result)
	}
}

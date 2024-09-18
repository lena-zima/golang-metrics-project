package metricstorage

// Supported metric types
type Gauge float64
type Counter int64

// Type to store metrics
type MemStorage struct {
	GaugeMetrics   map[string]Gauge
	CounterMetrics map[string]Counter
}

func (storage *MemStorage) AddGauge(name string, value float64) {
	storage.GaugeMetrics[name] = Gauge(value)
}

func (storage *MemStorage) AddCounter(name string, value int64) {
	storage.CounterMetrics[name] += Counter(value)
}

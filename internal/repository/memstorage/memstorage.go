package memstorage

import (
	"context"

	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

// Type to store metrics in memory
type MemStorage struct {
	GaugeMetrics   map[string]repository.Gauge
	CounterMetrics map[string]repository.Counter
}

// Type to initiate new memory storage
// #question: It is the same as type MemStorage - is it OK?
type NewMemStorageParams struct {
	GaugeMetrics   map[string]repository.Gauge
	CounterMetrics map[string]repository.Counter
}

// Initiate new memory storage
// Input: ctx context.Context - ???
// Input: p *NewMemStorageParams - storage parameters
// Output: *MemStorage - metrics storage in memory
// Output: error - error while memory storage creation
func NewMemStorage(ctx context.Context, p *NewMemStorageParams) (*MemStorage, error) {

	// define variable for new memory storage
	var storage MemStorage

	// create initial structure for variable
	storage.GaugeMetrics = make(map[string]repository.Gauge)
	storage.CounterMetrics = make(map[string]repository.Counter)

	// fill in variable structure
	for k, v := range p.GaugeMetrics {
		storage.PostGauge(k, v)
	}
	for k, v := range p.CounterMetrics {
		storage.PostCounter(k, v)
	}

	//return outputs
	return &storage, nil
}

func (storage *MemStorage) GetAll() (map[string]repository.Gauge, map[string]repository.Counter) {
	return storage.GaugeMetrics, storage.CounterMetrics
}

// Get gauge metric by name
func (storage *MemStorage) GetGauge(name string) (repository.Gauge, bool) {
	value, exists := storage.GaugeMetrics[name]
	return value, exists
}

// Get counter metric by name
func (storage *MemStorage) GetCounter(name string) (repository.Counter, bool) {
	value, exists := storage.CounterMetrics[name]
	return value, exists
}

// Create or update gauge metric
func (storage *MemStorage) PostGauge(name string, value repository.Gauge) {
	storage.GaugeMetrics[name] = value
}

// Create or update counter metric
func (storage *MemStorage) PostCounter(name string, value repository.Counter) {
	storage.CounterMetrics[name] += value
}

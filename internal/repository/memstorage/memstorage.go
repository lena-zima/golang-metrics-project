package memstorage

import (
	"github.com/lena-zima/golang-metrics-project/internal/repository"
)

// Type to store metrics in memory
type memStorage struct {
	gaugeMetrics   map[string]repository.Gauge
	counterMetrics map[string]repository.Counter
}

// Initiate new memory storage
// Input: ctx context.Context - ???
// Input: p *NewMemStorageParams - storage parameters
// Output: *MemStorage - metrics storage in memory
// Output: error - error while memory storage creation
func NewMemStorage() (*memStorage, error) {

	// define variable for new memory storage
	var storage memStorage

	// create initial structure for variable
	storage.gaugeMetrics = make(map[string]repository.Gauge)
	storage.counterMetrics = make(map[string]repository.Counter)

	//return outputs
	return &storage, nil
}

func (storage *memStorage) GetAll() (map[string]repository.Gauge, map[string]repository.Counter, error) {
	return storage.gaugeMetrics, storage.counterMetrics, nil
}

// Get gauge metric by name
func (storage *memStorage) GetGauge(name string) (*repository.Gauge, error) {
	var value *repository.Gauge

	for k, v := range storage.gaugeMetrics {
		if k == name {
			value = &v
		}

	}
	return value, nil
}

// Get counter metric by name
func (storage *memStorage) GetCounter(name string) (*repository.Counter, error) {
	var value *repository.Counter

	for k, v := range storage.counterMetrics {
		if k == name {
			value = &v
		}

	}
	return value, nil
}

// Create or update gauge metric
func (storage *memStorage) PostGauge(name string, value repository.Gauge) error {
	storage.gaugeMetrics[name] = value
	return nil
}

// Create or update counter metric
func (storage *memStorage) PostCounter(name string, value repository.Counter) error {
	storage.counterMetrics[name] += value
	return nil
}

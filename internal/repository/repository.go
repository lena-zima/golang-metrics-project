package repository

// Aliases for supported metric types
type Gauge float64
type Counter int64

// Repository interface for metric operations in any type of storage (memory/file/db/etc)
type Repository interface {
	GetAll() (map[string]Gauge, map[string]Counter)
	GetGauge(name string) (Gauge, bool)
	GetCounter(name string) (Counter, bool)
	PostGauge(name string, value Gauge)
	PostCounter(name string, value Counter)
}

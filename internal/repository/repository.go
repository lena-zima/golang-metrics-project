package repository

// Aliases for supported metric types
type Gauge float64
type Counter int64

// Repository interface for metric operations in any type of storage (memory/file/db/etc)
type Repository interface {
	GetGauge(name string) (Gauge, bool, error)
	GetCounter(name string) (Counter, bool, error)
	PostGauge(name string, value Gauge)
	PostCounter(name string, value Counter)
}

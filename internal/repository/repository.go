package repository

// Aliases for supported metric types
type Gauge float64
type Counter int64

// Repository interface for metric operations in any type of storage (memory/file/db/etc)
type Repository interface {
	GetAll() (map[string]Gauge, map[string]Counter, error)
	GetGauge(name string) (*Gauge, error)
	GetCounter(name string) (*Counter, error)
	PostGauge(name string, value Gauge) error
	PostCounter(name string, value Counter) error
}

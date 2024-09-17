package storage

type Repository interface {
	AddGauge(name string, value float64)
	AddCounter(name string, value int64)
}

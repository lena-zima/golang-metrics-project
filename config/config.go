package config

// Config interface for config operations with any system element (server/agent/etc)
type Config interface {
	GetConfig()
}

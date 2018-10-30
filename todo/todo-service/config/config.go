package config

import (
	commonCfg "github.com/Microkubes/microservice-tools/config"
)

// ServiceConfig extends common config definition
type ServiceConfig struct {
	// Inherit common service config
	commonCfg.ServiceConfig
}

// ToStandardConfig maps this config to a standard config structure.
func (s *ServiceConfig) ToStandardConfig() *commonCfg.ServiceConfig {
	return &commonCfg.ServiceConfig{
		Service:         s.Service,
		SecurityConfig:  s.SecurityConfig,
		DBConfig:        s.DBConfig,
		GatewayURL:      s.GatewayURL,
		GatewayAdminURL: s.GatewayAdminURL,
	}
}

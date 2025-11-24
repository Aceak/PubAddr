package config

import (
	"os"
	"time"

	"PubAddr/internal/logger"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Addr      string `yaml:"addr"`
	EnableTCP bool   `yaml:"enable_tcp"`
	TCPAddr   string `yaml:"tcp_addr"`
	LogLevel  string `yaml:"log_level"`
}

type IPHeaderConfig struct {
	TrustedRealIPHeader string `yaml:"trusted_real_ip_header"`
}

type SecurityConfig struct {
	AccessToken       string        `yaml:"access_token"`
	EnableUABlock     bool          `yaml:"enable_ua_block"`
	RateDuration      time.Duration `yaml:"rate_duration"`
	BlockedUserAgents []string      `yaml:"blocked_user_agents"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	IPHeader IPHeaderConfig `yaml:"ip_header"`
	Security SecurityConfig `yaml:"security"`
}

func Load(path string) (*Config, error) {
	logger.Info("Loading config from %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8080"
	}

	if cfg.Server.LogLevel == "" {
		cfg.Server.LogLevel = "info"
	}

	if cfg.IPHeader.TrustedRealIPHeader == "" {
		cfg.IPHeader.TrustedRealIPHeader = "X-Real-IP"
	}

	return &cfg, nil
}

package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type VaultConfig struct {
	Address       string
	ClientTimeout time.Duration
	Token         string
	MountPath     string
	CertPaths     []string
}

func GetVaultConfig() VaultConfig {
	return VaultConfig{
		Address:       viper.GetString(VaultAddress),
		ClientTimeout: viper.GetDuration(VaultClientTimeout),
		Token:         viper.GetString(VaultToken),
		MountPath:     viper.GetString(VaultMountPath),
		CertPaths:     viper.GetStringSlice(VaultCertPaths),
	}
}

func (c *VaultConfig) Validate() {
	if c.Address == "" {
		log.Fatal("vault address is empty")
	}

	if c.ClientTimeout.Milliseconds() == 0 {
		log.Fatal("vault request timeout is empty")
	}

	if c.Token == "" {
		log.Fatal("vault token is empty")
	}

	if len(c.CertPaths) == 0 {
		log.Fatal("vault folder certs is empty")
	}
}

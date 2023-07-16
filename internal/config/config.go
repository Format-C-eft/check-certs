package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	VaultAddress       = "vault_address"
	VaultClientTimeout = "vault_client_timeout"
	VaultToken         = "vault_token"
	VaultMountPath     = "vault_mount_path"
	VaultCertPaths     = "vault_cert_paths"
)

func InitConfig() error {
	viper.SetConfigType("yaml")

	viper.SetConfigName("config")
	if viper.GetBool(localRunKey) {
		viper.SetConfigName("config_local")
	}

	viper.AddConfigPath(".k8s/")
	viper.AddConfigPath(".")

	if folder := viper.GetString(configCustomFolderKey); folder != "" {
		viper.AddConfigPath(folder)
	}

	return errors.Wrap(viper.ReadInConfig(), "viper.ReadInConfig")
}

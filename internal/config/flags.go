package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	configCustomFolderKey = "config-custom-folder"
	localRunKey           = "local-run"
	VersionCommandKey     = "version"
)

var (
	VersionCommand bool
)

func ParseServiceFlags() {
	initFlags()
	parseFlags()
}

func initFlags() {
	viper.SetEnvPrefix(strings.ReplaceAll(AppName, "-", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	pflag.String(configCustomFolderKey, "", "config custom folder")
	pflag.Bool(localRunKey, false, "local run")
	pflag.Bool(VersionCommandKey, false, "print version of application")
}

func parseFlags() {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	//_ = viper.BindEnv(app.GracefulDelayKey, strings.ToUpper(app.GracefulDelayKey))

	VersionCommand = viper.GetBool(VersionCommandKey)
}

package bootstrap

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/Format-C-eft/check-certs/internal/app"
	"github.com/Format-C-eft/check-certs/internal/config"
	"github.com/Format-C-eft/check-certs/internal/pkg/vault"
)

var vaultStoreOnce sync.Once
var vaultStore vault.Store

func newVaultStore() (vault.Store, error) {
	var err error
	vaultStoreOnce.Do(func() {
		vaultStore, err = vault.New(config.GetVaultConfig())
	})

	return vaultStore, errors.Wrap(err, "newVaultStore")
}

var appStoreOnce sync.Once
var appStore *app.Store

func newAppStore(vault vault.Store) *app.Store {
	appStoreOnce.Do(func() {
		appStore = app.New(vault)
	})

	return appStore
}

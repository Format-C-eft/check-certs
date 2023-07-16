//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"

	"github.com/google/wire"

	"github.com/Format-C-eft/check-certs/internal/app"
)

func Initialize(_ context.Context) (*app.Store, error) {
	wire.Build(
		newVaultStore,
		newAppStore,
	)
	return &app.Store{}, nil
}

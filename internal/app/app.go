package app

import (
	"github.com/Format-C-eft/check-certs/internal/pkg/vault"
)

type Store struct {
	vault vault.Store
}

func New(vault vault.Store) *Store {
	return &Store{
		vault: vault,
	}
}

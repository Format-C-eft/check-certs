package vault

import (
	"time"
)

const (
	vaultMaxRetries   = 100
	vaultMinRetryWait = 1 * time.Second
	vaultMaxRetryWait = 8 * time.Second

	vaultMaxJitter = 1000
)

type KeyValues map[string]map[string]string

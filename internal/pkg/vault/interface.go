package vault

import (
	"context"
)

type Store interface {
	GetAllKeys(ctx context.Context) KeyValues
}

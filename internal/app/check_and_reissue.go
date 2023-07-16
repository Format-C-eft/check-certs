package app

import (
	"context"
	"log"

	"github.com/Format-C-eft/check-certs/internal/pkg/vault"
)

func (s *Store) CheckAndReissue(ctx context.Context, done chan struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	values := s.vault.GetAllKeys(ctx)

	valuesReissue := getCertsListFromReissue(ctx, values)

	if len(valuesReissue) == 0 {
		log.Println("All certs is ok")
		return
	}

	return
}

func getCertsListFromReissue(ctx context.Context, values vault.KeyValues) vault.KeyValues {

	//tls.loa

	return nil
}

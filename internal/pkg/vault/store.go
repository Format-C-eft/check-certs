package vault

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"

	"github.com/Format-C-eft/check-certs/internal/config"
)

type store struct {
	client *api.KVv2

	certPaths []string
}

func New(cfg config.VaultConfig) (Store, error) {
	cfg.Validate()

	cfgClient := api.DefaultConfig()
	cfgClient.Address = cfg.Address

	client, err := api.NewClient(cfgClient)
	if err != nil {
		return nil, errors.Wrap(err, "api.NewClient")
	}

	client.SetToken(cfg.Token)
	client.SetClientTimeout(cfg.ClientTimeout)

	client.SetMaxRetryWait(vaultMaxRetryWait)
	client.SetMinRetryWait(vaultMinRetryWait)
	client.SetMaxRetries(vaultMaxRetries)

	client.SetCheckRetry(func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp != nil {
			if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
				log.Printf("vault: run backoff retry. statusCode: %d", resp.StatusCode)
				return true, nil
			}
		}

		return api.DefaultRetryPolicy(ctx, resp, err)
	})

	client.SetBackoff(retryablehttp.DefaultBackoff)

	return &store{
		client:    client.KVv2(cfg.MountPath),
		certPaths: cfg.CertPaths,
	}, nil
}

package vault

import (
	"context"
	"log"
)

func (s *store) GetAllKeys(ctx context.Context) KeyValues {
	result := make(KeyValues, len(s.certPaths))

	for _, path := range s.certPaths {
		values, err := s.client.Get(ctx, path)
		if err != nil {
			log.Printf("Cant get vault value from key %s: %s", path, err)
			continue
		}

		if len(values.Data) == 0 {
			log.Printf("Count values from key: %s is zero", path)
			continue
		}

		if !values.VersionMetadata.DeletionTime.IsZero() {
			log.Printf("Values from key: %s is deleted", path)
			continue
		}

		resultSecondary := make(map[string]string, len(values.Data))
		for key, value := range values.Data {
			resultSecondary[key] = value.(string)
		}

		result[path] = resultSecondary
	}

	return result
}

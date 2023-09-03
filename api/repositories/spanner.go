// repositories/spanner.go
package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"

	"cloud-spanner-go/config"
)

func NewSpannerClient(cfg *config.Config) (*spanner.Client, error) {
	ctx := context.Background()
	client, err := spanner.NewClient(
		ctx,
		fmt.Sprintf("projects/%s/instances/%s/databases/%s",
			cfg.ProjectID,
			cfg.InstanceID,
			cfg.DatabaseID,
		),
		option.WithCredentialsFile(cfg.CF),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %v", err)
	}
	return client, nil
}

package main

import (
	"context"
	"os"

	"cleaning-repos/src/domain"
	"cleaning-repos/src/services"
)

func main() {
	ctx := context.Background()

	options := domain.NewOptions(
		domain.WithWhitelist("1", "2", "3"),
		domain.WithOwner("OwnerName"),
		domain.WithListRepositoryOptions(100, domain.All),
		domain.WithDeleteEnabled(false),
		domain.WithFilename("repositories"),
	)

	err := services.ListRepository(ctx, *options)
	if err != nil {
		os.Exit(2)
	}

	services.DeleteRepository(ctx, *options)
}

package main

import (
	"context"
	"os"

	"cleaning-repos/src/services"
)

func main() {
	ctx := context.Background()
	fileName, err := services.ListRepository(ctx)
	if err != nil {
		os.Exit(2)
	}

	services.DeleteRepository(ctx, fileName, false)
}

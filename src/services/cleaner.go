package services

import (
	"cleaning-repos/src/clients/github"
	"cleaning-repos/src/domain"
	"cleaning-repos/src/logger"
	"cleaning-repos/src/services/models"
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
)

func DeleteRepository(ctx context.Context, options domain.Options) error {
	logger.Info("Reading repository file")

	filename := fmt.Sprintf("%s.csv", options.Filename)
	b, err := os.ReadFile(filename)
	if err != nil {
		logger.Error("Failed reading file",
			zap.String("Filename", filename),
			zap.Error(err))
		return err
	}

	logger.Info("Deserializing file to struct")
	var repositories []models.Repository
	if err := gocsv.UnmarshalBytes(b, &repositories); err != nil {
		logger.Error("Failed deserializing file",
			zap.Error(err))
		return err
	}

	logger.Info("Deleting repositories concurrently with goroutines")

	var wg sync.WaitGroup

	for _, repo := range repositories {
		wg.Add(1)

		go func(goroutineRepo models.Repository) {
			defer wg.Done()

			if !options.DeleteEnabled {
				logger.Info("Repository deleted",
					zap.String("Name", goroutineRepo.Name))

				return
			}

			err := github.DeleteRepository(ctx, goroutineRepo.Name, options)
			if err != nil {
				logger.Error("Failed to delete repository",
					zap.String("Repository name", goroutineRepo.Name),
					zap.Error(err))
				return
			}

			logger.Info("Repository deleted",
				zap.String("Name", goroutineRepo.Name))

		}(repo)
	}

	wg.Wait()
	logger.Info("Done")

	return nil
}

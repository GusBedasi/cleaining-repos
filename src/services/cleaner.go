package services

import (
	"cleaning-repos/src/clients/github"
	"cleaning-repos/src/logger"
	"cleaning-repos/src/services/models"
	"context"
	"os"
	"sync"

	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
)

func DeleteRepository(ctx context.Context, fileName string, delete bool) error {
	logger.Info("Reading repository file")
	b, err := os.ReadFile(fileName)
	if err != nil {
		logger.Error("Failed reading file",
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

			logger.Info("Deleting repository",
				zap.String("Name", goroutineRepo.Name))

			if !delete {
				logger.Info("Repository deleted",
					zap.String("Name", goroutineRepo.Name))

				return
			}

			err := github.DeleteRepository(ctx, goroutineRepo.Name)
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

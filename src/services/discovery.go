package services

import (
	"context"
	"fmt"
	"os"
	"slices"

	"cleaning-repos/src/clients/github"
	"cleaning-repos/src/domain"
	"cleaning-repos/src/logger"

	"github.com/schollz/progressbar/v3"
	"go.uber.org/zap"
)

const csv_header_name = "name"

func ListRepository(ctx context.Context, options domain.Options) error {
	logger.Info("Listing repositorires")

	repositories, err := github.ListRepositories(ctx, options)
	if err != nil {
		logger.Error("Error listing repositories on github",
			zap.String("Error", err.Error()))
		return err
	}

	logger.Info("Repositories loaded")
	logger.Info("Writing repositories to file")

	f, err := os.Create(fmt.Sprintf("%s.txt", options.Filename))
	if err != nil {
		logger.Error("Error creating a file",
			zap.String("Error", err.Error()))
		return err
	}

	defer f.Close()

	_, err = f.WriteString(csv_header_name + ",\n")
	if err != nil {
		logger.Error("Error writing csv heaeder",
			zap.String("Error", err.Error()))
		return err
	}

	size := len(repositories)
	bar := progressbar.Default(int64(size))

	for _, repo := range repositories {
		bar.Add(1)

		if !isAllowed(options.Whitelist, repo.Name) {
			continue
		}

		_, err := f.WriteString(fmt.Sprintf("%s,\n", repo.Name))
		if err != nil {
			logger.Error("Error writing repository to the file",
				zap.String("Error", err.Error()))
			return err
		}
	}

	logger.Info("Done")

	return nil
}

func isAllowed(whiteList []string, repo string) bool {
	return !slices.Contains(whiteList, repo)
}

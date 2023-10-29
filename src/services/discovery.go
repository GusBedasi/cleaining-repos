package services

import (
	"context"
	"fmt"
	"os"
	"slices"

	"cleaning-repos/src/clients/github"
	"cleaning-repos/src/logger"

	"github.com/schollz/progressbar/v3"
	"go.uber.org/zap"
)

const csv_header_name = "name"

var whiteList []string = []string{
	"GusBedasi",
	"Yield-return",
	"YieldReturn2",
	"PollyFallbackPOC",
	"groove-tech-test",
	"flappy-bird-js",
}

func ListRepository(ctx context.Context) (string, error) {
	logger.Info("Listing repositorires")

	repositories, err := github.ListRepositories(ctx)
	if err != nil {
		logger.Error("Error listing repositories on github",
			zap.String("Error", err.Error()))
		return "", err
	}

	logger.Info("Repositories loaded")
	logger.Info("Writing repositories to file")

	f, err := os.Create("repositories.txt")
	if err != nil {
		logger.Error("Error creating a file",
			zap.String("Error", err.Error()))
		return "", err
	}

	defer f.Close()

	_, err = f.WriteString(csv_header_name + ",\n")
	if err != nil {
		logger.Error("Error writing csv heaeder",
			zap.String("Error", err.Error()))
		return "", err
	}

	size := len(repositories)
	bar := progressbar.Default(int64(size))

	for _, repo := range repositories {
		bar.Add(1)

		if !isAllowed(repo.Name) {
			continue
		}

		_, err := f.WriteString(fmt.Sprintf("%s,\n", repo.Name))
		if err != nil {
			logger.Error("Error writing repository to the file",
				zap.String("Error", err.Error()))
			return "", err
		}
	}

	logger.Info("Done")

	return f.Name(), nil
}

func isAllowed(repo string) bool {
	if slices.Contains(whiteList, repo) {
		return false
	}

	return true
}

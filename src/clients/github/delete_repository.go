package github

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"cleaning-repos/src/logger"

	"go.uber.org/zap"
)

const delete_repo_path = "repos/GusBedasi"

func DeleteRepository(ctx context.Context, repoName string) error {
	uri := fmt.Sprintf("%s/%s/%s", options.BaseUrl, delete_repo_path, repoName)

	logger.Info("Making request",
		zap.String("URI", uri))

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", options.ApiKey))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		logger.Error("Error loading repositories",
			zap.String("Status code", resp.Status),
			zap.String("Message:", string(body)))
	}

	return nil
}

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"cleaning-repos/src/clients/github/models"
	"cleaning-repos/src/logger"

	"go.uber.org/zap"
)

const list_repo_path = "users/GusBedasi/repos"

func ListRepositories(ctx context.Context) ([]models.RespositoryResponse, error) {
	uri := fmt.Sprintf("%s/%s?", options.BaseUrl, list_repo_path)

	params := url.Values{
		"type":     {"private"},
		"per_page": {"100"},
	}

	uri = uri + params.Encode()

	logger.Info("Making request",
		zap.String("URI", uri))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", options.ApiKey))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("Error loading repositories",
			zap.String("Status code", resp.Status),
			zap.String("Message:", string(body)))
	}

	repository := &[]models.RespositoryResponse{}
	err = json.Unmarshal([]byte(body), repository)
	if err != nil {
		return nil, err
	}

	return *repository, nil
}

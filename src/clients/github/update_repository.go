package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cleaning-repos/src/clients/github/models"
	"cleaning-repos/src/domain"
)

func ModifyRepositoryVisbility(
	ctx context.Context,
	repoName string,
	request models.UpdateRepositoryRequest,
	options domain.Options) error {

	resource := fmt.Sprintf("repos/%s", options.Owner)
	uri := fmt.Sprintf("%s/%s/%s", githubOption.BaseUrl, resource, repoName)

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, uri, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", githubOption.ApiKey))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	repository := &[]models.RespositoryResponse{}
	err = json.Unmarshal([]byte(responseBody), repository)
	if err != nil {
		return err
	}

	return nil
}

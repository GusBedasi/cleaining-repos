package github

import "os"

type githubOptions struct {
	ApiKey  string
	BaseUrl string
}

const (
	github_api_key  = "github_api_key"
	github_base_url = "https://api.github.com"
)

var githubOption *githubOptions

func init() {
	githubOption = &githubOptions{
		ApiKey:  os.Getenv(github_api_key),
		BaseUrl: github_base_url,
	}
}

package gitlfs

import (
	"github.com/dre2004/go-nexus-client/nexus3/pkg/client"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	gitLfsAPIEndpoint = common.RepositoryAPIEndpoint + "/gitlfs"
)

type RepositoryGitLfsService struct {
	client *client.Client

	Hosted *RepositoryGitLfsHostedService
}

func NewRepositoryGitLfsService(c *client.Client) *RepositoryGitLfsService {
	return &RepositoryGitLfsService{
		client: c,

		Hosted: NewRepositoryGitLfsHostedService(c),
	}
}

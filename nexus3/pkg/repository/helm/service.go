package helm

import (
	"github.com/dre2004/go-nexus-client/nexus3/pkg/client"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	helmAPIEndpoint = common.RepositoryAPIEndpoint + "/helm"
)

type RepositoryHelmService struct {
	client *client.Client

	Hosted *RepositoryHelmHostedService
	Proxy  *RepositoryHelmProxyService
}

func NewRepositoryHelmService(c *client.Client) *RepositoryHelmService {
	return &RepositoryHelmService{
		client: c,

		Hosted: NewRepositoryHelmHostedService(c),
		Proxy:  NewRepositoryHelmProxyService(c),
	}
}

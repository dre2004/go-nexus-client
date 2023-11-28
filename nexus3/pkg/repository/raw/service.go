package raw

import (
	"github.com/dre2004/go-nexus-client/nexus3/pkg/client"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	rawAPIEndpoint = common.RepositoryAPIEndpoint + "/raw"
)

type RepositoryRawService struct {
	client *client.Client

	Group  *RepositoryRawGroupService
	Hosted *RepositoryRawHostedService
	Proxy  *RepositoryRawProxyService
}

func NewRepositoryRawService(c *client.Client) *RepositoryRawService {
	return &RepositoryRawService{
		client: c,

		Group:  NewRepositoryRawGroupService(c),
		Hosted: NewRepositoryRawHostedService(c),
		Proxy:  NewRepositoryRawProxyService(c),
	}
}

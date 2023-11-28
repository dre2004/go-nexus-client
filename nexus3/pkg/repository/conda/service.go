package conda

import (
	"github.com/dre2004/go-nexus-client/nexus3/pkg/client"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	condaAPIEndpoint = common.RepositoryAPIEndpoint + "/conda"
)

type RepositoryCondaService struct {
	client *client.Client

	Proxy *RepositoryCondaProxyService
}

func NewRepositoryCondaService(c *client.Client) *RepositoryCondaService {
	return &RepositoryCondaService{
		client: c,

		Proxy: NewRepositoryCondaProxyService(c),
	}
}

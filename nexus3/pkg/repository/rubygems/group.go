package rubygems

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dre2004/go-nexus-client/nexus3/pkg/client"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/dre2004/go-nexus-client/nexus3/pkg/tools"
	"github.com/dre2004/go-nexus-client/nexus3/schema/repository"
)

const (
	rubyGemsGroupAPIEndpoint = rubyGemsAPIEndpoint + "/group"
)

type RepositoryRubyGemsGroupService struct {
	client *client.Client
}

func NewRepositoryRubyGemsGroupService(c *client.Client) *RepositoryRubyGemsGroupService {
	return &RepositoryRubyGemsGroupService{
		client: c,
	}
}

func (s *RepositoryRubyGemsGroupService) Create(repo repository.RubyGemsGroupRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(rubyGemsGroupAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsGroupService) Get(id string) (*repository.RubyGemsGroupRepository, error) {
	var repo repository.RubyGemsGroupRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", rubyGemsGroupAPIEndpoint, id), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &repo); err != nil {
		return nil, fmt.Errorf("could not unmarshal repository: %v", err)
	}
	return &repo, nil
}

func (s *RepositoryRubyGemsGroupService) Update(id string, repo repository.RubyGemsGroupRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", rubyGemsGroupAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsGroupService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}

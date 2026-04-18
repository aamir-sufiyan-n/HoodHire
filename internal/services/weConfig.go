package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
)

type WebConfigserv struct {
	Repo *repositories.WebRepo
}

func NewWebConfigService(repo *repositories.WebRepo) *WebConfigserv {
    return &WebConfigserv{Repo: repo}
}

func ( s *WebConfigserv)GetAllConfig()([]models.WebConfig,error){
	return s.Repo.GetAllConfig()
}
func (s *WebConfigserv) ToggleConfig(key string, isActive bool) error {
    _, err := s.Repo.GetConfigByKey(key)
    if err != nil {
        return errors.New("config key not found")
    }
    return s.Repo.UpdateConfig(key, isActive)
}

func (s *WebConfigserv) IsFeatureActive(key string) (bool, error) {
    config, err := s.Repo.GetConfigByKey(key)
    if err != nil {
        return false, err
    }
    return config.IsActive, nil
}
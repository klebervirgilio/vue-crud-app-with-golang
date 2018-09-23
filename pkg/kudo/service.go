package kudo

import (
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/core"
)

type GitHubRepo struct {
	RepoID   string `json:"repo_id"`
	RepoURL  string `json:"repo_url"`
	RepoName string `json:"name"`
	Language string `json:"language"`
}

type Service struct {
	userId string
	repo   core.Repository
}

func (s Service) GetKudos() ([]*core.Kudo, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

func (s Service) CreateKudoFor(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := &core.Kudo{
		UserID:   s.userId,
		RepoID:   githubRepo.RepoID,
		RepoName: githubRepo.RepoName,
		RepoURL:  githubRepo.RepoURL,
	}
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) UpdateKudoWith(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := &core.Kudo{
		UserID:   s.userId,
		RepoID:   githubRepo.RepoID,
		RepoName: githubRepo.RepoName,
		RepoURL:  githubRepo.RepoURL,
	}
	err := s.repo.Update(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) RemoveKudo(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := &core.Kudo{
		UserID:   s.userId,
		RepoID:   githubRepo.RepoID,
		RepoName: githubRepo.RepoName,
		RepoURL:  githubRepo.RepoURL,
	}
	err := s.repo.Delete(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func NewService(repo core.Repository, userId string) Service {
	return Service{
		repo:   repo,
		userId: userId,
	}
}

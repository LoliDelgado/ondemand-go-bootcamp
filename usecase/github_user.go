//Package usecase processes the information from repository, making any needed calculations, data transformation, etc.
//If it recieved data it's from controllers, coming from delivery layer
package usecase

import (
	"github.com/LoliDelgado/ondemand-go-bootcamp/model"
	"github.com/LoliDelgado/ondemand-go-bootcamp/repository"
)

type GithubUserUseCase struct {
	githubUserRepo *repository.GithubUser
}

func NewGithubUser(githubUserRepo *repository.GithubUser) *GithubUserUseCase {
	return &GithubUserUseCase{
		githubUserRepo,
	}
}

func (g *GithubUserUseCase) FetchAll() ([]model.GithubUser, error) {
	return g.githubUserRepo.FetchAll()
}

func (g *GithubUserUseCase) GetById(id int) (model.GithubUser, error) {
	return g.githubUserRepo.GetById(id)
}

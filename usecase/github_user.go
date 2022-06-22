//Package usecase processes the information from repository, making any needed calculations, data transformation, etc.
//If it received data it's from controllers, coming from delivery layer
package usecase

import (
	"log"

	"github.com/LoliDelgado/ondemand-go-bootcamp/model"
	"github.com/LoliDelgado/ondemand-go-bootcamp/repository"
	"github.com/LoliDelgado/ondemand-go-bootcamp/util"
)

type GithubUserUseCase struct {
	githubUserRepo *repository.GithubUser
}

type IGithubUserUseCase interface {
	FetchAll() ([]model.GithubUser, error)
	GetById(id int) (model.GithubUser, error)
}

func NewGithubUser(githubUserRepo *repository.GithubUser) *GithubUserUseCase {
	return &GithubUserUseCase{
		githubUserRepo,
	}
}

// FetchAll gets info from repository layer and handles any error that might come
// Handling errors at this point provides a more detailed information to the controller layer to provide a better response
func (g *GithubUserUseCase) FetchAll() ([]model.GithubUser, *util.RequestError) {
	logger := log.Default()

	githubUsers, err := g.githubUserRepo.FetchAll()

	if err != nil {
		var responseError *util.RequestError
		errorInfo := err.WithoutContext()

		if errorInfo == repository.ErrReadingLineCSV || errorInfo == repository.ErrOpeningCSV {
			logger.Println("Error reading CSV file in " + err.Context)
			responseError = util.NewRequestError(500, err)
		}

		if responseError != nil {
			return nil, responseError
		}
	}

	return githubUsers, nil
}

// GetById gets info from repository layer and returns the githubUser matching the received id
// Handling errors at this point provides a more detailed information to the controller layer to provide a better response
func (g *GithubUserUseCase) GetById(id int) (model.GithubUser, *util.RequestError) {
	logger := log.Default()

	githubUser, err := g.githubUserRepo.GetById(id)
	if err != nil {
		var responseError *util.RequestError
		errorInfo := err.WithoutContext()

		if errorInfo == repository.ErrReadingLineCSV || errorInfo == repository.ErrOpeningCSV {
			logger.Println("Error reading CSV file in " + err.Context)
			responseError = util.NewRequestError(500, err)
		}

		if responseError != nil {
			return model.GithubUser{}, responseError
		}
	}
	return githubUser, nil
}

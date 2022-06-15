package main

import (
	"log"
	"net/http"

	"github.com/LoliDelgado/ondemand-go-bootcamp/controller"
	"github.com/LoliDelgado/ondemand-go-bootcamp/delivery"
	"github.com/LoliDelgado/ondemand-go-bootcamp/repository"
	"github.com/LoliDelgado/ondemand-go-bootcamp/usecase"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	const fileName = "github_users.csv"
	//repository
	githubUserRepo := repository.NewGithubUser(fileName)

	//useCase
	githubUserUseCase := usecase.NewGithubUser(githubUserRepo)

	//controllers for Rest
	httpRender := render.New()
	usersController := controller.NewGithubUser(httpRender, githubUserUseCase)

	httpRouter := mux.NewRouter()
	delivery.Setup(
		usersController,
		httpRouter,
	)

	//start server
	log.Fatal(http.ListenAndServe(":7000", httpRouter))
}

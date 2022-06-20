package main

import (
	"net/http"

	"github.com/LoliDelgado/ondemand-go-bootcamp/controller"
	"github.com/LoliDelgado/ondemand-go-bootcamp/delivery"
	"github.com/LoliDelgado/ondemand-go-bootcamp/repository"
	"github.com/LoliDelgado/ondemand-go-bootcamp/usecase"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	// logger
	var log = logrus.New()

	const fileName = "github_users.csv"
	// repository
	githubUserRepo := repository.NewGithubUser(fileName)

	// useCase
	githubUserUseCase := usecase.NewGithubUser(githubUserRepo)

	// controllers for Rest
	httpRender := render.New()
	usersController := controller.NewGithubUser(httpRender, githubUserUseCase)

	httpRouter := mux.NewRouter()
	delivery.Setup(
		usersController,
		httpRouter,
	)

	// start server
	log.Info("Starting server at port 7000")

	err := http.ListenAndServe(":7000", httpRouter)
	if err != nil {
		log.Fatal("starting server:", err)
	}
}

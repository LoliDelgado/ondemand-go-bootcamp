package main

import (
	"log"
	"net/http"

	"github.com/LoliDelgado/ondemand-go-bootcamp/config"
	"github.com/LoliDelgado/ondemand-go-bootcamp/controller"
	"github.com/LoliDelgado/ondemand-go-bootcamp/delivery"
	"github.com/LoliDelgado/ondemand-go-bootcamp/repository"
	"github.com/LoliDelgado/ondemand-go-bootcamp/usecase"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	// logger
	var log = logrus.New()

	// repository
	githubUserRepo := repository.NewGithubUser(config.SourceFileName, config.SourceFilePath)

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
	log.Info("Starting server at port " + config.Port)

	err = http.ListenAndServe(":"+config.Port, httpRouter)
	if err != nil {
		log.Fatal("starting server:", err)
	}
}

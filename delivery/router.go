package delivery

import (
	"github.com/LoliDelgado/ondemand-go-bootcamp/controller"

	"github.com/gorilla/mux"
)

func Setup(githubUserHandler *controller.GithubUser, router *mux.Router) {

	router.HandleFunc("/github-users", githubUserHandler.GetGithubUsers)
	router.HandleFunc("/github-users/{id}", githubUserHandler.GetGithubUserById)
}

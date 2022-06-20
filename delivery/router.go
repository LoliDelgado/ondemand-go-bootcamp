package delivery

import (
	"net/http"

	"github.com/LoliDelgado/ondemand-go-bootcamp/controller"

	"github.com/gorilla/mux"
)

// Setup defines all the available routes
func Setup(githubUserHandler *controller.GithubUser, router *mux.Router) {
	// github users routes
	router.HandleFunc("/github-users", githubUserHandler.GetGithubUsers).Methods(http.MethodGet)
	router.HandleFunc("/github-users/{id}", githubUserHandler.GetGithubUserById).Methods(http.MethodGet)
}

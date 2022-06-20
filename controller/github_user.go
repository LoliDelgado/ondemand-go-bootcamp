// Package controller provides the communication between the delivery layer and the usecase layer
package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LoliDelgado/ondemand-go-bootcamp/usecase"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// GithubUser struct defines the structure to be used in the controller, so that it can be
// communicated with the usecase and able to render a response to the router
type GithubUser struct {
	render *render.Render
	g      *usecase.GithubUserUseCase
}

// NewGithubUser initializes the controller
func NewGithubUser(r *render.Render, g *usecase.GithubUserUseCase) *GithubUser {
	return &GithubUser{r, g}
}

// GetGithubUsers fetches all the available github users from the usecase
func (c *GithubUser) GetGithubUsers(rw http.ResponseWriter, req *http.Request) {
	githubUsers, err := c.g.FetchAll()
	if err != nil {
		c.render.Text(rw, http.StatusInternalServerError, fmt.Sprintf("This is not fine🔥\nAnd the reason is: %s", err))
		return
	}

	c.render.JSON(rw, http.StatusOK, githubUsers)
}

// GetGithubUserById fetches the matching github user filtered by the requested id
func (c *GithubUser) GetGithubUserById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	numericId, err := strconv.Atoi(id)
	if err != nil {
		c.render.Text(rw, http.StatusBadRequest, "invalid id")
		return
	}

	githubUser, err := c.g.GetById(numericId)
	if err != nil {
		c.render.Text(rw, http.StatusInternalServerError, fmt.Sprintf("This is not fine🔥\nAnd the reason is: %s", err))
		return
	}

	c.render.JSON(rw, http.StatusOK, githubUser)
}

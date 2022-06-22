// Package controller provides the communication between the delivery layer and the usecase layer
package controller

import (
	"net/http"
	"strconv"

	"github.com/LoliDelgado/ondemand-go-bootcamp/model"
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

type IGithubUserControler interface {
	GetGithubUsers(rw http.ResponseWriter, req http.Request)
	GetGithubUserById(rw http.ResponseWriter, req http.Request)
}

// NewGithubUser initializes the controller
func NewGithubUser(r *render.Render, g *usecase.GithubUserUseCase) *GithubUser {
	return &GithubUser{r, g}
}

// GetGithubUsers fetches all the available github users from the usecase
func (c *GithubUser) GetGithubUsers(rw http.ResponseWriter, req *http.Request) {
	githubUsers, err := c.g.FetchAll()
	if err != nil {
		c.render.JSON(rw, err.StatusCode, map[string]interface{}{
			"message": "This is not fine🔥",
			"reason":  err.ErrorWithoutContext(),
		})
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
		c.render.JSON(rw, http.StatusBadRequest, map[string]interface{}{
			"message": "This is not fine🔥",
			"reason":  "Invalid ID",
		})
		return
	}

	githubUser, _err := c.g.GetById(numericId)
	if _err != nil {
		if _err.StatusCode >= 500 {
			c.render.JSON(rw, _err.StatusCode, map[string]interface{}{
				"message": "This is not fine🔥",
				"reason":  _err.ErrorWithoutContext(),
			})
			return
		}
		if _err.StatusCode > 400 {
			c.render.JSON(rw, http.StatusNotFound, model.GithubUser{})
			return
		}
	}

	c.render.JSON(rw, http.StatusOK, githubUser)
}

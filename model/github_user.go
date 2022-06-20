package model

type GithubUser struct {
	ID          int
	Login       string
	Name        string
	Company     string
	Bio         string
	PublicRepos int
}

func NewGithubUser(g GithubUser) GithubUser {
	return GithubUser{
		ID:          g.ID,
		Login:       g.Login,
		Name:        g.Name,
		Company:     g.Company,
		Bio:         g.Bio,
		PublicRepos: g.PublicRepos,
	}
}

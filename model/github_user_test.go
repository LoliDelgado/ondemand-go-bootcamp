package model

import "testing"

func TestNewGithubUser(t *testing.T) {
	tests := []struct {
		name       string
		githubuser GithubUser
		expected   GithubUser
	}{
		{
			name: "Create a new github user",
			githubuser: GithubUser{
				ID:          1,
				Login:       "madodela",
				Name:        "Loli Delgado",
				Company:     "Wizeline",
				Bio:         "Software Engineer",
				PublicRepos: 20,
			},
			expected: GithubUser{
				ID:          1,
				Login:       "madodela",
				Name:        "Loli Delgado",
				Company:     "Wizeline",
				Bio:         "Software Engineer",
				PublicRepos: 20,
			},
		},
		{
			name: "Create a new github user without some fields",
			githubuser: GithubUser{
				ID:      1,
				Login:   "madodela",
				Company: "Wizeline",
			},
			expected: GithubUser{
				ID:          1,
				Login:       "madodela",
				Name:        "",
				Company:     "Wizeline",
				Bio:         "",
				PublicRepos: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGithubUser(tt.githubuser)

			if got != tt.expected {
				t.Errorf("NewGithuUser(%+v,%+v,%+v,%+v,%+v,%+v) = %+v; expected %+v", tt.githubuser.ID, tt.githubuser.Login, tt.githubuser.Name, tt.githubuser.Company, tt.githubuser.Bio, tt.githubuser.PublicRepos, got, tt.expected)
			}
		})
	}
}

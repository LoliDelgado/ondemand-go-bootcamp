//Package repository container the fuctions to handle githubusers model information
package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/LoliDelgado/ondemand-go-bootcamp/model"
)

type GithubUser struct {
	githubUserCsvFileName string
}

type IGithubUserRepository interface {
	FetchAll() ([]model.GithubUser, error)
	GetById(id int) (model.GithubUser, error)
}

const staticFilesPath = "static/"

func NewGithubUser(fileName string) *GithubUser {
	return &GithubUser{
		fileName,
	}
}

// FetchAll returns an array of all the available users in the csv file, and any error if exist
func (g *GithubUser) FetchAll() ([]model.GithubUser, error) {
	lines, err := g.readCsv()
	if err != nil {
		return nil, err
	}
	return arrayToGithubUser(lines), nil
}

// GetById returns a github user if found in the csv file, and any error if exist
func (g *GithubUser) GetById(id int) (model.GithubUser, error) {
	lines, err := g.readCsv()
	if err != nil {
		return model.GithubUser{}, fmt.Errorf("Error getting data from csv file: %w", err)
	}
	var user [][]string
	for _, line := range lines {
		idFromFile, err := strconv.Atoi(line[0])
		if err != nil {
			return model.GithubUser{}, fmt.Errorf("Eror converting id from file: %w", err)
		}
		if idFromFile == id {
			user = append(user, line)
			break
		}
	}

	if len(user) == 1 {
		githubUser := user[0]
		publicRepos, err := strconv.Atoi(githubUser[5])
		if err != nil {
			publicRepos = 0
		}
		return model.NewGithubUser(model.GithubUser{ID: id, Login: githubUser[1], Name: githubUser[2], Company: githubUser[3], Bio: githubUser[4], PublicRepos: publicRepos}), nil
	}
	return model.GithubUser{}, errors.New("Github User not found")
}

func (g *GithubUser) readCsv() ([][]string, error) {
	//csv reader
	file, err := os.Open(staticFilesPath + g.githubUserCsvFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	var lines [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if _, err := strconv.Atoi(record[0]); err == nil {
			lines = append(lines, record)
		}
	}
	return lines, nil
}

func arrayToGithubUser(lines [][]string) []model.GithubUser {
	// Create a slice of length 0, but capacity for the amount of data received
	githubUsers := make([]model.GithubUser, 0, len(lines))
	for _, line := range lines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		publicRepos, err := strconv.Atoi(line[5])
		if err != nil {
			publicRepos = 0
		}

		data := model.NewGithubUser(model.GithubUser{ID: id, Login: line[1], Name: line[2], Company: line[3], Bio: line[4], PublicRepos: publicRepos})
		githubUsers = append(githubUsers, data)
	}
	return githubUsers
}

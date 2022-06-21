//Package repository container the fuctions to handle githubusers model information
package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/LoliDelgado/ondemand-go-bootcamp/model"
	"github.com/LoliDelgado/ondemand-go-bootcamp/util"
)

type GithubUser struct {
	csvFileName string
	csvFilePath string
}

type IGithubUserRepository interface {
	FetchAll() ([]model.GithubUser, error)
	GetById(id int) (model.GithubUser, error)
}

func NewGithubUser(fileName, filePath string) *GithubUser {
	return &GithubUser{
		fileName,
		filePath,
	}
}

// FetchAll returns an array of all the available users in the csv file, and any error if exist
func (g *GithubUser) FetchAll() ([]model.GithubUser, *util.WrappedError) {
	lines, err := g.readCsv()
	if err != nil {
		return nil, util.Wrap(err, "repo")
	}

	return arrayToGithubUser(lines), nil
}

// GetById returns a github user if found in the csv file, and any error if exist
func (g *GithubUser) GetById(id int) (model.GithubUser, *util.WrappedError) {
	lines, err := g.readCsv()
	if err != nil {
		return model.GithubUser{}, util.Wrap(err, "repo")
	}

	var user [][]string
	for _, line := range lines {
		idFromFile, err := strconv.Atoi(line[0])
		if err != nil {
			return model.GithubUser{}, util.Wrap(ErrInvalidIdCSV, "repo")
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

	return model.GithubUser{}, util.Wrap(ErrGithubUserNotFound, "repo")
}

// readCsv is in chage of opening and getting the info stored in a csv file if available
func (g *GithubUser) readCsv() ([][]string, error) {
	file, err := os.Open(g.csvFilePath + "/" + g.csvFileName)
	if err != nil {
		return nil, ErrOpeningCSV
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
			return nil, ErrReadingLineCSV
		}
		if _, err := strconv.Atoi(record[0]); err == nil {
			lines = append(lines, record)
		}
	}

	return lines, nil
}

// arrayToGithubUser receives an array of arrays of strings with all the information read from the csv file and converts it to our expected model struct of GithubUser
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

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
	githubUsers, err := g.readCsv()
	if err != nil {
		return nil, util.Wrap(err, "repo")
	}

	return githubUsers, nil
}

// GetById returns a github user if found in the csv file, and any error if exist
func (g *GithubUser) GetById(id int) (model.GithubUser, *util.WrappedError) {
	githubUsers, err := g.readCsv()
	if err != nil {
		return model.GithubUser{}, util.Wrap(err, "repo")
	}

	var matchedUser model.GithubUser
	for _, user := range githubUsers {
		if user.ID == id {
			matchedUser = user
			break
		}
	}

	return matchedUser, nil
}

// readCsv is in chage of opening and getting the info stored in a csv file if available
func (g *GithubUser) readCsv() ([]model.GithubUser, error) {
	file, err := os.Open(g.csvFilePath + "/" + g.csvFileName)
	if err != nil {
		return nil, ErrOpeningCSV
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	var githubUsers []model.GithubUser
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return []model.GithubUser{}, ErrReadingLineCSV
		}

		if len(record) < 6 {
			return []model.GithubUser{}, ErrInvalidFileStruct
		}

		publicRepos, err := strconv.Atoi(record[5])
		if err != nil {
			publicRepos = 0
		}

		if ID, err := strconv.Atoi(record[0]); err == nil {
			githubUsers = append(githubUsers, model.GithubUser{
				ID:          ID,
				Login:       record[1],
				Name:        record[2],
				Company:     record[3],
				Bio:         record[4],
				PublicRepos: publicRepos,
			})
		}
	}

	return githubUsers, nil
}

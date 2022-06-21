package repository

import "errors"

// The following are the possible errors that could happen in the repo layer
var (
	ErrOpeningCSV         = errors.New("opening-csv-file")
	ErrReadingLineCSV     = errors.New("reading-line-csv-file")
	ErrInvalidIdCSV       = errors.New("invalid-id-csv-file")
	ErrGithubUserNotFound = errors.New("github-user-not-found")
)

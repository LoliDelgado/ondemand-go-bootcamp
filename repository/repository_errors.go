package repository

import "errors"

// The following are the possible errors that could happen in the repo layer
var (
	ErrOpeningCSV        = errors.New("opening-csv-file")
	ErrInvalidFileStruct = errors.New("invalid-csv-file-struct")
	ErrReadingLineCSV    = errors.New("reading-line-csv-file")
)

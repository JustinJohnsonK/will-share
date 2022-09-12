package utils

import "errors"

type userErrors struct {
	MissingAdmin      error
	MissingGroup      error
	MissingUserValues error
	Failed            error
}

var Errors = userErrors{
	MissingAdmin:      errors.New("missing admin details"),
	MissingUserValues: errors.New("missing user values"),
	Failed:            errors.New("operation failed"),
	MissingGroup:      errors.New("group not found"),
}

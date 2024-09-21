package blog

import (
	"errors"
)

type Error struct {
	Code int
	Err  error
}

func ValidateBlogData(data Blog) Error {
	if data.Title == "" {
		return Error{
			Code: 1,
			Err:  errors.New("invalid title"),
		}
	}

	if data.Content == "" {
		return Error{
			Code: 2,
			Err:  errors.New("invalid content"),
		}
	}

	if data.Category == "" {
		return Error{
			Code: 3,
			Err:  errors.New("invalid category"),
		}
	}

	if len(data.Tags) == 0 {
		return Error{
			Code: 4,
			Err:  errors.New("invalid tags"),
		}
	}

	return Error{
		Code: 0,
	}
}

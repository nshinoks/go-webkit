package errors

import (
	stdErrors "errors"
	"net/http"
)

type HTTPError interface {
	error
	Problem() Problem
}

type httpErr struct {
	problem Problem
	cause   error
}

func (e httpErr) Error() string {
	if e.cause != nil {
		return e.cause.Error()
	}
	if e.problem.Detail != "" {
		return e.problem.Detail
	}
	return e.problem.Title
}
func (e httpErr) Unwrap() error { return e.cause }
func (e httpErr) Problem() Problem {
	p := e.problem
	if p.Status == 0 {
		p.Status = http.StatusInternalServerError
	}
	return p
}

func New(status int, title, detail string, cause error) error {
	return httpErr{
		problem: Problem{
			Status: status,
			Title:  title,
			Detail: detail,
		},
		cause: cause,
	}
}

func BadRequest(detail string) error { return New(http.StatusBadRequest, "Bad Request", detail, nil) }
func NotFound(detail string) error   { return New(http.StatusNotFound, "Not Found", detail, nil) }
func Unauthorized(detail string) error {
	return New(http.StatusUnauthorized, "Unauthorized", detail, nil)
}

func ToProblem(err error) Problem {
	if err == nil {
		return Problem{Status: http.StatusInternalServerError, Title: "Internal Server Error"}
	}
	var he HTTPError
	if stdErrors.As(err, &he) {
		return he.Problem()
	}
	return Problem{
		Status: http.StatusInternalServerError,
		Title:  "Internal Server Error",
		Detail: "unexpected error",
	}
}

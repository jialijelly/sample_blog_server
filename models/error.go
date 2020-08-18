package models

import (
	"fmt"
	"net/http"
)

type Error interface {
	Response() Response
	StatusCode() int
}

type notFound struct {
	ID string
}

func NotFound(id string) Error {
	return &notFound{ID: id}
}

func (e *notFound) Response() Response {
	return Response{
		Status:  e.StatusCode(),
		Message: fmt.Sprintf("Article with id %v not found.", e.ID),
	}
}

func (e *notFound) StatusCode() int {
	return http.StatusNotFound
}

type missingRequiredParam struct {
	param string
}

func MissingRequiredParam(param string) Error {
	return &missingRequiredParam{param: param}
}

func (e *missingRequiredParam) Response() Response {
	return Response{
		Status:  e.StatusCode(),
		Message: fmt.Sprintf("Missing required field '%v' in request body.", e.param),
	}
}

func (e *missingRequiredParam) StatusCode() int {
	return http.StatusBadRequest
}

type failedToReadRequest struct{}

func FailedToReadRequest() Error {
	return &failedToReadRequest{}
}

func (e *failedToReadRequest) Response() Response {
	return Response{
		Status:  e.StatusCode(),
		Message: "Failed to read request body.",
	}
}

func (e *failedToReadRequest) StatusCode() int {
	return http.StatusInternalServerError
}

type databaseError struct{}

func DatabaseError() Error {
	return &databaseError{}
}

func (e *databaseError) Response() Response {
	return Response{
		Status:  e.StatusCode(),
		Message: "Database error.",
	}
}

func (e *databaseError) StatusCode() int {
	return http.StatusInternalServerError
}

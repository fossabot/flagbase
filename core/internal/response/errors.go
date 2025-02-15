package response

import "github.com/sirupsen/logrus"

// Errors array of standard errors
type Errors struct {
	Errors []*Error `json:"errors"`
}

// Error standard error response object, used inside an error response
type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// IsEmpty returns true if there are no errors
func (e *Errors) IsEmpty() bool {
	return len(e.Errors) == 0
}

// Append new error to an existing error object
func (e *Errors) Append(c string, m string) {
	e.Errors = append(e.Errors, &Error{
		Code:    c,
		Message: m,
	})
	logrus.WithFields(logrus.Fields{
		"code": c,
	}).Error(m)
}

// Extend extend an existing errors response with another errors response
func (e *Errors) Extend(_e *Errors) {
	e.Errors = append(e.Errors, _e.Errors...)
}

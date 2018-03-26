package filesystem

import "errors"

// ErrEmptyLine is returned when the line to parse is empty or only contains a comment.
var ErrEmptyLine = errors.New("the line is empty")

// InvalidLineError is the error returned when a line is invalid.
type InvalidLineError struct {
	line string
}

// newInvalidLineError returns an InvalidLineError.
func newInvalidLineError(line string) *InvalidLineError {
	return &InvalidLineError{line: line}
}

// Error implements the error interface.
func (e *InvalidLineError) Error() string {
	return "invalid line: " + e.line
}

// InvalidPositionError is the error returned when a position is invalid.
type InvalidPositionError struct {
	position string
}

// newInvalidPositionError returns an InvalidPositionError.
func newInvalidPositionError(position string) *InvalidPositionError {
	return &InvalidPositionError{position: position}
}

// Error implements the error interface.
func (e *InvalidPositionError) Error() string {
	return "invalid position: " + e.position
}

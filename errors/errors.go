package errors

import "errors"

var (
	// ResourceNotFound represents when a specific resource was not found regardless of a layer.
	ResourceNotFound = errors.New("resource was not found")
)

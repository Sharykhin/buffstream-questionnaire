package errors

import (
	"errors"
	"fmt"
	"os"
)

var (
	// ResourceNotFound represents when a specific resource was not found regardless of a layer.
	ResourceNotFound = errors.New("resource was not found")
)

type DeferredFunc func() error

// CheckDefferError checks error on any deferred call
func CheckDefferError(fn DeferredFunc) {
	err := fn()
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("got error on deferred call: %v", err))
	}
}

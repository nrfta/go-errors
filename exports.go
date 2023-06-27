//go:build go1.13
// +build go1.13

package errors

import "errors"

var (
	Is = errors.Is
	As = errors.As
)

package ecode

import (
	"fmt"
	"strconv"

	"github.com/subscan-explorer/subscan-common/core/ecode/types"
)

// Error new status with code and message
func Error(code Code, message string) *Status {
	return &Status{s: &types.Status{Code: int32(code.Code()), Message: message}}
}

// Errorf new status with code and message
func Errorf(code Code, format string, args ...interface{}) *Status {
	return Error(code, fmt.Sprintf(format, args...))
}

var _ Codes = &Status{}

// Status statusError is an alias of a status proto
// implement ecode.Codes
type Status struct {
	s *types.Status
}

// Error implement error
func (s *Status) Error() string {
	return s.Message()
}

// Code return error code
func (s *Status) Code() int {
	return int(s.s.Code)
}

// Message return error message for developer
func (s *Status) Message() string {
	if s.s.Message == "" {
		return strconv.Itoa(int(s.s.Code))
	}
	return s.s.Message
}

// Proto return origin protobuf message
func (s *Status) Proto() *types.Status {
	return s.s
}

// FromCode create status from ecode
func FromCode(code Code) *Status {
	return &Status{s: &types.Status{Code: int32(code), Message: code.Message()}}
}

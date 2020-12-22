// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You don't have permission to perform the action specified in the request.
type AccessDeniedException struct {
	Message *string

	Code *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Internal server error.
type InternalException struct {
	Message *string

	Code *string
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string             { return "InternalException" }
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// AWS Security Hub isn't enabled for the account used to make this request.
type InvalidAccessException struct {
	Message *string

	Code *string
}

func (e *InvalidAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAccessException) ErrorCode() string             { return "InvalidAccessException" }
func (e *InvalidAccessException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because you supplied an invalid or out-of-range value
// for an input parameter.
type InvalidInputException struct {
	Message *string

	Code *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to create resources beyond the
// current AWS account limits. The error code describes the limit exceeded.
type LimitExceededException struct {
	Message *string

	Code *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified in the request conflicts with an existing resource.
type ResourceConflictException struct {
	Message *string

	Code *string
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because we can't find the specified resource.
type ResourceNotFoundException struct {
	Message *string

	Code *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

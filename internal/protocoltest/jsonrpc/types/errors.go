// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string

	TopLevel *string
	Nested   *ComplexNestedErrorData
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string             { return "ComplexError" }
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ErrorWithMembers struct {
	Message *string

	Code         *string
	ComplexData  *KitchenSink
	IntegerField *int32
	ListField    []string
	MapField     map[string]string
	StringField  *string
}

func (e *ErrorWithMembers) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ErrorWithMembers) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ErrorWithMembers) ErrorCode() string             { return "ErrorWithMembers" }
func (e *ErrorWithMembers) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ErrorWithoutMembers struct {
	Message *string
}

func (e *ErrorWithoutMembers) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ErrorWithoutMembers) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ErrorWithoutMembers) ErrorCode() string             { return "ErrorWithoutMembers" }
func (e *ErrorWithoutMembers) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// This error has test cases that test some of the dark corners of Amazon service
// framework history. It should only be implemented by clients.
type FooError struct {
	Message *string
}

func (e *FooError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FooError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FooError) ErrorCode() string             { return "FooError" }
func (e *FooError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string             { return "InvalidGreeting" }
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

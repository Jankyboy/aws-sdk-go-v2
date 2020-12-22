// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The user has the required permissions, so the request would have succeeded, but
// a dry run was performed.
type DryRunOperationException struct {
	Message *string
}

func (e *DryRunOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DryRunOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DryRunOperationException) ErrorCode() string             { return "DryRunOperationException" }
func (e *DryRunOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal error occurred.
type InternalError struct {
	Message *string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalError) ErrorCode() string             { return "InternalError" }
func (e *InternalError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A specified parameter is not valid.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A required parameter is missing.
type MissingRequiredParameterException struct {
	Message *string
}

func (e *MissingRequiredParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRequiredParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRequiredParameterException) ErrorCode() string {
	return "MissingRequiredParameterException"
}
func (e *MissingRequiredParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There are no connectors available.
type NoConnectorsAvailableException struct {
	Message *string
}

func (e *NoConnectorsAvailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoConnectorsAvailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoConnectorsAvailableException) ErrorCode() string             { return "NoConnectorsAvailableException" }
func (e *NoConnectorsAvailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This operation is not allowed.
type OperationNotPermittedException struct {
	Message *string
}

func (e *OperationNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationNotPermittedException) ErrorCode() string             { return "OperationNotPermittedException" }
func (e *OperationNotPermittedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified replication job already exists.
type ReplicationJobAlreadyExistsException struct {
	Message *string
}

func (e *ReplicationJobAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationJobAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationJobAlreadyExistsException) ErrorCode() string {
	return "ReplicationJobAlreadyExistsException"
}
func (e *ReplicationJobAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified replication job does not exist.
type ReplicationJobNotFoundException struct {
	Message *string
}

func (e *ReplicationJobNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationJobNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationJobNotFoundException) ErrorCode() string {
	return "ReplicationJobNotFoundException"
}
func (e *ReplicationJobNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the number of on-demand replication runs you can request in a
// 24-hour period.
type ReplicationRunLimitExceededException struct {
	Message *string
}

func (e *ReplicationRunLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationRunLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationRunLimitExceededException) ErrorCode() string {
	return "ReplicationRunLimitExceededException"
}
func (e *ReplicationRunLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified server cannot be replicated.
type ServerCannotBeReplicatedException struct {
	Message *string
}

func (e *ServerCannotBeReplicatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerCannotBeReplicatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerCannotBeReplicatedException) ErrorCode() string {
	return "ServerCannotBeReplicatedException"
}
func (e *ServerCannotBeReplicatedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is temporarily unavailable.
type TemporarilyUnavailableException struct {
	Message *string
}

func (e *TemporarilyUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TemporarilyUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TemporarilyUnavailableException) ErrorCode() string {
	return "TemporarilyUnavailableException"
}
func (e *TemporarilyUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// You lack permissions needed to perform this operation. Check your IAM policies,
// and ensure that you are using the correct access keys.
type UnauthorizedOperationException struct {
	Message *string
}

func (e *UnauthorizedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedOperationException) ErrorCode() string             { return "UnauthorizedOperationException" }
func (e *UnauthorizedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

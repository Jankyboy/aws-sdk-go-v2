// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateReplicationJob struct {
}

func (*validateOpCreateReplicationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateReplicationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateReplicationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateReplicationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAppValidationConfiguration struct {
}

func (*validateOpDeleteAppValidationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAppValidationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAppValidationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAppValidationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteReplicationJob struct {
}

func (*validateOpDeleteReplicationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteReplicationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteReplicationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteReplicationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateConnector struct {
}

func (*validateOpDisassociateConnector) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateConnectorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAppValidationConfiguration struct {
}

func (*validateOpGetAppValidationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAppValidationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAppValidationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAppValidationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAppValidationOutput struct {
}

func (*validateOpGetAppValidationOutput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAppValidationOutput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAppValidationOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAppValidationOutputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetReplicationRuns struct {
}

func (*validateOpGetReplicationRuns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetReplicationRuns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetReplicationRunsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetReplicationRunsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpNotifyAppValidationOutput struct {
}

func (*validateOpNotifyAppValidationOutput) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpNotifyAppValidationOutput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*NotifyAppValidationOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpNotifyAppValidationOutputInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAppValidationConfiguration struct {
}

func (*validateOpPutAppValidationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAppValidationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAppValidationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAppValidationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartOnDemandAppReplication struct {
}

func (*validateOpStartOnDemandAppReplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartOnDemandAppReplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartOnDemandAppReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartOnDemandAppReplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartOnDemandReplicationRun struct {
}

func (*validateOpStartOnDemandReplicationRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartOnDemandReplicationRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartOnDemandReplicationRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartOnDemandReplicationRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateReplicationJob struct {
}

func (*validateOpUpdateReplicationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateReplicationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateReplicationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateReplicationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateReplicationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateReplicationJob{}, middleware.After)
}

func addOpDeleteAppValidationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAppValidationConfiguration{}, middleware.After)
}

func addOpDeleteReplicationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteReplicationJob{}, middleware.After)
}

func addOpDisassociateConnectorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateConnector{}, middleware.After)
}

func addOpGetAppValidationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAppValidationConfiguration{}, middleware.After)
}

func addOpGetAppValidationOutputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAppValidationOutput{}, middleware.After)
}

func addOpGetReplicationRunsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetReplicationRuns{}, middleware.After)
}

func addOpNotifyAppValidationOutputValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpNotifyAppValidationOutput{}, middleware.After)
}

func addOpPutAppValidationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAppValidationConfiguration{}, middleware.After)
}

func addOpStartOnDemandAppReplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartOnDemandAppReplication{}, middleware.After)
}

func addOpStartOnDemandReplicationRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartOnDemandReplicationRun{}, middleware.After)
}

func addOpUpdateReplicationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateReplicationJob{}, middleware.After)
}

func validateOpCreateReplicationJobInput(v *CreateReplicationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateReplicationJobInput"}
	if v.ServerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServerId"))
	}
	if v.SeedReplicationTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SeedReplicationTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAppValidationConfigurationInput(v *DeleteAppValidationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAppValidationConfigurationInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteReplicationJobInput(v *DeleteReplicationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteReplicationJobInput"}
	if v.ReplicationJobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationJobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateConnectorInput(v *DisassociateConnectorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateConnectorInput"}
	if v.ConnectorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectorId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAppValidationConfigurationInput(v *GetAppValidationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAppValidationConfigurationInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAppValidationOutputInput(v *GetAppValidationOutputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAppValidationOutputInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetReplicationRunsInput(v *GetReplicationRunsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetReplicationRunsInput"}
	if v.ReplicationJobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationJobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpNotifyAppValidationOutputInput(v *NotifyAppValidationOutputInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotifyAppValidationOutputInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAppValidationConfigurationInput(v *PutAppValidationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAppValidationConfigurationInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartOnDemandAppReplicationInput(v *StartOnDemandAppReplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartOnDemandAppReplicationInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartOnDemandReplicationRunInput(v *StartOnDemandReplicationRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartOnDemandReplicationRunInput"}
	if v.ReplicationJobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationJobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateReplicationJobInput(v *UpdateReplicationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateReplicationJobInput"}
	if v.ReplicationJobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationJobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

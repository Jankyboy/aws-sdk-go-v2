// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateAdminAccount struct {
}

func (*validateOpAssociateAdminAccount) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateAdminAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateAdminAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateAdminAccountInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAppsList struct {
}

func (*validateOpDeleteAppsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAppsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAppsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAppsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePolicy struct {
}

func (*validateOpDeletePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProtocolsList struct {
}

func (*validateOpDeleteProtocolsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProtocolsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProtocolsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProtocolsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAppsList struct {
}

func (*validateOpGetAppsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAppsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAppsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAppsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetComplianceDetail struct {
}

func (*validateOpGetComplianceDetail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetComplianceDetail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetComplianceDetailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetComplianceDetailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPolicy struct {
}

func (*validateOpGetPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProtectionStatus struct {
}

func (*validateOpGetProtectionStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProtectionStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProtectionStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProtectionStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProtocolsList struct {
}

func (*validateOpGetProtocolsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProtocolsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProtocolsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProtocolsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetViolationDetails struct {
}

func (*validateOpGetViolationDetails) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetViolationDetails) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetViolationDetailsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetViolationDetailsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAppsLists struct {
}

func (*validateOpListAppsLists) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAppsLists) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAppsListsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAppsListsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListComplianceStatus struct {
}

func (*validateOpListComplianceStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListComplianceStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListComplianceStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListComplianceStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProtocolsLists struct {
}

func (*validateOpListProtocolsLists) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProtocolsLists) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProtocolsListsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProtocolsListsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAppsList struct {
}

func (*validateOpPutAppsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAppsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAppsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAppsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutNotificationChannel struct {
}

func (*validateOpPutNotificationChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutNotificationChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutNotificationChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutNotificationChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutPolicy struct {
}

func (*validateOpPutPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutProtocolsList struct {
}

func (*validateOpPutProtocolsList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutProtocolsList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutProtocolsListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutProtocolsListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateAdminAccountValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateAdminAccount{}, middleware.After)
}

func addOpDeleteAppsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAppsList{}, middleware.After)
}

func addOpDeletePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePolicy{}, middleware.After)
}

func addOpDeleteProtocolsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProtocolsList{}, middleware.After)
}

func addOpGetAppsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAppsList{}, middleware.After)
}

func addOpGetComplianceDetailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetComplianceDetail{}, middleware.After)
}

func addOpGetPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPolicy{}, middleware.After)
}

func addOpGetProtectionStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProtectionStatus{}, middleware.After)
}

func addOpGetProtocolsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProtocolsList{}, middleware.After)
}

func addOpGetViolationDetailsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetViolationDetails{}, middleware.After)
}

func addOpListAppsListsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAppsLists{}, middleware.After)
}

func addOpListComplianceStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListComplianceStatus{}, middleware.After)
}

func addOpListProtocolsListsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProtocolsLists{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutAppsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAppsList{}, middleware.After)
}

func addOpPutNotificationChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutNotificationChannel{}, middleware.After)
}

func addOpPutPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutPolicy{}, middleware.After)
}

func addOpPutProtocolsListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutProtocolsList{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateApp(v *types.App) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "App"}
	if v.AppName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppName"))
	}
	if v.Protocol == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Protocol"))
	}
	if v.Port == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Port"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAppsList(v []types.App) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AppsList"}
	for i := range v {
		if err := validateApp(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAppsListData(v *types.AppsListData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AppsListData"}
	if v.ListName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListName"))
	}
	if v.AppsList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppsList"))
	} else if v.AppsList != nil {
		if err := validateAppsList(v.AppsList); err != nil {
			invalidParams.AddNested("AppsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.PreviousAppsList != nil {
		if err := validatePreviousAppsList(v.PreviousAppsList); err != nil {
			invalidParams.AddNested("PreviousAppsList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicy(v *types.Policy) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Policy"}
	if v.PolicyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyName"))
	}
	if v.SecurityServicePolicyData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityServicePolicyData"))
	} else if v.SecurityServicePolicyData != nil {
		if err := validateSecurityServicePolicyData(v.SecurityServicePolicyData); err != nil {
			invalidParams.AddNested("SecurityServicePolicyData", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.ResourceTags != nil {
		if err := validateResourceTags(v.ResourceTags); err != nil {
			invalidParams.AddNested("ResourceTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePreviousAppsList(v map[string][]types.App) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PreviousAppsList"}
	for key := range v {
		if err := validateAppsList(v[key]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProtocolsListData(v *types.ProtocolsListData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProtocolsListData"}
	if v.ListName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListName"))
	}
	if v.ProtocolsList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtocolsList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceTag(v *types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceTags(v []types.ResourceTag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceTags"}
	for i := range v {
		if err := validateResourceTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSecurityServicePolicyData(v *types.SecurityServicePolicyData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SecurityServicePolicyData"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateAdminAccountInput(v *AssociateAdminAccountInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateAdminAccountInput"}
	if v.AdminAccount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdminAccount"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAppsListInput(v *DeleteAppsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAppsListInput"}
	if v.ListId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePolicyInput(v *DeletePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProtocolsListInput(v *DeleteProtocolsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProtocolsListInput"}
	if v.ListId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAppsListInput(v *GetAppsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAppsListInput"}
	if v.ListId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetComplianceDetailInput(v *GetComplianceDetailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetComplianceDetailInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if v.MemberAccount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberAccount"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPolicyInput(v *GetPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProtectionStatusInput(v *GetProtectionStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProtectionStatusInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProtocolsListInput(v *GetProtocolsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProtocolsListInput"}
	if v.ListId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ListId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetViolationDetailsInput(v *GetViolationDetailsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetViolationDetailsInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if v.MemberAccount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberAccount"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAppsListsInput(v *ListAppsListsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAppsListsInput"}
	if v.MaxResults == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MaxResults"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListComplianceStatusInput(v *ListComplianceStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListComplianceStatusInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProtocolsListsInput(v *ListProtocolsListsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProtocolsListsInput"}
	if v.MaxResults == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MaxResults"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAppsListInput(v *PutAppsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAppsListInput"}
	if v.AppsList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppsList"))
	} else if v.AppsList != nil {
		if err := validateAppsListData(v.AppsList); err != nil {
			invalidParams.AddNested("AppsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutNotificationChannelInput(v *PutNotificationChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutNotificationChannelInput"}
	if v.SnsTopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsTopicArn"))
	}
	if v.SnsRoleName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsRoleName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutPolicyInput(v *PutPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutPolicyInput"}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	} else if v.Policy != nil {
		if err := validatePolicy(v.Policy); err != nil {
			invalidParams.AddNested("Policy", err.(smithy.InvalidParamsError))
		}
	}
	if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutProtocolsListInput(v *PutProtocolsListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutProtocolsListInput"}
	if v.ProtocolsList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProtocolsList"))
	} else if v.ProtocolsList != nil {
		if err := validateProtocolsListData(v.ProtocolsList); err != nil {
			invalidParams.AddNested("ProtocolsList", err.(smithy.InvalidParamsError))
		}
	}
	if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagList"))
	} else if v.TagList != nil {
		if err := validateTagList(v.TagList); err != nil {
			invalidParams.AddNested("TagList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

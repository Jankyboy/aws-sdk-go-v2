// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateDeviceWithPlacement struct {
}

func (*validateOpAssociateDeviceWithPlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateDeviceWithPlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateDeviceWithPlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateDeviceWithPlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePlacement struct {
}

func (*validateOpCreatePlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProject struct {
}

func (*validateOpCreateProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePlacement struct {
}

func (*validateOpDeletePlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProject struct {
}

func (*validateOpDeleteProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePlacement struct {
}

func (*validateOpDescribePlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProject struct {
}

func (*validateOpDescribeProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateDeviceFromPlacement struct {
}

func (*validateOpDisassociateDeviceFromPlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateDeviceFromPlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateDeviceFromPlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateDeviceFromPlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDevicesInPlacement struct {
}

func (*validateOpGetDevicesInPlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDevicesInPlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDevicesInPlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDevicesInPlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPlacements struct {
}

func (*validateOpListPlacements) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPlacements) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPlacementsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPlacementsInput(input); err != nil {
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

type validateOpUpdatePlacement struct {
}

func (*validateOpUpdatePlacement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePlacement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePlacementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePlacementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateProject struct {
}

func (*validateOpUpdateProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateDeviceWithPlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateDeviceWithPlacement{}, middleware.After)
}

func addOpCreatePlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePlacement{}, middleware.After)
}

func addOpCreateProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProject{}, middleware.After)
}

func addOpDeletePlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePlacement{}, middleware.After)
}

func addOpDeleteProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProject{}, middleware.After)
}

func addOpDescribePlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePlacement{}, middleware.After)
}

func addOpDescribeProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProject{}, middleware.After)
}

func addOpDisassociateDeviceFromPlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateDeviceFromPlacement{}, middleware.After)
}

func addOpGetDevicesInPlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDevicesInPlacement{}, middleware.After)
}

func addOpListPlacementsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPlacements{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdatePlacementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePlacement{}, middleware.After)
}

func addOpUpdateProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProject{}, middleware.After)
}

func validateOpAssociateDeviceWithPlacementInput(v *AssociateDeviceWithPlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateDeviceWithPlacementInput"}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if v.DeviceTemplateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceTemplateName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.DeviceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePlacementInput(v *CreatePlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePlacementInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProjectInput(v *CreateProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProjectInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePlacementInput(v *DeletePlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePlacementInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProjectInput(v *DeleteProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProjectInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePlacementInput(v *DescribePlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePlacementInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProjectInput(v *DescribeProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProjectInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateDeviceFromPlacementInput(v *DisassociateDeviceFromPlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateDeviceFromPlacementInput"}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.DeviceTemplateName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceTemplateName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDevicesInPlacementInput(v *GetDevicesInPlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDevicesInPlacementInput"}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPlacementsInput(v *ListPlacementsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPlacementsInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
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

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdatePlacementInput(v *UpdatePlacementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePlacementInput"}
	if v.PlacementName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlacementName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateProjectInput(v *UpdateProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProjectInput"}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

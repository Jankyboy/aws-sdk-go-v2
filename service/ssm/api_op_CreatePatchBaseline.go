// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a patch baseline. For information about valid key and value pairs in
// PatchFilters for each supported operating system type, see PatchFilter
// (http://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html).
func (c *Client) CreatePatchBaseline(ctx context.Context, params *CreatePatchBaselineInput, optFns ...func(*Options)) (*CreatePatchBaselineOutput, error) {
	if params == nil {
		params = &CreatePatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePatchBaseline", params, optFns, addOperationCreatePatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePatchBaselineInput struct {

	// The name of the patch baseline.
	//
	// This member is required.
	Name *string

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline. For information about
	// accepted formats for lists of approved patches and rejected patches, see About
	// package name formats for approved and rejected patch lists
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	ApprovedPatches []string

	// Defines the compliance level for approved patches. This means that if an
	// approved patch is reported as missing, this is the severity of the compliance
	// violation. The default value is UNSPECIFIED.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is 'false'. Applies
	// to Linux instances only.
	ApprovedPatchesEnableNonSecurity bool

	// User-provided idempotency token.
	ClientToken *string

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to include patches in the baseline.
	GlobalFilters *types.PatchFilterGroup

	// Defines the operating system the patch baseline applies to. The Default value is
	// WINDOWS.
	OperatingSystem types.OperatingSystem

	// A list of explicitly rejected patches for the baseline. For information about
	// accepted formats for lists of approved patches and rejected patches, see About
	// package name formats for approved and rejected patch lists
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	RejectedPatches []string

	// The action for Patch Manager to take on patches included in the RejectedPackages
	// list.
	//
	// * ALLOW_AS_DEPENDENCY: A package in the Rejected patches list is
	// installed only if it is a dependency of another package. It is considered
	// compliant with the patch baseline, and its status is reported as InstalledOther.
	// This is the default action if no option is specified.
	//
	// * BLOCK: Packages in the
	// RejectedPatches list, and packages that include them as dependencies, are not
	// installed under any circumstances. If a package was installed before it was
	// added to the Rejected patches list, it is considered non-compliant with the
	// patch baseline, and its status is reported as InstalledRejected.
	RejectedPatchesAction types.PatchAction

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []types.PatchSource

	// Optional metadata that you assign to a resource. Tags enable you to categorize a
	// resource in different ways, such as by purpose, owner, or environment. For
	// example, you might want to tag a patch baseline to identify the severity level
	// of patches it specifies and the operating system family it applies to. In this
	// case, you could specify the following key name/value pairs:
	//
	// *
	// Key=PatchSeverity,Value=Critical
	//
	// * Key=OS,Value=Windows
	//
	// To add tags to an
	// existing patch baseline, use the AddTagsToResource action.
	Tags []types.Tag
}

type CreatePatchBaselineOutput struct {

	// The ID of the created patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreatePatchBaselineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePatchBaseline(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreatePatchBaseline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePatchBaseline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePatchBaseline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePatchBaselineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePatchBaselineInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePatchBaselineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePatchBaseline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreatePatchBaseline",
	}
}

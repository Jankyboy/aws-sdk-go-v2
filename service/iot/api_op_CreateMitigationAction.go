// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines an action that can be applied to audit findings by using
// StartAuditMitigationActionsTask. Only certain types of mitigation actions can be
// applied to specific check names. For more information, see Mitigation actions
// (https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-mitigation-actions.html).
// Each mitigation action can apply only one type of change.
func (c *Client) CreateMitigationAction(ctx context.Context, params *CreateMitigationActionInput, optFns ...func(*Options)) (*CreateMitigationActionOutput, error) {
	if params == nil {
		params = &CreateMitigationActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMitigationAction", params, optFns, addOperationCreateMitigationActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMitigationActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMitigationActionInput struct {

	// A friendly name for the action. Choose a friendly name that accurately describes
	// the action (for example, EnableLoggingAction).
	//
	// This member is required.
	ActionName *string

	// Defines the type of action and the parameters for that action.
	//
	// This member is required.
	ActionParams *types.MitigationActionParams

	// The ARN of the IAM role that is used to apply the mitigation action.
	//
	// This member is required.
	RoleArn *string

	// Metadata that can be used to manage the mitigation action.
	Tags []types.Tag
}

type CreateMitigationActionOutput struct {

	// The ARN for the new mitigation action.
	ActionArn *string

	// A unique identifier for the new mitigation action.
	ActionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMitigationActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMitigationAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMitigationAction{}, middleware.After)
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
	if err = addOpCreateMitigationActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMitigationAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMitigationAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateMitigationAction",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified workflow. All deployed systems and system instances that
// use the workflow will see the changes in the flow when it is redeployed. If you
// don't want this behavior, copy the workflow (creating a new workflow with a
// different ID), and update the copy. The workflow can contain only entities in
// the specified namespace.
func (c *Client) UpdateFlowTemplate(ctx context.Context, params *UpdateFlowTemplateInput, optFns ...func(*Options)) (*UpdateFlowTemplateOutput, error) {
	if params == nil {
		params = &UpdateFlowTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowTemplate", params, optFns, addOperationUpdateFlowTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFlowTemplateInput struct {

	// The DefinitionDocument that contains the updated workflow definition.
	//
	// This member is required.
	Definition *types.DefinitionDocument

	// The ID of the workflow to be updated. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:workflow:WORKFLOWNAME
	//
	// This member is required.
	Id *string

	// The version of the user's namespace. If no value is specified, the latest
	// version is used by default. Use the GetFlowTemplateRevisions if you want to find
	// earlier revisions of the flow to update.
	CompatibleNamespaceVersion *int64
}

type UpdateFlowTemplateOutput struct {

	// An object containing summary information about the updated workflow.
	Summary *types.FlowTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFlowTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFlowTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFlowTemplate{}, middleware.After)
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
	if err = addOpUpdateFlowTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "UpdateFlowTemplate",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the status of a bulk deployment.
func (c *Client) GetBulkDeploymentStatus(ctx context.Context, params *GetBulkDeploymentStatusInput, optFns ...func(*Options)) (*GetBulkDeploymentStatusOutput, error) {
	if params == nil {
		params = &GetBulkDeploymentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBulkDeploymentStatus", params, optFns, addOperationGetBulkDeploymentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBulkDeploymentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBulkDeploymentStatusInput struct {

	// The ID of the bulk deployment.
	//
	// This member is required.
	BulkDeploymentId *string
}

type GetBulkDeploymentStatusOutput struct {

	// Relevant metrics on input records processed during bulk deployment.
	BulkDeploymentMetrics *types.BulkDeploymentMetrics

	// The status of the bulk deployment.
	BulkDeploymentStatus types.BulkDeploymentStatus

	// The time, in ISO format, when the deployment was created.
	CreatedAt *string

	// Error details
	ErrorDetails []types.ErrorDetail

	// Error message
	ErrorMessage *string

	// Tag(s) attached to the resource arn.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBulkDeploymentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBulkDeploymentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBulkDeploymentStatus{}, middleware.After)
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
	if err = addOpGetBulkDeploymentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBulkDeploymentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBulkDeploymentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetBulkDeploymentStatus",
	}
}

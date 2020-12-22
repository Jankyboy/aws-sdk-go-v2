// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the endpoint groups that are associated with a listener. To see an AWS CLI
// example of listing the endpoint groups for listener, scroll down to Example.
func (c *Client) ListEndpointGroups(ctx context.Context, params *ListEndpointGroupsInput, optFns ...func(*Options)) (*ListEndpointGroupsOutput, error) {
	if params == nil {
		params = &ListEndpointGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointGroups", params, optFns, addOperationListEndpointGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointGroupsInput struct {

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
}

type ListEndpointGroupsOutput struct {

	// The list of the endpoint groups associated with a listener.
	EndpointGroups []types.EndpointGroup

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEndpointGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointGroups{}, middleware.After)
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
	if err = addOpListEndpointGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEndpointGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListEndpointGroups",
	}
}

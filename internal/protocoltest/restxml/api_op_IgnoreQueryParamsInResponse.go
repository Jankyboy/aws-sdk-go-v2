// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This example ensures that query string bound request parameters are serialized
// in the body of responses if the structure is used in both the request and
// response.
func (c *Client) IgnoreQueryParamsInResponse(ctx context.Context, params *IgnoreQueryParamsInResponseInput, optFns ...func(*Options)) (*IgnoreQueryParamsInResponseOutput, error) {
	if params == nil {
		params = &IgnoreQueryParamsInResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IgnoreQueryParamsInResponse", params, optFns, addOperationIgnoreQueryParamsInResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IgnoreQueryParamsInResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoreQueryParamsInResponseInput struct {
}

type IgnoreQueryParamsInResponseOutput struct {
	Baz *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIgnoreQueryParamsInResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpIgnoreQueryParamsInResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpIgnoreQueryParamsInResponse{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIgnoreQueryParamsInResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIgnoreQueryParamsInResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IgnoreQueryParamsInResponse",
	}
}

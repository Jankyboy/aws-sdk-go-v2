// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a subscription. If the subscription requires authentication for
// deletion, only the owner of the subscription or the topic's owner can
// unsubscribe, and an AWS signature is required. If the Unsubscribe call does not
// require authentication and the requester is not the subscription owner, a final
// cancellation message is delivered to the endpoint, so that the endpoint owner
// can easily resubscribe to the topic if the Unsubscribe request was unintended.
// This action is throttled at 100 transactions per second (TPS).
func (c *Client) Unsubscribe(ctx context.Context, params *UnsubscribeInput, optFns ...func(*Options)) (*UnsubscribeOutput, error) {
	if params == nil {
		params = &UnsubscribeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Unsubscribe", params, optFns, addOperationUnsubscribeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnsubscribeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Unsubscribe action.
type UnsubscribeInput struct {

	// The ARN of the subscription to be deleted.
	//
	// This member is required.
	SubscriptionArn *string
}

type UnsubscribeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnsubscribeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUnsubscribe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUnsubscribe{}, middleware.After)
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
	if err = addOpUnsubscribeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnsubscribe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "Unsubscribe",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes tags from the specified Kinesis data stream. Removed tags are deleted
// and cannot be recovered after this operation successfully completes. If you
// specify a tag that does not exist, it is ignored. RemoveTagsFromStream has a
// limit of five transactions per second per account.
func (c *Client) RemoveTagsFromStream(ctx context.Context, params *RemoveTagsFromStreamInput, optFns ...func(*Options)) (*RemoveTagsFromStreamOutput, error) {
	if params == nil {
		params = &RemoveTagsFromStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTagsFromStream", params, optFns, addOperationRemoveTagsFromStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsFromStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for RemoveTagsFromStream.
type RemoveTagsFromStreamInput struct {

	// The name of the stream.
	//
	// This member is required.
	StreamName *string

	// A list of tag keys. Each corresponding tag is removed from the stream.
	//
	// This member is required.
	TagKeys []string
}

type RemoveTagsFromStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsFromStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromStream{}, middleware.After)
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
	if err = addOpRemoveTagsFromStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveTagsFromStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "RemoveTagsFromStream",
	}
}

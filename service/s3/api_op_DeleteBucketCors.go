// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the cors configuration information set for the bucket. To use this
// operation, you must have permission to perform the s3:PutBucketCORS action. The
// bucket owner has this permission by default and can grant this permission to
// others. For information about cors, see Enabling Cross-Origin Resource Sharing
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon Simple
// Storage Service Developer Guide.  <p class="title"> <b>Related Resources:</b>
// </p> <ul> <li> <p> </p> </li> <li> <p> <a>RESTOPTIONSobject</a> </p> </li> </ul>
func (c *Client) DeleteBucketCors(ctx context.Context, params *DeleteBucketCorsInput, optFns ...func(*Options)) (*DeleteBucketCorsOutput, error) {
	stack := middleware.NewStack("DeleteBucketCors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteBucketCorsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteBucketCorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketCors(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteBucketCors",
			Err:           err,
		}
	}
	out := result.(*DeleteBucketCorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketCorsInput struct {
	// Specifies the bucket whose cors configuration is being deleted.
	Bucket *string
}

type DeleteBucketCorsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteBucketCorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketCors{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketCors{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketCors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketCors",
	}
}

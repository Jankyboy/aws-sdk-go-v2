// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the PublicAccessBlock configuration for an Amazon S3 bucket. To use
// this operation, you must have the s3:GetBucketPublicAccessBlock permission. For
// more information about Amazon S3 permissions, see Specifying Permissions in a
// Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// <important> <p>When Amazon S3 evaluates the <code>PublicAccessBlock</code>
// configuration for a bucket or an object, it checks the
// <code>PublicAccessBlock</code> configuration for both the bucket (or the bucket
// that contains the object) and the bucket owner's account. If the
// <code>PublicAccessBlock</code> settings are different between the bucket and the
// account, Amazon S3 uses the most restrictive combination of the bucket-level and
// account-level settings.</p> </important> <p>For more information about when
// Amazon S3 considers a bucket or an object public, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The
// Meaning of "Public"</a>.</p> <p>The following operations are related to
// <code>GetPublicAccessBlock</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html">Using
// Amazon S3 Block Public Access</a> </p> </li> <li> <p>
// <a>PutPublicAccessBlock</a> </p> </li> <li> <p> <a>GetPublicAccessBlock</a> </p>
// </li> <li> <p> <a>DeletePublicAccessBlock</a> </p> </li> </ul>
func (c *Client) GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error) {
	stack := middleware.NewStack("GetPublicAccessBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetPublicAccessBlockMiddlewares(stack)
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
	addOpGetPublicAccessBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicAccessBlock(options.Region), middleware.Before)

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
			OperationName: "GetPublicAccessBlock",
			Err:           err,
		}
	}
	out := result.(*GetPublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicAccessBlockInput struct {
	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you want
	// to retrieve.
	Bucket *string
}

type GetPublicAccessBlockOutput struct {
	// The PublicAccessBlock configuration currently in effect for this Amazon S3
	// bucket.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetPublicAccessBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetPublicAccessBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetPublicAccessBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPublicAccessBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetPublicAccessBlock",
	}
}

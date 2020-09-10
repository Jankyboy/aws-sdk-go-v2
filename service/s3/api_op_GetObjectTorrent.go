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
	"io"
)

// Return torrent files from a bucket. BitTorrent can save you bandwidth when
// you're distributing large files. For more information about BitTorrent, see
// Amazon S3 Torrent
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html).  <note> <p>You
// can get torrent only for objects that are less than 5 GB in size and that are
// not encrypted using server-side encryption with customer-provided encryption
// key.</p> </note> <p>To use GET, you must have READ access to the object.</p>
// <p>The following operation is related to <code>GetObjectTorrent</code>:</p> <ul>
// <li> <p> <a>GetObject</a> </p> </li> </ul>
func (c *Client) GetObjectTorrent(ctx context.Context, params *GetObjectTorrentInput, optFns ...func(*Options)) (*GetObjectTorrentOutput, error) {
	stack := middleware.NewStack("GetObjectTorrent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectTorrentMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetObjectTorrentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectTorrent(options.Region), middleware.Before)

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
			OperationName: "GetObjectTorrent",
			Err:           err,
		}
	}
	out := result.(*GetObjectTorrentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectTorrentInput struct {
	// The object key for which to get the information.
	Key *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// The name of the bucket containing the object for which to get the torrent files.
	Bucket *string
}

type GetObjectTorrentOutput struct {
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	// A Bencoded dictionary as defined by the BitTorrent specification
	Body io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectTorrentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObjectTorrent{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectTorrent{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObjectTorrent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObjectTorrent",
	}
}

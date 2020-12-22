// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new channel and an associated stream key to start streaming.
func (c *Client) CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) {
	if params == nil {
		params = &CreateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannel", params, optFns, addOperationCreateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelInput struct {

	// Whether the channel is authorized. Default: false.
	Authorized bool

	// Channel latency mode. Default: LOW.
	LatencyMode types.ChannelLatencyMode

	// Channel name.
	Name *string

	// See Channel$tags.
	Tags map[string]string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable resolution or bitrate, the stream probably will disconnect
	// immediately. Valid values:
	//
	// * STANDARD: Multiple qualities are generated from
	// the original input, to automatically give viewers the best experience for their
	// devices and network conditions. Vertical resolution can be up to 1080 and
	// bitrate can be up to 8.5 Mbps.
	//
	// * BASIC: Amazon IVS delivers the original input
	// to viewers. The viewer’s video-quality choice is limited to the original input.
	// Vertical resolution can be up to 480 and bitrate can be up to 1.5
	// Mbps.
	//
	// Default: STANDARD.
	Type types.ChannelType
}

type CreateChannelOutput struct {

	// Object specifying a channel.
	Channel *types.Channel

	// Object specifying a stream key.
	StreamKey *types.StreamKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannel{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "CreateChannel",
	}
}

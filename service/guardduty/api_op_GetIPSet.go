// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the IPSet specified by the ipSetId.
func (c *Client) GetIPSet(ctx context.Context, params *GetIPSetInput, optFns ...func(*Options)) (*GetIPSetOutput, error) {
	if params == nil {
		params = &GetIPSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIPSet", params, optFns, addOperationGetIPSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIPSetInput struct {

	// The unique ID of the detector that the IPSet is associated with.
	//
	// This member is required.
	DetectorId *string

	// The unique ID of the IPSet to retrieve.
	//
	// This member is required.
	IpSetId *string
}

type GetIPSetOutput struct {

	// The format of the file that contains the IPSet.
	//
	// This member is required.
	Format types.IpSetFormat

	// The URI of the file that contains the IPSet. For example:
	// https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key.
	//
	// This member is required.
	Location *string

	// The user-friendly name for the IPSet.
	//
	// This member is required.
	Name *string

	// The status of IPSet file that was uploaded.
	//
	// This member is required.
	Status types.IpSetStatus

	// The tags of the IPSet resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIPSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIPSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIPSet{}, middleware.After)
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
	if err = addOpGetIPSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIPSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIPSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetIPSet",
	}
}

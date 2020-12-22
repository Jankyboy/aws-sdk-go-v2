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

// Returns the details of the filter specified by the filter name.
func (c *Client) GetFilter(ctx context.Context, params *GetFilterInput, optFns ...func(*Options)) (*GetFilterOutput, error) {
	if params == nil {
		params = &GetFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFilter", params, optFns, addOperationGetFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFilterInput struct {

	// The unique ID of the detector that the filter is associated with.
	//
	// This member is required.
	DetectorId *string

	// The name of the filter you want to get.
	//
	// This member is required.
	FilterName *string
}

type GetFilterOutput struct {

	// Specifies the action that is to be applied to the findings that match the
	// filter.
	//
	// This member is required.
	Action types.FilterAction

	// Represents the criteria to be used in the filter for querying findings.
	//
	// This member is required.
	FindingCriteria *types.FindingCriteria

	// The name of the filter.
	//
	// This member is required.
	Name *string

	// The description of the filter.
	Description *string

	// Specifies the position of the filter in the list of current filters. Also
	// specifies the order in which this filter is applied to the findings.
	Rank int32

	// The tags of the filter resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFilter{}, middleware.After)
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
	if err = addOpGetFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetFilter",
	}
}

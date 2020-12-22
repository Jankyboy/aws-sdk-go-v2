// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Permanently deletes a SizeConstraintSet. You can't delete a
// SizeConstraintSet if it's still used in any Rules or if it still includes any
// SizeConstraint objects (any filters). If you just want to remove a
// SizeConstraintSet from a Rule, use UpdateRule. To permanently delete a
// SizeConstraintSet, perform the following steps:
//
// * Update the SizeConstraintSet
// to remove filters, if any. For more information, see UpdateSizeConstraintSet.
//
// *
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteSizeConstraintSet request.
//
// * Submit a
// DeleteSizeConstraintSet request.
func (c *Client) DeleteSizeConstraintSet(ctx context.Context, params *DeleteSizeConstraintSetInput, optFns ...func(*Options)) (*DeleteSizeConstraintSetOutput, error) {
	if params == nil {
		params = &DeleteSizeConstraintSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSizeConstraintSet", params, optFns, addOperationDeleteSizeConstraintSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSizeConstraintSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSizeConstraintSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The SizeConstraintSetId of the SizeConstraintSet that you want to delete.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet and by
	// ListSizeConstraintSets.
	//
	// This member is required.
	SizeConstraintSetId *string
}

type DeleteSizeConstraintSetOutput struct {

	// The ChangeToken that you used to submit the DeleteSizeConstraintSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSizeConstraintSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSizeConstraintSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSizeConstraintSet{}, middleware.After)
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
	if err = addOpDeleteSizeConstraintSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSizeConstraintSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSizeConstraintSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteSizeConstraintSet",
	}
}

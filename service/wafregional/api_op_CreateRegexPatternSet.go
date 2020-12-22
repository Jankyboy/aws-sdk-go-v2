// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Creates a RegexPatternSet. You then use UpdateRegexPatternSet to
// specify the regular expression (regex) pattern that you want AWS WAF to search
// for, such as B[a@]dB[o0]t. You can then configure AWS WAF to reject those
// requests. To create and configure a RegexPatternSet, perform the following
// steps:
//
// * Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of a CreateRegexPatternSet request.
//
// * Submit a
// CreateRegexPatternSet request.
//
// * Use GetChangeToken to get the change token
// that you provide in the ChangeToken parameter of an UpdateRegexPatternSet
// request.
//
// * Submit an UpdateRegexPatternSet request to specify the string that
// you want AWS WAF to watch for.
//
// For more information about how to use the AWS
// WAF API to allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) {
	if params == nil {
		params = &CreateRegexPatternSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegexPatternSet", params, optFns, addOperationCreateRegexPatternSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegexPatternSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the RegexPatternSet. You can't change Name
	// after you create a RegexPatternSet.
	//
	// This member is required.
	Name *string
}

type CreateRegexPatternSetOutput struct {

	// The ChangeToken that you used to submit the CreateRegexPatternSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string

	// A RegexPatternSet that contains no objects.
	RegexPatternSet *types.RegexPatternSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRegexPatternSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegexPatternSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegexPatternSet{}, middleware.After)
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
	if err = addOpCreateRegexPatternSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegexPatternSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegexPatternSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateRegexPatternSet",
	}
}

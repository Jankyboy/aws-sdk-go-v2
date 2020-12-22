// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a ReqeustValidator of a given RestApi.
func (c *Client) CreateRequestValidator(ctx context.Context, params *CreateRequestValidatorInput, optFns ...func(*Options)) (*CreateRequestValidatorOutput, error) {
	if params == nil {
		params = &CreateRequestValidatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRequestValidator", params, optFns, addOperationCreateRequestValidatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRequestValidatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a RequestValidator of a given RestApi.
type CreateRequestValidatorInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A Boolean flag to indicate whether to validate request body according to the
	// configured model schema for the method (true) or not (false).
	ValidateRequestBody bool

	// A Boolean flag to indicate whether to validate request parameters, true, or not
	// false.
	ValidateRequestParameters bool
}

// A set of validation rules for incoming Method requests. In OpenAPI, a
// RequestValidator of an API is defined by the
// x-amazon-apigateway-request-validators.requestValidator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validators.requestValidator.html)
// object. It the referenced using the x-amazon-apigateway-request-validator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validator)
// property. Enable Basic Request Validation in API Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)
type CreateRequestValidatorOutput struct {

	// The identifier of this RequestValidator.
	Id *string

	// The name of this RequestValidator
	Name *string

	// A Boolean flag to indicate whether to validate a request body according to the
	// configured Model schema.
	ValidateRequestBody bool

	// A Boolean flag to indicate whether to validate request parameters (true) or not
	// (false).
	ValidateRequestParameters bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRequestValidatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRequestValidator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRequestValidator{}, middleware.After)
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
	if err = addOpCreateRequestValidatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRequestValidator(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateRequestValidator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRequestValidator",
	}
}

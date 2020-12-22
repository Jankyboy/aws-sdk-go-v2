// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures an IndexField for the search domain. Used to create new fields and
// modify existing ones. You must specify the name of the domain you are
// configuring and an index field configuration. The index field configuration
// specifies a unique name, the index field type, and the options you want to
// configure for the field. The options you can specify depend on the
// IndexFieldType. If the field exists, the new configuration replaces the old one.
// For more information, see Configuring Index Fields
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DefineIndexField(ctx context.Context, params *DefineIndexFieldInput, optFns ...func(*Options)) (*DefineIndexFieldOutput, error) {
	if params == nil {
		params = &DefineIndexFieldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DefineIndexField", params, optFns, addOperationDefineIndexFieldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DefineIndexFieldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DefineIndexField operation. Specifies the
// name of the domain you want to update and the index field configuration.
type DefineIndexFieldInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// The index field and field options you want to configure.
	//
	// This member is required.
	IndexField *types.IndexField
}

// The result of a DefineIndexField request. Contains the status of the
// newly-configured index field.
type DefineIndexFieldOutput struct {

	// The value of an IndexField and its current status.
	//
	// This member is required.
	IndexField *types.IndexFieldStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDefineIndexFieldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDefineIndexField{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDefineIndexField{}, middleware.After)
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
	if err = addOpDefineIndexFieldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDefineIndexField(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDefineIndexField(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DefineIndexField",
	}
}

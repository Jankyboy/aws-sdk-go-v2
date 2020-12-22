// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new, empty repository.
func (c *Client) CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) {
	if params == nil {
		params = &CreateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRepository", params, optFns, addOperationCreateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a create repository operation.
type CreateRepositoryInput struct {

	// The name of the new repository to be created. The repository name must be unique
	// across the calling AWS account. Repository names are limited to 100
	// alphanumeric, dash, and underscore characters, and cannot include certain
	// characters. For more information about the limits on repository names, see
	// Limits (https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html) in
	// the AWS CodeCommit User Guide. The suffix .git is prohibited.
	//
	// This member is required.
	RepositoryName *string

	// A comment or description about the new repository. The description field for a
	// repository accepts all HTML characters and all valid Unicode characters.
	// Applications that do not HTML-encode the description and display it in a webpage
	// can expose users to potentially malicious code. Make sure that you HTML-encode
	// the description field in any application that uses this API to display the
	// repository description on a webpage.
	RepositoryDescription *string

	// One or more tag key-value pairs to use when tagging this repository.
	Tags map[string]string
}

// Represents the output of a create repository operation.
type CreateRepositoryOutput struct {

	// Information about the newly created repository.
	RepositoryMetadata *types.RepositoryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRepository{}, middleware.After)
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
	if err = addOpCreateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "CreateRepository",
	}
}

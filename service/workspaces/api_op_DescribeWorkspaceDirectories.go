// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the available directories that are registered with Amazon WorkSpaces.
func (c *Client) DescribeWorkspaceDirectories(ctx context.Context, params *DescribeWorkspaceDirectoriesInput, optFns ...func(*Options)) (*DescribeWorkspaceDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceDirectories", params, optFns, addOperationDescribeWorkspaceDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceDirectoriesInput struct {

	// The identifiers of the directories. If the value is null, all directories are
	// retrieved.
	DirectoryIds []string

	// The maximum number of directories to return.
	Limit *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string
}

type DescribeWorkspaceDirectoriesOutput struct {

	// Information about the directories.
	Directories []types.WorkspaceDirectory

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorkspaceDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceDirectories{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceDirectories(options.Region), middleware.Before); err != nil {
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

// DescribeWorkspaceDirectoriesAPIClient is a client that implements the
// DescribeWorkspaceDirectories operation.
type DescribeWorkspaceDirectoriesAPIClient interface {
	DescribeWorkspaceDirectories(context.Context, *DescribeWorkspaceDirectoriesInput, ...func(*Options)) (*DescribeWorkspaceDirectoriesOutput, error)
}

var _ DescribeWorkspaceDirectoriesAPIClient = (*Client)(nil)

// DescribeWorkspaceDirectoriesPaginatorOptions is the paginator options for
// DescribeWorkspaceDirectories
type DescribeWorkspaceDirectoriesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeWorkspaceDirectoriesPaginator is a paginator for
// DescribeWorkspaceDirectories
type DescribeWorkspaceDirectoriesPaginator struct {
	options   DescribeWorkspaceDirectoriesPaginatorOptions
	client    DescribeWorkspaceDirectoriesAPIClient
	params    *DescribeWorkspaceDirectoriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeWorkspaceDirectoriesPaginator returns a new
// DescribeWorkspaceDirectoriesPaginator
func NewDescribeWorkspaceDirectoriesPaginator(client DescribeWorkspaceDirectoriesAPIClient, params *DescribeWorkspaceDirectoriesInput, optFns ...func(*DescribeWorkspaceDirectoriesPaginatorOptions)) *DescribeWorkspaceDirectoriesPaginator {
	options := DescribeWorkspaceDirectoriesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeWorkspaceDirectoriesInput{}
	}

	return &DescribeWorkspaceDirectoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeWorkspaceDirectoriesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeWorkspaceDirectories page.
func (p *DescribeWorkspaceDirectoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeWorkspaceDirectoriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeWorkspaceDirectories(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeWorkspaceDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceDirectories",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns comments made on a pull request. Reaction counts might include numbers
// from user identities who were deleted after the reaction was made. For a count
// of reactions from active identities, use GetCommentReactions.
func (c *Client) GetCommentsForPullRequest(ctx context.Context, params *GetCommentsForPullRequestInput, optFns ...func(*Options)) (*GetCommentsForPullRequestOutput, error) {
	if params == nil {
		params = &GetCommentsForPullRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCommentsForPullRequest", params, optFns, addOperationGetCommentsForPullRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommentsForPullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentsForPullRequestInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests.
	//
	// This member is required.
	PullRequestId *string

	// The full commit ID of the commit in the source branch that was the tip of the
	// branch at the time the comment was made.
	AfterCommitId *string

	// The full commit ID of the commit in the destination branch that was the tip of
	// the branch at the time the pull request was created.
	BeforeCommitId *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// The name of the repository that contains the pull request.
	RepositoryName *string
}

type GetCommentsForPullRequestOutput struct {

	// An array of comment objects on the pull request.
	CommentsForPullRequestData []types.CommentsForPullRequest

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCommentsForPullRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentsForPullRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentsForPullRequest{}, middleware.After)
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
	if err = addOpGetCommentsForPullRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentsForPullRequest(options.Region), middleware.Before); err != nil {
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

// GetCommentsForPullRequestAPIClient is a client that implements the
// GetCommentsForPullRequest operation.
type GetCommentsForPullRequestAPIClient interface {
	GetCommentsForPullRequest(context.Context, *GetCommentsForPullRequestInput, ...func(*Options)) (*GetCommentsForPullRequestOutput, error)
}

var _ GetCommentsForPullRequestAPIClient = (*Client)(nil)

// GetCommentsForPullRequestPaginatorOptions is the paginator options for
// GetCommentsForPullRequest
type GetCommentsForPullRequestPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCommentsForPullRequestPaginator is a paginator for GetCommentsForPullRequest
type GetCommentsForPullRequestPaginator struct {
	options   GetCommentsForPullRequestPaginatorOptions
	client    GetCommentsForPullRequestAPIClient
	params    *GetCommentsForPullRequestInput
	nextToken *string
	firstPage bool
}

// NewGetCommentsForPullRequestPaginator returns a new
// GetCommentsForPullRequestPaginator
func NewGetCommentsForPullRequestPaginator(client GetCommentsForPullRequestAPIClient, params *GetCommentsForPullRequestInput, optFns ...func(*GetCommentsForPullRequestPaginatorOptions)) *GetCommentsForPullRequestPaginator {
	options := GetCommentsForPullRequestPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetCommentsForPullRequestInput{}
	}

	return &GetCommentsForPullRequestPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCommentsForPullRequestPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetCommentsForPullRequest page.
func (p *GetCommentsForPullRequestPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCommentsForPullRequestOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetCommentsForPullRequest(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCommentsForPullRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetCommentsForPullRequest",
	}
}

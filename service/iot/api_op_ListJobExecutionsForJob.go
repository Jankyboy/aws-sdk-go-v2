// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the job executions for a job.
func (c *Client) ListJobExecutionsForJob(ctx context.Context, params *ListJobExecutionsForJobInput, optFns ...func(*Options)) (*ListJobExecutionsForJobOutput, error) {
	if params == nil {
		params = &ListJobExecutionsForJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobExecutionsForJob", params, optFns, addOperationListJobExecutionsForJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobExecutionsForJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobExecutionsForJobInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	// The status of the job.
	Status types.JobExecutionStatus
}

type ListJobExecutionsForJobOutput struct {

	// A list of job execution summaries.
	ExecutionSummaries []types.JobExecutionSummaryForJob

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobExecutionsForJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobExecutionsForJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobExecutionsForJob{}, middleware.After)
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
	if err = addOpListJobExecutionsForJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobExecutionsForJob(options.Region), middleware.Before); err != nil {
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

// ListJobExecutionsForJobAPIClient is a client that implements the
// ListJobExecutionsForJob operation.
type ListJobExecutionsForJobAPIClient interface {
	ListJobExecutionsForJob(context.Context, *ListJobExecutionsForJobInput, ...func(*Options)) (*ListJobExecutionsForJobOutput, error)
}

var _ ListJobExecutionsForJobAPIClient = (*Client)(nil)

// ListJobExecutionsForJobPaginatorOptions is the paginator options for
// ListJobExecutionsForJob
type ListJobExecutionsForJobPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobExecutionsForJobPaginator is a paginator for ListJobExecutionsForJob
type ListJobExecutionsForJobPaginator struct {
	options   ListJobExecutionsForJobPaginatorOptions
	client    ListJobExecutionsForJobAPIClient
	params    *ListJobExecutionsForJobInput
	nextToken *string
	firstPage bool
}

// NewListJobExecutionsForJobPaginator returns a new
// ListJobExecutionsForJobPaginator
func NewListJobExecutionsForJobPaginator(client ListJobExecutionsForJobAPIClient, params *ListJobExecutionsForJobInput, optFns ...func(*ListJobExecutionsForJobPaginatorOptions)) *ListJobExecutionsForJobPaginator {
	options := ListJobExecutionsForJobPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListJobExecutionsForJobInput{}
	}

	return &ListJobExecutionsForJobPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobExecutionsForJobPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJobExecutionsForJob page.
func (p *ListJobExecutionsForJobPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobExecutionsForJobOutput, error) {
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

	result, err := p.client.ListJobExecutionsForJob(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobExecutionsForJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListJobExecutionsForJob",
	}
}

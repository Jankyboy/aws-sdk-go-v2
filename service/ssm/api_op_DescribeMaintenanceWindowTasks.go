// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tasks in a maintenance window.
func (c *Client) DescribeMaintenanceWindowTasks(ctx context.Context, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowTasks", params, optFns, addOperationDescribeMaintenanceWindowTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowTasksInput struct {

	// The ID of the maintenance window whose tasks should be retrieved.
	//
	// This member is required.
	WindowId *string

	// Optional filters used to narrow down the scope of the returned tasks. The
	// supported filter keys are WindowTaskId, TaskArn, Priority, and TaskType.
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeMaintenanceWindowTasksOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the tasks in the maintenance window.
	Tasks []types.MaintenanceWindowTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMaintenanceWindowTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowTasksAPIClient is a client that implements the
// DescribeMaintenanceWindowTasks operation.
type DescribeMaintenanceWindowTasksAPIClient interface {
	DescribeMaintenanceWindowTasks(context.Context, *DescribeMaintenanceWindowTasksInput, ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error)
}

var _ DescribeMaintenanceWindowTasksAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowTasksPaginatorOptions is the paginator options for
// DescribeMaintenanceWindowTasks
type DescribeMaintenanceWindowTasksPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowTasksPaginator is a paginator for
// DescribeMaintenanceWindowTasks
type DescribeMaintenanceWindowTasksPaginator struct {
	options   DescribeMaintenanceWindowTasksPaginatorOptions
	client    DescribeMaintenanceWindowTasksAPIClient
	params    *DescribeMaintenanceWindowTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowTasksPaginator returns a new
// DescribeMaintenanceWindowTasksPaginator
func NewDescribeMaintenanceWindowTasksPaginator(client DescribeMaintenanceWindowTasksAPIClient, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*DescribeMaintenanceWindowTasksPaginatorOptions)) *DescribeMaintenanceWindowTasksPaginator {
	options := DescribeMaintenanceWindowTasksPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeMaintenanceWindowTasksInput{}
	}

	return &DescribeMaintenanceWindowTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowTasksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeMaintenanceWindowTasks page.
func (p *DescribeMaintenanceWindowTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeMaintenanceWindowTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowTasks",
	}
}

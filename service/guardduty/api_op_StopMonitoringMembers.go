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

// Stops GuardDuty monitoring for the specified member accounts. Use the
// StartMonitoringMembers operation to restart monitoring for those accounts.
func (c *Client) StopMonitoringMembers(ctx context.Context, params *StopMonitoringMembersInput, optFns ...func(*Options)) (*StopMonitoringMembersOutput, error) {
	if params == nil {
		params = &StopMonitoringMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopMonitoringMembers", params, optFns, addOperationStopMonitoringMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopMonitoringMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopMonitoringMembersInput struct {

	// A list of account IDs for the member accounts to stop monitoring.
	//
	// This member is required.
	AccountIds []string

	// The unique ID of the detector associated with the GuardDuty master account that
	// is monitoring member accounts.
	//
	// This member is required.
	DetectorId *string
}

type StopMonitoringMembersOutput struct {

	// A list of objects that contain an accountId for each account that could not be
	// processed, and a result string that indicates why the account was not processed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopMonitoringMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopMonitoringMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMonitoringMembers{}, middleware.After)
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
	if err = addOpStopMonitoringMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopMonitoringMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopMonitoringMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "StopMonitoringMembers",
	}
}

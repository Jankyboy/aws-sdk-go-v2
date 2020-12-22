// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The SendBonus operation issues a payment of money from your account to a Worker.
// This payment happens separately from the reward you pay to the Worker when you
// approve the Worker's assignment. The SendBonus operation requires the Worker's
// ID and the assignment ID as parameters to initiate payment of the bonus. You
// must include a message that explains the reason for the bonus payment, as the
// Worker may not be expecting the payment. Amazon Mechanical Turk collects a fee
// for bonus payments, similar to the HIT listing fee. This operation fails if your
// account does not have enough funds to pay for both the bonus and the fees.
func (c *Client) SendBonus(ctx context.Context, params *SendBonusInput, optFns ...func(*Options)) (*SendBonusOutput, error) {
	if params == nil {
		params = &SendBonusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendBonus", params, optFns, addOperationSendBonusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendBonusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendBonusInput struct {

	// The ID of the assignment for which this bonus is paid.
	//
	// This member is required.
	AssignmentId *string

	// The Bonus amount is a US Dollar amount specified using a string (for example,
	// "5" represents $5.00 USD and "101.42" represents $101.42 USD). Do not include
	// currency symbols or currency codes.
	//
	// This member is required.
	BonusAmount *string

	// A message that explains the reason for the bonus payment. The Worker receiving
	// the bonus can see this message.
	//
	// This member is required.
	Reason *string

	// The ID of the Worker being paid the bonus.
	//
	// This member is required.
	WorkerId *string

	// A unique identifier for this request, which allows you to retry the call on
	// error without granting multiple bonuses. This is useful in cases such as network
	// timeouts where it is unclear whether or not the call succeeded on the server. If
	// the bonus already exists in the system from a previous call using the same
	// UniqueRequestToken, subsequent calls will return an error with a message
	// containing the request ID.
	UniqueRequestToken *string
}

type SendBonusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendBonusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendBonus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendBonus{}, middleware.After)
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
	if err = addOpSendBonusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendBonus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendBonus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "SendBonus",
	}
}

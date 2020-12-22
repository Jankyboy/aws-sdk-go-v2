// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a device certificate with AWS IoT. If you have more than one CA
// certificate that has the same subject field, you must specify the CA certificate
// that was used to sign the device certificate being registered.
func (c *Client) RegisterCertificate(ctx context.Context, params *RegisterCertificateInput, optFns ...func(*Options)) (*RegisterCertificateOutput, error) {
	if params == nil {
		params = &RegisterCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCertificate", params, optFns, addOperationRegisterCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the RegisterCertificate operation.
type RegisterCertificateInput struct {

	// The certificate data, in PEM format.
	//
	// This member is required.
	CertificatePem *string

	// The CA certificate used to sign the device certificate being registered.
	CaCertificatePem *string

	// A boolean value that specifies if the certificate is set to active.
	//
	// Deprecated: This member has been deprecated.
	SetAsActive *bool

	// The status of the register certificate request.
	Status types.CertificateStatus
}

// The output from the RegisterCertificate operation.
type RegisterCertificateOutput struct {

	// The certificate ARN.
	CertificateArn *string

	// The certificate identifier.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCertificate{}, middleware.After)
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
	if err = addOpRegisterCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RegisterCertificate",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified SSH public key, including metadata about the key. The
// SSH public key retrieved by this operation is used only for authenticating the
// associated IAM user to an AWS CodeCommit repository. For more information about
// using SSH keys to authenticate to an AWS CodeCommit repository, see Set up AWS
// CodeCommit for SSH Connections
// (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
func (c *Client) GetSSHPublicKey(ctx context.Context, params *GetSSHPublicKeyInput, optFns ...func(*Options)) (*GetSSHPublicKeyOutput, error) {
	if params == nil {
		params = &GetSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSSHPublicKey", params, optFns, addOperationGetSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSSHPublicKeyInput struct {

	// Specifies the public key encoding format to use in the response. To retrieve the
	// public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format,
	// use PEM.
	//
	// This member is required.
	Encoding types.EncodingType

	// The unique identifier for the SSH public key. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters that can
	// consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	SSHPublicKeyId *string

	// The name of the IAM user associated with the SSH public key. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

// Contains the response to a successful GetSSHPublicKey request.
type GetSSHPublicKeyOutput struct {

	// A structure containing details about the SSH public key.
	SSHPublicKey *types.SSHPublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSSHPublicKey{}, middleware.After)
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
	if err = addOpGetSSHPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSSHPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSSHPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetSSHPublicKey",
	}
}

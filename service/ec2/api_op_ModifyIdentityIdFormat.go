// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the ID format of a resource for a specified IAM user, IAM role, or the
// root user for an account; or all IAM users, IAM roles, and the root user for an
// account. You can specify that resources should receive longer IDs (17-character
// IDs) when they are created. This request can only be used to modify longer ID
// settings for resource types that are within the opt-in period. Resources
// currently in their opt-in period include: bundle | conversion-task |
// customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | internet-gateway | network-acl
// | network-acl-association | network-interface | network-interface-attachment |
// prefix-list | route-table | route-table-association | security-group | subnet |
// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
// | vpc-peering-connection | vpn-connection | vpn-gateway. For more information,
// see Resource IDs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
// Amazon Elastic Compute Cloud User Guide. This setting applies to the principal
// specified in the request; it does not apply to the principal that makes the
// request. Resources created with longer IDs are visible to all IAM roles and
// users, regardless of these settings and provided that they have permission to
// use the relevant Describe command for the resource type.
func (c *Client) ModifyIdentityIdFormat(ctx context.Context, params *ModifyIdentityIdFormatInput, optFns ...func(*Options)) (*ModifyIdentityIdFormatOutput, error) {
	if params == nil {
		params = &ModifyIdentityIdFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyIdentityIdFormat", params, optFns, addOperationModifyIdentityIdFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyIdentityIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyIdentityIdFormatInput struct {

	// The ARN of the principal, which can be an IAM user, IAM role, or the root user.
	// Specify all to modify the ID format for all IAM users, IAM roles, and the root
	// user of the account.
	//
	// This member is required.
	PrincipalArn *string

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log |
	// image | import-task | internet-gateway | network-acl | network-acl-association |
	// network-interface | network-interface-attachment | prefix-list | route-table |
	// route-table-association | security-group | subnet |
	// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
	// | vpc-peering-connection | vpn-connection | vpn-gateway. Alternatively, use the
	// all-current option to include all resource types that are currently within their
	// opt-in period for longer IDs.
	//
	// This member is required.
	Resource *string

	// Indicates whether the resource should use longer IDs (17-character IDs)
	//
	// This member is required.
	UseLongIds bool
}

type ModifyIdentityIdFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyIdentityIdFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyIdentityIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyIdentityIdFormat{}, middleware.After)
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
	if err = addOpModifyIdentityIdFormatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyIdentityIdFormat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyIdentityIdFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyIdentityIdFormat",
	}
}

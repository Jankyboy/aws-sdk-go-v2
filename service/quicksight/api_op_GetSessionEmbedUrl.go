// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a session URL and authorization code that you can use to embed the
// Amazon QuickSight console in your web server code. Use GetSessionEmbedUrl where
// you want to provide an authoring portal that allows users to create data
// sources, datasets, analyses, and dashboards. The users who access an embedded
// QuickSight console need belong to the author or admin security cohort. If you
// want to restrict permissions to some of these features, add a custom permissions
// profile to the user with the UpdateUser API operation. Use RegisterUser API
// operation to add a new user with a custom permission profile attached. For more
// information, see the following sections in the Amazon QuickSight User Guide:
//
// *
// Embedding the Amazon QuickSight Console
// (https://docs.aws.amazon.com/quicksight/latest/user/embedding-the-quicksight-console.html)
//
// *
// Customizing Access to the Amazon QuickSight Console
// (https://docs.aws.amazon.com/quicksight/latest/user/customizing-permissions-to-the-quicksight-console.html)
func (c *Client) GetSessionEmbedUrl(ctx context.Context, params *GetSessionEmbedUrlInput, optFns ...func(*Options)) (*GetSessionEmbedUrlOutput, error) {
	if params == nil {
		params = &GetSessionEmbedUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSessionEmbedUrl", params, optFns, addOperationGetSessionEmbedUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSessionEmbedUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionEmbedUrlInput struct {

	// The ID for the AWS account associated with your QuickSight subscription.
	//
	// This member is required.
	AwsAccountId *string

	// The URL you use to access the embedded session. The entry point URL is
	// constrained to the following paths:
	//
	// * /start
	//
	// * /start/analyses
	//
	// *
	// /start/dashboards
	//
	// * /start/favorites
	//
	// * /dashboards/DashboardId  - where
	// DashboardId is the actual ID key from the QuickSight console URL of the
	// dashboard
	//
	// * /analyses/AnalysisId  - where AnalysisId is the actual ID key from
	// the QuickSight console URL of the analysis
	EntryPoint *string

	// How many minutes the session is valid. The session lifetime must be 15-600
	// minutes.
	SessionLifetimeInMinutes *int64

	// The Amazon QuickSight user's Amazon Resource Name (ARN), for use with QUICKSIGHT
	// identity type. You can use this for any type of Amazon QuickSight users in your
	// account (readers, authors, or admins). They need to be authenticated as one of
	// the following:
	//
	// * Active Directory (AD) users or group members
	//
	// * Invited
	// nonfederated users
	//
	// * IAM users and IAM role-based sessions authenticated
	// through Federated Single Sign-On using SAML, OpenID Connect, or IAM
	// federation
	//
	// Omit this parameter for users in the third group – IAM users and IAM
	// role-based sessions.
	UserArn *string
}

type GetSessionEmbedUrlOutput struct {

	// A single-use URL that you can put into your server-side web page to embed your
	// QuickSight session. This URL is valid for 5 minutes. The API operation provides
	// the URL with an auth_code value that enables one (and only one) sign-on to a
	// user session that is valid for 10 hours.
	EmbedUrl *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSessionEmbedUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSessionEmbedUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSessionEmbedUrl{}, middleware.After)
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
	if err = addOpGetSessionEmbedUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSessionEmbedUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSessionEmbedUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GetSessionEmbedUrl",
	}
}

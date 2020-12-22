// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new database in Amazon Lightsail. The create relational database
// operation supports tag-based access control via request tags. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateRelationalDatabase(ctx context.Context, params *CreateRelationalDatabaseInput, optFns ...func(*Options)) (*CreateRelationalDatabaseOutput, error) {
	if params == nil {
		params = &CreateRelationalDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRelationalDatabase", params, optFns, addOperationCreateRelationalDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRelationalDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRelationalDatabaseInput struct {

	// The name of the master database created when the Lightsail database resource is
	// created. Constraints:
	//
	// * Must contain from 1 to 64 alphanumeric characters.
	//
	// *
	// Cannot be a word reserved by the specified database engine
	//
	// This member is required.
	MasterDatabaseName *string

	// The master user name for your new database. Constraints:
	//
	// * Master user name is
	// required.
	//
	// * Must contain from 1 to 16 alphanumeric characters.
	//
	// * The first
	// character must be a letter.
	//
	// * Cannot be a reserved word for the database engine
	// you choose. For more information about reserved words in MySQL 5.6 or 5.7, see
	// the Keywords and Reserved Words articles for MySQL 5.6
	// (https://dev.mysql.com/doc/refman/5.6/en/keywords.html) or MySQL 5.7
	// (https://dev.mysql.com/doc/refman/5.7/en/keywords.html) respectively.
	//
	// This member is required.
	MasterUsername *string

	// The blueprint ID for your new database. A blueprint describes the major engine
	// version of a database. You can get a list of database blueprints IDs by using
	// the get relational database blueprints operation.
	//
	// This member is required.
	RelationalDatabaseBlueprintId *string

	// The bundle ID for your new database. A bundle describes the performance
	// specifications for your database. You can get a list of database bundle IDs by
	// using the get relational database bundles operation.
	//
	// This member is required.
	RelationalDatabaseBundleId *string

	// The name to use for your new database. Constraints:
	//
	// * Must contain from 2 to
	// 255 alphanumeric characters, or hyphens.
	//
	// * The first and last character must be
	// a letter or number.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The Availability Zone in which to create your new database. Use the us-east-2a
	// case-sensitive format. You can get a list of Availability Zones by using the get
	// regions operation. Be sure to add the include relational database Availability
	// Zones parameter to your request.
	AvailabilityZone *string

	// The password for the master user of your new database. The password can include
	// any printable ASCII character except "/", """, or "@". Constraints: Must contain
	// 8 to 41 characters.
	MasterUserPassword *string

	// The daily time range during which automated backups are created for your new
	// database if automated backups are enabled. The default is a 30-minute window
	// selected at random from an 8-hour block of time for each AWS Region. For more
	// information about the preferred backup window time blocks for each region, see
	// the Working With Backups
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)
	// guide in the Amazon Relational Database Service (Amazon RDS) documentation.
	// Constraints:
	//
	// * Must be in the hh24:mi-hh24:mi format. Example: 16:00-16:30
	//
	// *
	// Specified in Coordinated Universal Time (UTC).
	//
	// * Must not conflict with the
	// preferred maintenance window.
	//
	// * Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur on your new
	// database. The default is a 30-minute window selected at random from an 8-hour
	// block of time for each AWS Region, occurring on a random day of the week.
	// Constraints:
	//
	// * Must be in the ddd:hh24:mi-ddd:hh24:mi format.
	//
	// * Valid days:
	// Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// * Must be at least 30 minutes.
	//
	// * Specified
	// in Coordinated Universal Time (UTC).
	//
	// * Example: Tue:17:00-Tue:17:30
	PreferredMaintenanceWindow *string

	// Specifies the accessibility options for your new database. A value of true
	// specifies a database that is available to resources outside of your Lightsail
	// account. A value of false specifies a database that is available only to your
	// Lightsail resources in the same region as your database.
	PubliclyAccessible *bool

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []types.Tag
}

type CreateRelationalDatabaseOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRelationalDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRelationalDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRelationalDatabase{}, middleware.After)
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
	if err = addOpCreateRelationalDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRelationalDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRelationalDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateRelationalDatabase",
	}
}

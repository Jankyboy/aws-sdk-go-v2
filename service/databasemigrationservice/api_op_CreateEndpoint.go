// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint using the provided settings.
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateEndpointInput struct {

	// The database endpoint identifier. Identifiers must begin with a letter and must
	// contain only ASCII letters, digits, and hyphens. They can't end with a hyphen,
	// or contain two consecutive hyphens.
	//
	// This member is required.
	EndpointIdentifier *string

	// The type of endpoint. Valid values are source and target.
	//
	// This member is required.
	EndpointType types.ReplicationEndpointTypeValue

	// The type of engine for the endpoint. Valid values, depending on the EndpointType
	// value, include "mysql", "oracle", "postgres", "mariadb", "aurora",
	// "aurora-postgresql", "redshift", "s3", "db2", "azuredb", "sybase", "dynamodb",
	// "mongodb", "kinesis", "kafka", "elasticsearch", "docdb", "sqlserver", and
	// "neptune".
	//
	// This member is required.
	EngineName *string

	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string

	// The name of the endpoint database.
	DatabaseName *string

	// The settings in JSON format for the DMS transfer type of source endpoint.
	// Possible settings include the following:
	//
	// * ServiceAccessRoleArn - The IAM role
	// that has permission to access the Amazon S3 bucket.
	//
	// * BucketName - The name of
	// the S3 bucket to use.
	//
	// * CompressionType - An optional parameter to use GZIP to
	// compress the target files. To use GZIP, set this value to NONE (the default). To
	// keep the files uncompressed, don't use this value.
	//
	// Shorthand syntax for these
	// settings is as follows:
	// ServiceAccessRoleArn=string,BucketName=string,CompressionType=string JSON syntax
	// for these settings is as follows: { "ServiceAccessRoleArn": "string",
	// "BucketName": "string", "CompressionType": "none"|"gzip" }
	DmsTransferSettings *types.DmsTransferSettings

	// Provides information that defines a DocumentDB endpoint.
	DocDbSettings *types.DocDbSettings

	// Settings in JSON format for the target Amazon DynamoDB endpoint. For information
	// about other available settings, see Using Object Mapping to Migrate Data to
	// DynamoDB
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html) in
	// the AWS Database Migration Service User Guide.
	DynamoDbSettings *types.DynamoDbSettings

	// Settings in JSON format for the target Elasticsearch endpoint. For more
	// information about the available settings, see Extra Connection Attributes When
	// Using Elasticsearch as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration)
	// in the AWS Database Migration Service User Guide.
	ElasticsearchSettings *types.ElasticsearchSettings

	// The external table definition.
	ExternalTableDefinition *string

	// Additional attributes associated with the connection. Each attribute is
	// specified as a name-value pair associated by an equal sign (=). Multiple
	// attributes are separated by a semicolon (;) with no additional white space. For
	// information on the attributes available for connecting your source or target
	// endpoint, see Working with AWS DMS Endpoints
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the
	// AWS Database Migration Service User Guide.
	ExtraConnectionAttributes *string

	// Settings in JSON format for the source IBM Db2 LUW endpoint. For information
	// about other available settings, see Extra connection attributes when using Db2
	// LUW as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html) in the
	// AWS Database Migration Service User Guide.
	IBMDb2Settings *types.IBMDb2Settings

	// Settings in JSON format for the target Apache Kafka endpoint. For more
	// information about the available settings, see Using Apache Kafka as a Target for
	// AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html) in the
	// AWS Database Migration Service User Guide.
	KafkaSettings *types.KafkaSettings

	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	// For more information about the available settings, see Using Amazon Kinesis Data
	// Streams as a Target for AWS Database Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html) in
	// the AWS Database Migration Service User Guide.
	KinesisSettings *types.KinesisSettings

	// An AWS KMS key identifier that is used to encrypt the connection parameters for
	// the endpoint. If you don't specify a value for the KmsKeyId parameter, then AWS
	// DMS uses your default encryption key. AWS KMS creates the default encryption key
	// for your AWS account. Your AWS account has a different default encryption key
	// for each AWS Region.
	KmsKeyId *string

	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	// For information about other available settings, see Extra connection attributes
	// when using SQL Server as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html)
	// and  Extra connection attributes when using SQL Server as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html) in
	// the AWS Database Migration Service User Guide.
	MicrosoftSQLServerSettings *types.MicrosoftSQLServerSettings

	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see Using MongoDB as a Target for AWS Database
	// Migration Service
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration)
	// in the AWS Database Migration Service User Guide.
	MongoDbSettings *types.MongoDbSettings

	// Settings in JSON format for the source and target MySQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using MySQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html) and
	// Extra connection attributes when using a MySQL-compatible database as a target
	// for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html) in the
	// AWS Database Migration Service User Guide.
	MySQLSettings *types.MySQLSettings

	// Settings in JSON format for the target Amazon Neptune endpoint. For more
	// information about the available settings, see Specifying Endpoint Settings for
	// Amazon Neptune as a Target
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings)
	// in the AWS Database Migration Service User Guide.
	NeptuneSettings *types.NeptuneSettings

	// Settings in JSON format for the source and target Oracle endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using Oracle as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html) and
	// Extra connection attributes when using Oracle as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html) in
	// the AWS Database Migration Service User Guide.
	OracleSettings *types.OracleSettings

	// The password to be used to log in to the endpoint database.
	Password *string

	// The port used by the endpoint database.
	Port *int32

	// Settings in JSON format for the source and target PostgreSQL endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using PostgreSQL as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html)
	// and  Extra connection attributes when using PostgreSQL as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html)
	// in the AWS Database Migration Service User Guide.
	PostgreSQLSettings *types.PostgreSQLSettings

	// Provides information that defines an Amazon Redshift endpoint.
	RedshiftSettings *types.RedshiftSettings

	// A friendly name for the resource identifier at the end of the EndpointArn
	// response parameter that is returned in the created Endpoint object. The value
	// for this parameter can have up to 31 characters. It can contain only ASCII
	// letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain
	// two consecutive hyphens, and can only begin with a letter, such as
	// Example-App-ARN1. For example, this value might result in the EndpointArn value
	// arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1. If you don't specify a
	// ResourceIdentifier value, AWS DMS generates a default identifier value for the
	// end of EndpointArn.
	ResourceIdentifier *string

	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using Amazon
	// S3 as a Target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring)
	// in the AWS Database Migration Service User Guide.
	S3Settings *types.S3Settings

	// The name of the server where the endpoint database resides.
	ServerName *string

	// The Amazon Resource Name (ARN) for the service access role that you want to use
	// to create the endpoint.
	ServiceAccessRoleArn *string

	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default
	// is none
	SslMode types.DmsSslModeValue

	// Settings in JSON format for the source and target SAP ASE endpoint. For
	// information about other available settings, see Extra connection attributes when
	// using SAP ASE as a source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html) and
	// Extra connection attributes when using SAP ASE as a target for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html) in the
	// AWS Database Migration Service User Guide.
	SybaseSettings *types.SybaseSettings

	// One or more tags to be assigned to the endpoint.
	Tags []types.Tag

	// The user name to be used to log in to the endpoint database.
	Username *string
}

//
type CreateEndpointOutput struct {

	// The endpoint that was created.
	Endpoint *types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
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
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateEndpoint",
	}
}

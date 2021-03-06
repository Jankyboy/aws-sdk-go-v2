// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computeoptimizer

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes the configuration of an Auto Scaling group.
type AutoScalingGroupConfiguration struct {
	_ struct{} `type:"structure"`

	// The desired capacity, or number of instances, for the Auto Scaling group.
	DesiredCapacity *int64 `locationName:"desiredCapacity" type:"integer"`

	// The instance type for the Auto Scaling group.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// The maximum size, or maximum number of instances, for the Auto Scaling group.
	MaxSize *int64 `locationName:"maxSize" type:"integer"`

	// The minimum size, or minimum number of instances, for the Auto Scaling group.
	MinSize *int64 `locationName:"minSize" type:"integer"`
}

// String returns the string representation
func (s AutoScalingGroupConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Describes an Auto Scaling group recommendation.
type AutoScalingGroupRecommendation struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the Auto Scaling group.
	AccountId *string `locationName:"accountId" type:"string"`

	// The Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoScalingGroupArn *string `locationName:"autoScalingGroupArn" type:"string"`

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `locationName:"autoScalingGroupName" type:"string"`

	// An array of objects that describe the current configuration of the Auto Scaling
	// group.
	CurrentConfiguration *AutoScalingGroupConfiguration `locationName:"currentConfiguration" type:"structure"`

	// The finding classification for the Auto Scaling group.
	//
	// Findings for Auto Scaling groups include:
	//
	//    * NotOptimized —An Auto Scaling group is considered not optimized when
	//    AWS Compute Optimizer identifies a recommendation that can provide better
	//    performance for your workload.
	//
	//    * Optimized —An Auto Scaling group is considered optimized when Compute
	//    Optimizer determines that the group is correctly provisioned to run your
	//    workload based on the chosen instance type. For optimized resources, Compute
	//    Optimizer might recommend a new generation instance type.
	//
	// The values that are returned might be NOT_OPTIMIZED or OPTIMIZED.
	Finding Finding `locationName:"finding" type:"string" enum:"true"`

	// The time stamp of when the Auto Scaling group recommendation was last refreshed.
	LastRefreshTimestamp *time.Time `locationName:"lastRefreshTimestamp" type:"timestamp"`

	// The number of days for which utilization metrics were analyzed for the Auto
	// Scaling group.
	LookBackPeriodInDays *float64 `locationName:"lookBackPeriodInDays" type:"double"`

	// An array of objects that describe the recommendation options for the Auto
	// Scaling group.
	RecommendationOptions []AutoScalingGroupRecommendationOption `locationName:"recommendationOptions" type:"list"`

	// An array of objects that describe the utilization metrics of the Auto Scaling
	// group.
	UtilizationMetrics []UtilizationMetric `locationName:"utilizationMetrics" type:"list"`
}

// String returns the string representation
func (s AutoScalingGroupRecommendation) String() string {
	return awsutil.Prettify(s)
}

// Describes a recommendation option for an Auto Scaling group.
type AutoScalingGroupRecommendationOption struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe an Auto Scaling group configuration.
	Configuration *AutoScalingGroupConfiguration `locationName:"configuration" type:"structure"`

	// The performance risk of the Auto Scaling group configuration recommendation.
	//
	// Performance risk is the likelihood of the recommended instance type not meeting
	// the performance requirement of your workload.
	//
	// The lowest performance risk is categorized as 0, and the highest as 5.
	PerformanceRisk *float64 `locationName:"performanceRisk" type:"double"`

	// An array of objects that describe the projected utilization metrics of the
	// Auto Scaling group recommendation option.
	ProjectedUtilizationMetrics []UtilizationMetric `locationName:"projectedUtilizationMetrics" type:"list"`

	// The rank of the Auto Scaling group recommendation option.
	//
	// The top recommendation option is ranked as 1.
	Rank *int64 `locationName:"rank" type:"integer"`
}

// String returns the string representation
func (s AutoScalingGroupRecommendationOption) String() string {
	return awsutil.Prettify(s)
}

// Describes the destination of the recommendations export and metadata files.
type ExportDestination struct {
	_ struct{} `type:"structure"`

	// An object that describes the destination Amazon Simple Storage Service (Amazon
	// S3) bucket name and object keys of a recommendations export file, and its
	// associated metadata file.
	S3 *S3Destination `locationName:"s3" type:"structure"`
}

// String returns the string representation
func (s ExportDestination) String() string {
	return awsutil.Prettify(s)
}

// Describes a filter that returns a more specific list of recommendations.
type Filter struct {
	_ struct{} `type:"structure"`

	// The name of the filter.
	//
	// Specify Finding to return recommendations with a specific findings classification
	// (e.g., Overprovisioned).
	//
	// Specify RecommendationSourceType to return recommendations of a specific
	// resource type (e.g., AutoScalingGroup).
	Name FilterName `locationName:"name" type:"string" enum:"true"`

	// The value of the filter.
	//
	// If you specify the name parameter as Finding, and you request recommendations
	// for an instance, then the valid values are Underprovisioned, Overprovisioned,
	// NotOptimized, or Optimized.
	//
	// If you specify the name parameter as Finding, and you request recommendations
	// for an Auto Scaling group, then the valid values are Optimized, or NotOptimized.
	//
	// If you specify the name parameter as RecommendationSourceType, then the valid
	// values are Ec2Instance, or AutoScalingGroup.
	Values []string `locationName:"values" type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Describes an error experienced when getting recommendations.
//
// For example, an error is returned if you request recommendations for an unsupported
// Auto Scaling group, or if you request recommendations for an instance of
// an unsupported instance family.
type GetRecommendationError struct {
	_ struct{} `type:"structure"`

	// The error code.
	Code *string `locationName:"code" type:"string"`

	// The ID of the error.
	Identifier *string `locationName:"identifier" type:"string"`

	// The message, or reason, for the error.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s GetRecommendationError) String() string {
	return awsutil.Prettify(s)
}

// Describes an Amazon EC2 instance recommendation.
type InstanceRecommendation struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the instance.
	AccountId *string `locationName:"accountId" type:"string"`

	// The instance type of the current instance.
	CurrentInstanceType *string `locationName:"currentInstanceType" type:"string"`

	// The finding classification for the instance.
	//
	// Findings for instances include:
	//
	//    * Underprovisioned —An instance is considered under-provisioned when
	//    at least one specification of your instance, such as CPU, memory, or network,
	//    does not meet the performance requirements of your workload. Under-provisioned
	//    instances may lead to poor application performance.
	//
	//    * Overprovisioned —An instance is considered over-provisioned when at
	//    least one specification of your instance, such as CPU, memory, or network,
	//    can be sized down while still meeting the performance requirements of
	//    your workload, and no specification is under-provisioned. Over-provisioned
	//    instances may lead to unnecessary infrastructure cost.
	//
	//    * Optimized —An instance is considered optimized when all specifications
	//    of your instance, such as CPU, memory, and network, meet the performance
	//    requirements of your workload and is not over provisioned. An optimized
	//    instance runs your workloads with optimal performance and infrastructure
	//    cost. For optimized resources, AWS Compute Optimizer might recommend a
	//    new generation instance type.
	//
	// The values that are returned might be UNDER_PROVISIONED, OVER_PROVISIONED,
	// or OPTIMIZED.
	Finding Finding `locationName:"finding" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the current instance.
	InstanceArn *string `locationName:"instanceArn" type:"string"`

	// The name of the current instance.
	InstanceName *string `locationName:"instanceName" type:"string"`

	// The time stamp of when the instance recommendation was last refreshed.
	LastRefreshTimestamp *time.Time `locationName:"lastRefreshTimestamp" type:"timestamp"`

	// The number of days for which utilization metrics were analyzed for the instance.
	LookBackPeriodInDays *float64 `locationName:"lookBackPeriodInDays" type:"double"`

	// An array of objects that describe the recommendation options for the instance.
	RecommendationOptions []InstanceRecommendationOption `locationName:"recommendationOptions" type:"list"`

	// An array of objects that describe the source resource of the recommendation.
	RecommendationSources []RecommendationSource `locationName:"recommendationSources" type:"list"`

	// An array of objects that describe the utilization metrics of the instance.
	UtilizationMetrics []UtilizationMetric `locationName:"utilizationMetrics" type:"list"`
}

// String returns the string representation
func (s InstanceRecommendation) String() string {
	return awsutil.Prettify(s)
}

// Describes a recommendation option for an Amazon EC2 instance.
type InstanceRecommendationOption struct {
	_ struct{} `type:"structure"`

	// The instance type of the instance recommendation.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// The performance risk of the instance recommendation option.
	//
	// Performance risk is the likelihood of the recommended instance type not meeting
	// the performance requirement of your workload.
	//
	// The lowest performance risk is categorized as 0, and the highest as 5.
	PerformanceRisk *float64 `locationName:"performanceRisk" type:"double"`

	// An array of objects that describe the projected utilization metrics of the
	// instance recommendation option.
	ProjectedUtilizationMetrics []UtilizationMetric `locationName:"projectedUtilizationMetrics" type:"list"`

	// The rank of the instance recommendation option.
	//
	// The top recommendation option is ranked as 1.
	Rank *int64 `locationName:"rank" type:"integer"`
}

// String returns the string representation
func (s InstanceRecommendationOption) String() string {
	return awsutil.Prettify(s)
}

// Describes a filter that returns a more specific list of recommendation export
// jobs.
//
// This filter is used with the DescribeRecommendationExportJobs action.
type JobFilter struct {
	_ struct{} `type:"structure"`

	// The name of the filter.
	//
	// Specify ResourceType to return export jobs of a specific resource type (e.g.,
	// Ec2Instance).
	//
	// Specify JobStatus to return export jobs with a specific status (e.g, Complete).
	Name JobFilterName `locationName:"name" type:"string" enum:"true"`

	// The value of the filter.
	//
	// If you specify the name parameter as ResourceType, the valid values are Ec2Instance
	// or AutoScalingGroup.
	//
	// If you specify the name parameter as JobStatus, the valid values are Queued,
	// InProgress, Complete, or Failed.
	Values []string `locationName:"values" type:"list"`
}

// String returns the string representation
func (s JobFilter) String() string {
	return awsutil.Prettify(s)
}

// Describes a projected utilization metric of a recommendation option, such
// as an Amazon EC2 instance.
type ProjectedMetric struct {
	_ struct{} `type:"structure"`

	// The name of the projected utilization metric.
	//
	// Memory metrics are only returned for resources that have the unified CloudWatch
	// agent installed on them. For more information, see Enabling Memory Utilization
	// with the CloudWatch Agent (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Install-CloudWatch-Agent.html).
	Name MetricName `locationName:"name" type:"string" enum:"true"`

	// The time stamps of the projected utilization metric.
	Timestamps []time.Time `locationName:"timestamps" type:"list"`

	// The values of the projected utilization metrics.
	Values []float64 `locationName:"values" type:"list"`
}

// String returns the string representation
func (s ProjectedMetric) String() string {
	return awsutil.Prettify(s)
}

// Describes a recommendation export job.
//
// Use the DescribeRecommendationExportJobs action to view your recommendation
// export jobs.
//
// Use the ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations
// actions to request an export of your recommendations.
type RecommendationExportJob struct {
	_ struct{} `type:"structure"`

	// The timestamp of when the export job was created.
	CreationTimestamp *time.Time `locationName:"creationTimestamp" type:"timestamp"`

	// An object that describes the destination of the export file.
	Destination *ExportDestination `locationName:"destination" type:"structure"`

	// The reason for an export job failure.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The identification number of the export job.
	JobId *string `locationName:"jobId" type:"string"`

	// The timestamp of when the export job was last updated.
	LastUpdatedTimestamp *time.Time `locationName:"lastUpdatedTimestamp" type:"timestamp"`

	// The resource type of the exported recommendations.
	ResourceType ResourceType `locationName:"resourceType" type:"string" enum:"true"`

	// The status of the export job.
	Status JobStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s RecommendationExportJob) String() string {
	return awsutil.Prettify(s)
}

// Describes the source of a recommendation, such as an Amazon EC2 instance
// or Auto Scaling group.
type RecommendationSource struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the recommendation source.
	RecommendationSourceArn *string `locationName:"recommendationSourceArn" type:"string"`

	// The resource type of the recommendation source.
	RecommendationSourceType RecommendationSourceType `locationName:"recommendationSourceType" type:"string" enum:"true"`
}

// String returns the string representation
func (s RecommendationSource) String() string {
	return awsutil.Prettify(s)
}

// A summary of a recommendation.
type RecommendationSummary struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the recommendation summary.
	AccountId *string `locationName:"accountId" type:"string"`

	// The resource type of the recommendation.
	RecommendationResourceType RecommendationSourceType `locationName:"recommendationResourceType" type:"string" enum:"true"`

	// An array of objects that describe a recommendation summary.
	Summaries []Summary `locationName:"summaries" type:"list"`
}

// String returns the string representation
func (s RecommendationSummary) String() string {
	return awsutil.Prettify(s)
}

// Describes a projected utilization metric of a recommendation option.
type RecommendedOptionProjectedMetric struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe a projected utilization metric.
	ProjectedMetrics []ProjectedMetric `locationName:"projectedMetrics" type:"list"`

	// The rank of the recommendation option projected metric.
	//
	// The top recommendation option is ranked as 1.
	//
	// The projected metric rank correlates to the recommendation option rank. For
	// example, the projected metric ranked as 1 is related to the recommendation
	// option that is also ranked as 1 in the same response.
	Rank *int64 `locationName:"rank" type:"integer"`

	// The recommended instance type.
	RecommendedInstanceType *string `locationName:"recommendedInstanceType" type:"string"`
}

// String returns the string representation
func (s RecommendedOptionProjectedMetric) String() string {
	return awsutil.Prettify(s)
}

// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket
// name and object keys of a recommendations export file, and its associated
// metadata file.
type S3Destination struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon S3 bucket used as the destination of an export file.
	Bucket *string `locationName:"bucket" type:"string"`

	// The Amazon S3 bucket key of an export file.
	//
	// The key uniquely identifies the object, or export file, in the S3 bucket.
	Key *string `locationName:"key" type:"string"`

	// The Amazon S3 bucket key of a metadata file.
	//
	// The key uniquely identifies the object, or metadata file, in the S3 bucket.
	MetadataKey *string `locationName:"metadataKey" type:"string"`
}

// String returns the string representation
func (s S3Destination) String() string {
	return awsutil.Prettify(s)
}

// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket
// name and key prefix for a recommendations export job.
//
// You must create the destination Amazon S3 bucket for your recommendations
// export before you create the export job. Compute Optimizer does not create
// the S3 bucket for you. After you create the S3 bucket, ensure that it has
// the required permission policy to allow Compute Optimizer to write the export
// file to it. If you plan to specify an object prefix when you create the export
// job, you must include the object prefix in the policy that you add to the
// S3 bucket. For more information, see Amazon S3 Bucket Policy for Compute
// Optimizer (https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html)
// in the Compute Optimizer user guide.
type S3DestinationConfig struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon S3 bucket to use as the destination for an export
	// job.
	Bucket *string `locationName:"bucket" type:"string"`

	// The Amazon S3 bucket prefix for an export job.
	KeyPrefix *string `locationName:"keyPrefix" type:"string"`
}

// String returns the string representation
func (s S3DestinationConfig) String() string {
	return awsutil.Prettify(s)
}

// The summary of a recommendation.
type Summary struct {
	_ struct{} `type:"structure"`

	// The finding classification of the recommendation.
	Name Finding `locationName:"name" type:"string" enum:"true"`

	// The value of the recommendation summary.
	Value *float64 `locationName:"value" type:"double"`
}

// String returns the string representation
func (s Summary) String() string {
	return awsutil.Prettify(s)
}

// Describes a utilization metric of a resource, such as an Amazon EC2 instance.
type UtilizationMetric struct {
	_ struct{} `type:"structure"`

	// The name of the utilization metric.
	//
	// Memory metrics are only returned for resources that have the unified CloudWatch
	// agent installed on them. For more information, see Enabling Memory Utilization
	// with the CloudWatch Agent (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Install-CloudWatch-Agent.html).
	Name MetricName `locationName:"name" type:"string" enum:"true"`

	// The statistic of the utilization metric.
	Statistic MetricStatistic `locationName:"statistic" type:"string" enum:"true"`

	// The value of the utilization metric.
	Value *float64 `locationName:"value" type:"double"`
}

// String returns the string representation
func (s UtilizationMetric) String() string {
	return awsutil.Prettify(s)
}

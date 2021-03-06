// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains details about a package version asset.
type AssetSummary struct {
	_ struct{} `type:"structure"`

	// The hashes of the asset.
	Hashes map[string]string `locationName:"hashes" type:"map"`

	// The name of the asset.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The size of the asset.
	Size *int64 `locationName:"size" type:"long"`
}

// String returns the string representation
func (s AssetSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssetSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Hashes != nil {
		v := s.Hashes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "hashes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Size != nil {
		v := *s.Size

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "size", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Information about a domain. A domain is a container for repositories. When
// you create a domain, it is empty until you add one or more repositories.
type DomainDescription struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the domain.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The total size of all assets in the domain.
	AssetSizeBytes *int64 `locationName:"assetSizeBytes" type:"long"`

	// A timestamp that represents the date and time the domain was created.
	CreatedTime *time.Time `locationName:"createdTime" type:"timestamp"`

	// The ARN of an AWS Key Management Service (AWS KMS) key associated with a
	// domain.
	EncryptionKey *string `locationName:"encryptionKey" min:"1" type:"string"`

	// The name of the domain.
	Name *string `locationName:"name" min:"2" type:"string"`

	// The AWS account ID that owns the domain.
	Owner *string `locationName:"owner" min:"12" type:"string"`

	// The number of repositories in the domain.
	RepositoryCount *int64 `locationName:"repositoryCount" type:"integer"`

	// The current status of a domain. The valid values are
	//
	//    * Active
	//
	//    * Deleted
	Status DomainStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DomainDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DomainDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetSizeBytes != nil {
		v := *s.AssetSizeBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "assetSizeBytes", protocol.Int64Value(v), metadata)
	}
	if s.CreatedTime != nil {
		v := *s.CreatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.EncryptionKey != nil {
		v := *s.EncryptionKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "encryptionKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Owner != nil {
		v := *s.Owner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RepositoryCount != nil {
		v := *s.RepositoryCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repositoryCount", protocol.Int64Value(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about a domain, including its name, Amazon Resource Name (ARN),
// and status. The ListDomains (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListDomains.html)
// operation returns a list of DomainSummary objects.
type DomainSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the domain.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// A timestamp that contains the date and time the domain was created.
	CreatedTime *time.Time `locationName:"createdTime" type:"timestamp"`

	// The key used to encrypt the domain.
	EncryptionKey *string `locationName:"encryptionKey" min:"1" type:"string"`

	// The name of the domain.
	Name *string `locationName:"name" min:"2" type:"string"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	Owner *string `locationName:"owner" min:"12" type:"string"`

	// A string that contains the status of the domain. The valid values are:
	//
	//    * Active
	//
	//    * Deleted
	Status DomainStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DomainSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DomainSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedTime != nil {
		v := *s.CreatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.EncryptionKey != nil {
		v := *s.EncryptionKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "encryptionKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Owner != nil {
		v := *s.Owner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Details of the license data.
type LicenseInfo struct {
	_ struct{} `type:"structure"`

	// Name of the license.
	Name *string `locationName:"name" type:"string"`

	// The URL for license data.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s LicenseInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LicenseInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Url != nil {
		v := *s.Url

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "url", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a package dependency.
type PackageDependency struct {
	_ struct{} `type:"structure"`

	// The type of a package dependency. The possible values depend on the package
	// type. Example types are compile, runtime, and test for Maven packages, and
	// dev, prod, and optional for npm packages.
	DependencyType *string `locationName:"dependencyType" type:"string"`

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//    * The namespace of a Maven package is its groupId.
	//
	//    * The namespace of an npm package is its scope.
	//
	//    * A Python package does not contain a corresponding component, so Python
	//    packages do not have a namespace.
	Namespace *string `locationName:"namespace" min:"1" type:"string"`

	// The name of the package that this package depends on.
	Package *string `locationName:"package" min:"1" type:"string"`

	// The required version, or version range, of the package that this package
	// depends on. The version format is specific to the package type. For example,
	// the following are possible valid required versions: 1.2.3, ^2.3.4, or 4.x.
	VersionRequirement *string `locationName:"versionRequirement" type:"string"`
}

// String returns the string representation
func (s PackageDependency) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackageDependency) MarshalFields(e protocol.FieldEncoder) error {
	if s.DependencyType != nil {
		v := *s.DependencyType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dependencyType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Package != nil {
		v := *s.Package

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "package", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionRequirement != nil {
		v := *s.VersionRequirement

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionRequirement", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a package, including its format, namespace, and name. The ListPackages
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackages.html)
// operation returns a list of PackageSummary objects.
type PackageSummary struct {
	_ struct{} `type:"structure"`

	// The format of the package. Valid values are:
	//
	//    * npm
	//
	//    * pypi
	//
	//    * maven
	Format PackageFormat `locationName:"format" type:"string" enum:"true"`

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//    * The namespace of a Maven package is its groupId.
	//
	//    * The namespace of an npm package is its scope.
	//
	//    * A Python package does not contain a corresponding component, so Python
	//    packages do not have a namespace.
	Namespace *string `locationName:"namespace" min:"1" type:"string"`

	// The name of the package.
	Package *string `locationName:"package" min:"1" type:"string"`
}

// String returns the string representation
func (s PackageSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackageSummary) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Package != nil {
		v := *s.Package

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "package", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a package version.
type PackageVersionDescription struct {
	_ struct{} `type:"structure"`

	// The name of the package that is displayed. The displayName varies depending
	// on the package version's format. For example, if an npm package is named
	// ui, is in the namespace vue, and has the format npm, then the displayName
	// is @vue/ui.
	DisplayName *string `locationName:"displayName" min:"1" type:"string"`

	// The format of the package version. The valid package formats are:
	//
	//    * npm: A Node Package Manager (npm) package.
	//
	//    * pypi: A Python Package Index (PyPI) package.
	//
	//    * maven: A Maven package that contains compiled code in a distributable
	//    format, such as a JAR file.
	Format PackageFormat `locationName:"format" type:"string" enum:"true"`

	// The homepage associated with the package.
	HomePage *string `locationName:"homePage" type:"string"`

	// Information about licenses associated with the package version.
	Licenses []LicenseInfo `locationName:"licenses" type:"list"`

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//    * The namespace of a Maven package is its groupId.
	//
	//    * The namespace of an npm package is its scope.
	//
	//    * A Python package does not contain a corresponding component, so Python
	//    packages do not have a namespace.
	Namespace *string `locationName:"namespace" min:"1" type:"string"`

	// The name of the requested package.
	PackageName *string `locationName:"packageName" min:"1" type:"string"`

	// A timestamp that contains the date and time the package version was published.
	PublishedTime *time.Time `locationName:"publishedTime" type:"timestamp"`

	// The revision of the package version.
	Revision *string `locationName:"revision" min:"1" type:"string"`

	// The repository for the source code in the package version, or the source
	// code used to build it.
	SourceCodeRepository *string `locationName:"sourceCodeRepository" type:"string"`

	// A string that contains the status of the package version. It can be one of
	// the following:
	//
	//    * Published
	//
	//    * Unfinished
	//
	//    * Unlisted
	//
	//    * Archived
	//
	//    * Disposed
	Status PackageVersionStatus `locationName:"status" type:"string" enum:"true"`

	// A summary of the package version. The summary is extracted from the package.
	// The information in and detail level of the summary depends on the package
	// version's format.
	Summary *string `locationName:"summary" type:"string"`

	// The version of the package.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s PackageVersionDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackageVersionDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "displayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.HomePage != nil {
		v := *s.HomePage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "homePage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Licenses != nil {
		v := s.Licenses

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "licenses", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageName != nil {
		v := *s.PackageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PublishedTime != nil {
		v := *s.PublishedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "publishedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceCodeRepository != nil {
		v := *s.SourceCodeRepository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceCodeRepository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Summary != nil {
		v := *s.Summary

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "summary", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An error associated with package.
type PackageVersionError struct {
	_ struct{} `type:"structure"`

	// The error code associated with the error. Valid error codes are:
	//
	//    * ALREADY_EXISTS
	//
	//    * MISMATCHED_REVISION
	//
	//    * MISMATCHED_STATUS
	//
	//    * NOT_ALLOWED
	//
	//    * NOT_FOUND
	//
	//    * SKIPPED
	ErrorCode PackageVersionErrorCode `locationName:"errorCode" type:"string" enum:"true"`

	// The error message associated with the error.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`
}

// String returns the string representation
func (s PackageVersionError) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackageVersionError) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ErrorCode) > 0 {
		v := s.ErrorCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ErrorMessage != nil {
		v := *s.ErrorMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a package version, including its status, version, and revision.
// The ListPackageVersions (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)
// operation returns a list of PackageVersionSummary objects.
type PackageVersionSummary struct {
	_ struct{} `type:"structure"`

	// The revision associated with a package version.
	Revision *string `locationName:"revision" min:"1" type:"string"`

	// A string that contains the status of the package version. It can be one of
	// the following:
	//
	//    * Published
	//
	//    * Unfinished
	//
	//    * Unlisted
	//
	//    * Archived
	//
	//    * Disposed
	//
	// Status is a required field
	Status PackageVersionStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// Information about a package version.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PackageVersionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackageVersionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The details of a repository stored in AWS CodeArtifact. A CodeArtifact repository
// contains a set of package versions, each of which maps to a set of assets.
// Repositories are polyglot—a single repository can contain packages of any
// supported type. Each repository exposes endpoints for fetching and publishing
// packages using tools like the npm CLI, the Maven CLI (mvn), and pip. You
// can create up to 100 repositories per AWS account.
type RepositoryDescription struct {
	_ struct{} `type:"structure"`

	// The 12-digit account number of the AWS account that manages the repository.
	AdministratorAccount *string `locationName:"administratorAccount" min:"12" type:"string"`

	// The Amazon Resource Name (ARN) of the repository.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// A text description of the repository.
	Description *string `locationName:"description" type:"string"`

	// The name of the domain that contains the repository.
	DomainName *string `locationName:"domainName" min:"2" type:"string"`

	// The 12-digit account number of the AWS account that owns the domain that
	// contains the repository. It does not include dashes or spaces.
	DomainOwner *string `locationName:"domainOwner" min:"12" type:"string"`

	// An array of external connections associated with the repository.
	ExternalConnections []RepositoryExternalConnectionInfo `locationName:"externalConnections" type:"list"`

	// The name of the repository.
	Name *string `locationName:"name" min:"2" type:"string"`

	// A list of upstream repositories to associate with the repository. The order
	// of the upstream repositories in the list determines their priority order
	// when AWS CodeArtifact looks for a requested package version. For more information,
	// see Working with upstream repositories (https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html).
	Upstreams []UpstreamRepositoryInfo `locationName:"upstreams" type:"list"`
}

// String returns the string representation
func (s RepositoryDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RepositoryDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdministratorAccount != nil {
		v := *s.AdministratorAccount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "administratorAccount", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExternalConnections != nil {
		v := s.ExternalConnections

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "externalConnections", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Upstreams != nil {
		v := s.Upstreams

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "upstreams", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Contains information about the external connection of a repository.
type RepositoryExternalConnectionInfo struct {
	_ struct{} `type:"structure"`

	// The name of the external connection associated with a repository.
	ExternalConnectionName *string `locationName:"externalConnectionName" type:"string"`

	// The package format associated with a repository's external connection. The
	// valid package formats are:
	//
	//    * npm: A Node Package Manager (npm) package.
	//
	//    * pypi: A Python Package Index (PyPI) package.
	//
	//    * maven: A Maven package that contains compiled code in a distributable
	//    format, such as a JAR file.
	PackageFormat PackageFormat `locationName:"packageFormat" type:"string" enum:"true"`

	// The status of the external connection of a repository. There is one valid
	// value, Available.
	Status ExternalConnectionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s RepositoryExternalConnectionInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RepositoryExternalConnectionInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExternalConnectionName != nil {
		v := *s.ExternalConnectionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "externalConnectionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PackageFormat) > 0 {
		v := s.PackageFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packageFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Details about a repository, including its Amazon Resource Name (ARN), description,
// and domain information. The ListRepositories (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListRepositories.html)
// operation returns a list of RepositorySummary objects.
type RepositorySummary struct {
	_ struct{} `type:"structure"`

	// The AWS account ID that manages the repository.
	AdministratorAccount *string `locationName:"administratorAccount" min:"12" type:"string"`

	// The ARN of the repository.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The description of the repository.
	Description *string `locationName:"description" type:"string"`

	// The name of the domain that contains the repository.
	DomainName *string `locationName:"domainName" min:"2" type:"string"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `locationName:"domainOwner" min:"12" type:"string"`

	// The name of the repository.
	Name *string `locationName:"name" min:"2" type:"string"`
}

// String returns the string representation
func (s RepositorySummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RepositorySummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdministratorAccount != nil {
		v := *s.AdministratorAccount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "administratorAccount", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An AWS CodeArtifact resource policy that contains a resource ARN, document
// details, and a revision.
type ResourcePolicy struct {
	_ struct{} `type:"structure"`

	// The resource policy formatted in JSON.
	Document *string `locationName:"document" min:"1" type:"string"`

	// The ARN of the resource associated with the resource policy
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string"`

	// The current revision of the resource policy.
	Revision *string `locationName:"revision" min:"1" type:"string"`
}

// String returns the string representation
func (s ResourcePolicy) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourcePolicy) MarshalFields(e protocol.FieldEncoder) error {
	if s.Document != nil {
		v := *s.Document

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "document", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the revision and status of a package version.
type SuccessfulPackageVersionInfo struct {
	_ struct{} `type:"structure"`

	// The revision of a package version.
	Revision *string `locationName:"revision" type:"string"`

	// The status of a package version. Valid statuses are:
	//
	//    * Published
	//
	//    * Unfinished
	//
	//    * Unlisted
	//
	//    * Archived
	//
	//    * Disposed
	Status PackageVersionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s SuccessfulPackageVersionInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SuccessfulPackageVersionInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about an upstream repository. A list of UpstreamRepository objects
// is an input parameter to CreateRepository (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_CreateRepository.html)
// and UpdateRepository (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdateRepository.html).
type UpstreamRepository struct {
	_ struct{} `type:"structure"`

	// The name of an upstream repository.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s UpstreamRepository) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpstreamRepository) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpstreamRepository"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpstreamRepository) MarshalFields(e protocol.FieldEncoder) error {
	if s.RepositoryName != nil {
		v := *s.RepositoryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repositoryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about an upstream repository.
type UpstreamRepositoryInfo struct {
	_ struct{} `type:"structure"`

	// The name of an upstream repository.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s UpstreamRepositoryInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpstreamRepositoryInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.RepositoryName != nil {
		v := *s.RepositoryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repositoryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

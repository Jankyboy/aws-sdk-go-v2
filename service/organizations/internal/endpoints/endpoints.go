// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver Organizations endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "organizations.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:       regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
		IsRegionalized:    false,
		PartitionEndpoint: "aws-global",
		Endpoints: endpoints.Endpoints{
			"aws-global": endpoints.Endpoint{
				Hostname: "organizations.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"fips-aws-global": endpoints.Endpoint{
				Hostname: "organizations-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "organizations.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:       regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
		IsRegionalized:    false,
		PartitionEndpoint: "aws-cn-global",
		Endpoints: endpoints.Endpoints{
			"aws-cn-global": endpoints.Endpoint{
				Hostname: "organizations.cn-northwest-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
			"fips-aws-cn-global": endpoints.Endpoint{
				Hostname: "organizations.cn-northwest-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "organizations.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "organizations.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "organizations.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:       regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
		IsRegionalized:    false,
		PartitionEndpoint: "aws-us-gov-global",
		Endpoints: endpoints.Endpoints{
			"aws-us-gov-global": endpoints.Endpoint{
				Hostname: "organizations.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"fips-aws-us-gov-global": endpoints.Endpoint{
				Hostname: "organizations.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
		},
	},
}

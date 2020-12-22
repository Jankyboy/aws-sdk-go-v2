// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

func ExampleAccessLog_outputUsage() {
	var union types.AccessLog
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AccessLogMemberFile:
		_ = v.Value // Value is types.FileAccessLog

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FileAccessLog

func ExampleBackend_outputUsage() {
	var union types.Backend
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.BackendMemberVirtualService:
		_ = v.Value // Value is types.VirtualServiceBackend

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VirtualServiceBackend

func ExampleGrpcRouteMetadataMatchMethod_outputUsage() {
	var union types.GrpcRouteMetadataMatchMethod
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GrpcRouteMetadataMatchMethodMemberExact:
		_ = v.Value // Value is string

	case *types.GrpcRouteMetadataMatchMethodMemberPrefix:
		_ = v.Value // Value is string

	case *types.GrpcRouteMetadataMatchMethodMemberRange:
		_ = v.Value // Value is types.MatchRange

	case *types.GrpcRouteMetadataMatchMethodMemberRegex:
		_ = v.Value // Value is string

	case *types.GrpcRouteMetadataMatchMethodMemberSuffix:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MatchRange
var _ *string

func ExampleHeaderMatchMethod_outputUsage() {
	var union types.HeaderMatchMethod
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.HeaderMatchMethodMemberExact:
		_ = v.Value // Value is string

	case *types.HeaderMatchMethodMemberPrefix:
		_ = v.Value // Value is string

	case *types.HeaderMatchMethodMemberRange:
		_ = v.Value // Value is types.MatchRange

	case *types.HeaderMatchMethodMemberRegex:
		_ = v.Value // Value is string

	case *types.HeaderMatchMethodMemberSuffix:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MatchRange
var _ *string

func ExampleListenerTimeout_outputUsage() {
	var union types.ListenerTimeout
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ListenerTimeoutMemberGrpc:
		_ = v.Value // Value is types.GrpcTimeout

	case *types.ListenerTimeoutMemberHttp:
		_ = v.Value // Value is types.HttpTimeout

	case *types.ListenerTimeoutMemberHttp2:
		_ = v.Value // Value is types.HttpTimeout

	case *types.ListenerTimeoutMemberTcp:
		_ = v.Value // Value is types.TcpTimeout

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TcpTimeout
var _ *types.HttpTimeout
var _ *types.GrpcTimeout

func ExampleListenerTlsCertificate_outputUsage() {
	var union types.ListenerTlsCertificate
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ListenerTlsCertificateMemberAcm:
		_ = v.Value // Value is types.ListenerTlsAcmCertificate

	case *types.ListenerTlsCertificateMemberFile:
		_ = v.Value // Value is types.ListenerTlsFileCertificate

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ListenerTlsAcmCertificate
var _ *types.ListenerTlsFileCertificate

func ExampleServiceDiscovery_outputUsage() {
	var union types.ServiceDiscovery
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ServiceDiscoveryMemberAwsCloudMap:
		_ = v.Value // Value is types.AwsCloudMapServiceDiscovery

	case *types.ServiceDiscoveryMemberDns:
		_ = v.Value // Value is types.DnsServiceDiscovery

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AwsCloudMapServiceDiscovery
var _ *types.DnsServiceDiscovery

func ExampleTlsValidationContextTrust_outputUsage() {
	var union types.TlsValidationContextTrust
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TlsValidationContextTrustMemberAcm:
		_ = v.Value // Value is types.TlsValidationContextAcmTrust

	case *types.TlsValidationContextTrustMemberFile:
		_ = v.Value // Value is types.TlsValidationContextFileTrust

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TlsValidationContextAcmTrust
var _ *types.TlsValidationContextFileTrust

func ExampleVirtualGatewayAccessLog_outputUsage() {
	var union types.VirtualGatewayAccessLog
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VirtualGatewayAccessLogMemberFile:
		_ = v.Value // Value is types.VirtualGatewayFileAccessLog

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VirtualGatewayFileAccessLog

func ExampleVirtualGatewayListenerTlsCertificate_outputUsage() {
	var union types.VirtualGatewayListenerTlsCertificate
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VirtualGatewayListenerTlsCertificateMemberAcm:
		_ = v.Value // Value is types.VirtualGatewayListenerTlsAcmCertificate

	case *types.VirtualGatewayListenerTlsCertificateMemberFile:
		_ = v.Value // Value is types.VirtualGatewayListenerTlsFileCertificate

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VirtualGatewayListenerTlsFileCertificate
var _ *types.VirtualGatewayListenerTlsAcmCertificate

func ExampleVirtualGatewayTlsValidationContextTrust_outputUsage() {
	var union types.VirtualGatewayTlsValidationContextTrust
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VirtualGatewayTlsValidationContextTrustMemberAcm:
		_ = v.Value // Value is types.VirtualGatewayTlsValidationContextAcmTrust

	case *types.VirtualGatewayTlsValidationContextTrustMemberFile:
		_ = v.Value // Value is types.VirtualGatewayTlsValidationContextFileTrust

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VirtualGatewayTlsValidationContextAcmTrust
var _ *types.VirtualGatewayTlsValidationContextFileTrust

func ExampleVirtualServiceProvider_outputUsage() {
	var union types.VirtualServiceProvider
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VirtualServiceProviderMemberVirtualNode:
		_ = v.Value // Value is types.VirtualNodeServiceProvider

	case *types.VirtualServiceProviderMemberVirtualRouter:
		_ = v.Value // Value is types.VirtualRouterServiceProvider

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VirtualRouterServiceProvider
var _ *types.VirtualNodeServiceProvider

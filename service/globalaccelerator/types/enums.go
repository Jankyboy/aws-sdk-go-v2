// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AcceleratorStatus string

// Enum values for AcceleratorStatus
const (
	AcceleratorStatusDeployed   AcceleratorStatus = "DEPLOYED"
	AcceleratorStatusInProgress AcceleratorStatus = "IN_PROGRESS"
)

// Values returns all known values for AcceleratorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AcceleratorStatus) Values() []AcceleratorStatus {
	return []AcceleratorStatus{
		"DEPLOYED",
		"IN_PROGRESS",
	}
}

type ByoipCidrState string

// Enum values for ByoipCidrState
const (
	ByoipCidrStatePendingProvisioning   ByoipCidrState = "PENDING_PROVISIONING"
	ByoipCidrStateReady                 ByoipCidrState = "READY"
	ByoipCidrStatePendingAdvertising    ByoipCidrState = "PENDING_ADVERTISING"
	ByoipCidrStateAdvertising           ByoipCidrState = "ADVERTISING"
	ByoipCidrStatePendingWithdrawing    ByoipCidrState = "PENDING_WITHDRAWING"
	ByoipCidrStatePendingDeprovisioning ByoipCidrState = "PENDING_DEPROVISIONING"
	ByoipCidrStateDeprovisioned         ByoipCidrState = "DEPROVISIONED"
	ByoipCidrStateFailedProvision       ByoipCidrState = "FAILED_PROVISION"
	ByoipCidrStateFailedAdvertising     ByoipCidrState = "FAILED_ADVERTISING"
	ByoipCidrStateFailedWithdraw        ByoipCidrState = "FAILED_WITHDRAW"
	ByoipCidrStateFailedDeprovision     ByoipCidrState = "FAILED_DEPROVISION"
)

// Values returns all known values for ByoipCidrState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ByoipCidrState) Values() []ByoipCidrState {
	return []ByoipCidrState{
		"PENDING_PROVISIONING",
		"READY",
		"PENDING_ADVERTISING",
		"ADVERTISING",
		"PENDING_WITHDRAWING",
		"PENDING_DEPROVISIONING",
		"DEPROVISIONED",
		"FAILED_PROVISION",
		"FAILED_ADVERTISING",
		"FAILED_WITHDRAW",
		"FAILED_DEPROVISION",
	}
}

type ClientAffinity string

// Enum values for ClientAffinity
const (
	ClientAffinityNone     ClientAffinity = "NONE"
	ClientAffinitySourceIp ClientAffinity = "SOURCE_IP"
)

// Values returns all known values for ClientAffinity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClientAffinity) Values() []ClientAffinity {
	return []ClientAffinity{
		"NONE",
		"SOURCE_IP",
	}
}

type HealthCheckProtocol string

// Enum values for HealthCheckProtocol
const (
	HealthCheckProtocolTcp   HealthCheckProtocol = "TCP"
	HealthCheckProtocolHttp  HealthCheckProtocol = "HTTP"
	HealthCheckProtocolHttps HealthCheckProtocol = "HTTPS"
)

// Values returns all known values for HealthCheckProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HealthCheckProtocol) Values() []HealthCheckProtocol {
	return []HealthCheckProtocol{
		"TCP",
		"HTTP",
		"HTTPS",
	}
}

type HealthState string

// Enum values for HealthState
const (
	HealthStateInitial   HealthState = "INITIAL"
	HealthStateHealthy   HealthState = "HEALTHY"
	HealthStateUnhealthy HealthState = "UNHEALTHY"
)

// Values returns all known values for HealthState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HealthState) Values() []HealthState {
	return []HealthState{
		"INITIAL",
		"HEALTHY",
		"UNHEALTHY",
	}
}

type IpAddressType string

// Enum values for IpAddressType
const (
	IpAddressTypeIpv4 IpAddressType = "IPV4"
)

// Values returns all known values for IpAddressType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressType) Values() []IpAddressType {
	return []IpAddressType{
		"IPV4",
	}
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolTcp Protocol = "TCP"
	ProtocolUdp Protocol = "UDP"
)

// Values returns all known values for Protocol. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Protocol) Values() []Protocol {
	return []Protocol{
		"TCP",
		"UDP",
	}
}

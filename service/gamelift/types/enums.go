// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AcceptanceType string

// Enum values for AcceptanceType
const (
	AcceptanceTypeAccept AcceptanceType = "ACCEPT"
	AcceptanceTypeReject AcceptanceType = "REJECT"
)

// Values returns all known values for AcceptanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AcceptanceType) Values() []AcceptanceType {
	return []AcceptanceType{
		"ACCEPT",
		"REJECT",
	}
}

type BackfillMode string

// Enum values for BackfillMode
const (
	BackfillModeAutomatic BackfillMode = "AUTOMATIC"
	BackfillModeManual    BackfillMode = "MANUAL"
)

// Values returns all known values for BackfillMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BackfillMode) Values() []BackfillMode {
	return []BackfillMode{
		"AUTOMATIC",
		"MANUAL",
	}
}

type BalancingStrategy string

// Enum values for BalancingStrategy
const (
	BalancingStrategySpotOnly      BalancingStrategy = "SPOT_ONLY"
	BalancingStrategySpotPreferred BalancingStrategy = "SPOT_PREFERRED"
	BalancingStrategyOnDemandOnly  BalancingStrategy = "ON_DEMAND_ONLY"
)

// Values returns all known values for BalancingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BalancingStrategy) Values() []BalancingStrategy {
	return []BalancingStrategy{
		"SPOT_ONLY",
		"SPOT_PREFERRED",
		"ON_DEMAND_ONLY",
	}
}

type BuildStatus string

// Enum values for BuildStatus
const (
	BuildStatusInitialized BuildStatus = "INITIALIZED"
	BuildStatusReady       BuildStatus = "READY"
	BuildStatusFailed      BuildStatus = "FAILED"
)

// Values returns all known values for BuildStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BuildStatus) Values() []BuildStatus {
	return []BuildStatus{
		"INITIALIZED",
		"READY",
		"FAILED",
	}
}

type CertificateType string

// Enum values for CertificateType
const (
	CertificateTypeDisabled  CertificateType = "DISABLED"
	CertificateTypeGenerated CertificateType = "GENERATED"
)

// Values returns all known values for CertificateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CertificateType) Values() []CertificateType {
	return []CertificateType{
		"DISABLED",
		"GENERATED",
	}
}

type ComparisonOperatorType string

// Enum values for ComparisonOperatorType
const (
	ComparisonOperatorTypeGreaterthanorequaltothreshold ComparisonOperatorType = "GreaterThanOrEqualToThreshold"
	ComparisonOperatorTypeGreaterthanthreshold          ComparisonOperatorType = "GreaterThanThreshold"
	ComparisonOperatorTypeLessthanthreshold             ComparisonOperatorType = "LessThanThreshold"
	ComparisonOperatorTypeLessthanorequaltothreshold    ComparisonOperatorType = "LessThanOrEqualToThreshold"
)

// Values returns all known values for ComparisonOperatorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperatorType) Values() []ComparisonOperatorType {
	return []ComparisonOperatorType{
		"GreaterThanOrEqualToThreshold",
		"GreaterThanThreshold",
		"LessThanThreshold",
		"LessThanOrEqualToThreshold",
	}
}

type EC2InstanceType string

// Enum values for EC2InstanceType
const (
	EC2InstanceTypeT2Micro    EC2InstanceType = "t2.micro"
	EC2InstanceTypeT2Small    EC2InstanceType = "t2.small"
	EC2InstanceTypeT2Medium   EC2InstanceType = "t2.medium"
	EC2InstanceTypeT2Large    EC2InstanceType = "t2.large"
	EC2InstanceTypeC3Large    EC2InstanceType = "c3.large"
	EC2InstanceTypeC3Xlarge   EC2InstanceType = "c3.xlarge"
	EC2InstanceTypeC32xlarge  EC2InstanceType = "c3.2xlarge"
	EC2InstanceTypeC34xlarge  EC2InstanceType = "c3.4xlarge"
	EC2InstanceTypeC38xlarge  EC2InstanceType = "c3.8xlarge"
	EC2InstanceTypeC4Large    EC2InstanceType = "c4.large"
	EC2InstanceTypeC4Xlarge   EC2InstanceType = "c4.xlarge"
	EC2InstanceTypeC42xlarge  EC2InstanceType = "c4.2xlarge"
	EC2InstanceTypeC44xlarge  EC2InstanceType = "c4.4xlarge"
	EC2InstanceTypeC48xlarge  EC2InstanceType = "c4.8xlarge"
	EC2InstanceTypeC5Large    EC2InstanceType = "c5.large"
	EC2InstanceTypeC5Xlarge   EC2InstanceType = "c5.xlarge"
	EC2InstanceTypeC52xlarge  EC2InstanceType = "c5.2xlarge"
	EC2InstanceTypeC54xlarge  EC2InstanceType = "c5.4xlarge"
	EC2InstanceTypeC59xlarge  EC2InstanceType = "c5.9xlarge"
	EC2InstanceTypeC512xlarge EC2InstanceType = "c5.12xlarge"
	EC2InstanceTypeC518xlarge EC2InstanceType = "c5.18xlarge"
	EC2InstanceTypeC524xlarge EC2InstanceType = "c5.24xlarge"
	EC2InstanceTypeR3Large    EC2InstanceType = "r3.large"
	EC2InstanceTypeR3Xlarge   EC2InstanceType = "r3.xlarge"
	EC2InstanceTypeR32xlarge  EC2InstanceType = "r3.2xlarge"
	EC2InstanceTypeR34xlarge  EC2InstanceType = "r3.4xlarge"
	EC2InstanceTypeR38xlarge  EC2InstanceType = "r3.8xlarge"
	EC2InstanceTypeR4Large    EC2InstanceType = "r4.large"
	EC2InstanceTypeR4Xlarge   EC2InstanceType = "r4.xlarge"
	EC2InstanceTypeR42xlarge  EC2InstanceType = "r4.2xlarge"
	EC2InstanceTypeR44xlarge  EC2InstanceType = "r4.4xlarge"
	EC2InstanceTypeR48xlarge  EC2InstanceType = "r4.8xlarge"
	EC2InstanceTypeR416xlarge EC2InstanceType = "r4.16xlarge"
	EC2InstanceTypeR5Large    EC2InstanceType = "r5.large"
	EC2InstanceTypeR5Xlarge   EC2InstanceType = "r5.xlarge"
	EC2InstanceTypeR52xlarge  EC2InstanceType = "r5.2xlarge"
	EC2InstanceTypeR54xlarge  EC2InstanceType = "r5.4xlarge"
	EC2InstanceTypeR58xlarge  EC2InstanceType = "r5.8xlarge"
	EC2InstanceTypeR512xlarge EC2InstanceType = "r5.12xlarge"
	EC2InstanceTypeR516xlarge EC2InstanceType = "r5.16xlarge"
	EC2InstanceTypeR524xlarge EC2InstanceType = "r5.24xlarge"
	EC2InstanceTypeM3Medium   EC2InstanceType = "m3.medium"
	EC2InstanceTypeM3Large    EC2InstanceType = "m3.large"
	EC2InstanceTypeM3Xlarge   EC2InstanceType = "m3.xlarge"
	EC2InstanceTypeM32xlarge  EC2InstanceType = "m3.2xlarge"
	EC2InstanceTypeM4Large    EC2InstanceType = "m4.large"
	EC2InstanceTypeM4Xlarge   EC2InstanceType = "m4.xlarge"
	EC2InstanceTypeM42xlarge  EC2InstanceType = "m4.2xlarge"
	EC2InstanceTypeM44xlarge  EC2InstanceType = "m4.4xlarge"
	EC2InstanceTypeM410xlarge EC2InstanceType = "m4.10xlarge"
	EC2InstanceTypeM5Large    EC2InstanceType = "m5.large"
	EC2InstanceTypeM5Xlarge   EC2InstanceType = "m5.xlarge"
	EC2InstanceTypeM52xlarge  EC2InstanceType = "m5.2xlarge"
	EC2InstanceTypeM54xlarge  EC2InstanceType = "m5.4xlarge"
	EC2InstanceTypeM58xlarge  EC2InstanceType = "m5.8xlarge"
	EC2InstanceTypeM512xlarge EC2InstanceType = "m5.12xlarge"
	EC2InstanceTypeM516xlarge EC2InstanceType = "m5.16xlarge"
	EC2InstanceTypeM524xlarge EC2InstanceType = "m5.24xlarge"
)

// Values returns all known values for EC2InstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EC2InstanceType) Values() []EC2InstanceType {
	return []EC2InstanceType{
		"t2.micro",
		"t2.small",
		"t2.medium",
		"t2.large",
		"c3.large",
		"c3.xlarge",
		"c3.2xlarge",
		"c3.4xlarge",
		"c3.8xlarge",
		"c4.large",
		"c4.xlarge",
		"c4.2xlarge",
		"c4.4xlarge",
		"c4.8xlarge",
		"c5.large",
		"c5.xlarge",
		"c5.2xlarge",
		"c5.4xlarge",
		"c5.9xlarge",
		"c5.12xlarge",
		"c5.18xlarge",
		"c5.24xlarge",
		"r3.large",
		"r3.xlarge",
		"r3.2xlarge",
		"r3.4xlarge",
		"r3.8xlarge",
		"r4.large",
		"r4.xlarge",
		"r4.2xlarge",
		"r4.4xlarge",
		"r4.8xlarge",
		"r4.16xlarge",
		"r5.large",
		"r5.xlarge",
		"r5.2xlarge",
		"r5.4xlarge",
		"r5.8xlarge",
		"r5.12xlarge",
		"r5.16xlarge",
		"r5.24xlarge",
		"m3.medium",
		"m3.large",
		"m3.xlarge",
		"m3.2xlarge",
		"m4.large",
		"m4.xlarge",
		"m4.2xlarge",
		"m4.4xlarge",
		"m4.10xlarge",
		"m5.large",
		"m5.xlarge",
		"m5.2xlarge",
		"m5.4xlarge",
		"m5.8xlarge",
		"m5.12xlarge",
		"m5.16xlarge",
		"m5.24xlarge",
	}
}

type EventCode string

// Enum values for EventCode
const (
	EventCodeGenericEvent                               EventCode = "GENERIC_EVENT"
	EventCodeFleetCreated                               EventCode = "FLEET_CREATED"
	EventCodeFleetDeleted                               EventCode = "FLEET_DELETED"
	EventCodeFleetScalingEvent                          EventCode = "FLEET_SCALING_EVENT"
	EventCodeFleetStateDownloading                      EventCode = "FLEET_STATE_DOWNLOADING"
	EventCodeFleetStateValidating                       EventCode = "FLEET_STATE_VALIDATING"
	EventCodeFleetStateBuilding                         EventCode = "FLEET_STATE_BUILDING"
	EventCodeFleetStateActivating                       EventCode = "FLEET_STATE_ACTIVATING"
	EventCodeFleetStateActive                           EventCode = "FLEET_STATE_ACTIVE"
	EventCodeFleetStateError                            EventCode = "FLEET_STATE_ERROR"
	EventCodeFleetInitializationFailed                  EventCode = "FLEET_INITIALIZATION_FAILED"
	EventCodeFleetBinaryDownloadFailed                  EventCode = "FLEET_BINARY_DOWNLOAD_FAILED"
	EventCodeFleetValidationLaunchPathNotFound          EventCode = "FLEET_VALIDATION_LAUNCH_PATH_NOT_FOUND"
	EventCodeFleetValidationExecutableRuntimeFailure    EventCode = "FLEET_VALIDATION_EXECUTABLE_RUNTIME_FAILURE"
	EventCodeFleetValidationTimedOut                    EventCode = "FLEET_VALIDATION_TIMED_OUT"
	EventCodeFleetActivationFailed                      EventCode = "FLEET_ACTIVATION_FAILED"
	EventCodeFleetActivationFailedNoInstances           EventCode = "FLEET_ACTIVATION_FAILED_NO_INSTANCES"
	EventCodeFleetNewGameSessionProtectionPolicyUpdated EventCode = "FLEET_NEW_GAME_SESSION_PROTECTION_POLICY_UPDATED"
	EventCodeServerProcessInvalidPath                   EventCode = "SERVER_PROCESS_INVALID_PATH"
	EventCodeServerProcessSdkInitializationTimeout      EventCode = "SERVER_PROCESS_SDK_INITIALIZATION_TIMEOUT"
	EventCodeServerProcessProcessReadyTimeout           EventCode = "SERVER_PROCESS_PROCESS_READY_TIMEOUT"
	EventCodeServerProcessCrashed                       EventCode = "SERVER_PROCESS_CRASHED"
	EventCodeServerProcessTerminatedUnhealthy           EventCode = "SERVER_PROCESS_TERMINATED_UNHEALTHY"
	EventCodeServerProcessForceTerminated               EventCode = "SERVER_PROCESS_FORCE_TERMINATED"
	EventCodeServerProcessProcessExitTimeout            EventCode = "SERVER_PROCESS_PROCESS_EXIT_TIMEOUT"
	EventCodeGameSessionActivationTimeout               EventCode = "GAME_SESSION_ACTIVATION_TIMEOUT"
	EventCodeFleetCreationExtractingBuild               EventCode = "FLEET_CREATION_EXTRACTING_BUILD"
	EventCodeFleetCreationRunningInstaller              EventCode = "FLEET_CREATION_RUNNING_INSTALLER"
	EventCodeFleetCreationValidatingRuntimeConfig       EventCode = "FLEET_CREATION_VALIDATING_RUNTIME_CONFIG"
	EventCodeFleetVpcPeeringSucceeded                   EventCode = "FLEET_VPC_PEERING_SUCCEEDED"
	EventCodeFleetVpcPeeringFailed                      EventCode = "FLEET_VPC_PEERING_FAILED"
	EventCodeFleetVpcPeeringDeleted                     EventCode = "FLEET_VPC_PEERING_DELETED"
	EventCodeInstanceInterrupted                        EventCode = "INSTANCE_INTERRUPTED"
)

// Values returns all known values for EventCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (EventCode) Values() []EventCode {
	return []EventCode{
		"GENERIC_EVENT",
		"FLEET_CREATED",
		"FLEET_DELETED",
		"FLEET_SCALING_EVENT",
		"FLEET_STATE_DOWNLOADING",
		"FLEET_STATE_VALIDATING",
		"FLEET_STATE_BUILDING",
		"FLEET_STATE_ACTIVATING",
		"FLEET_STATE_ACTIVE",
		"FLEET_STATE_ERROR",
		"FLEET_INITIALIZATION_FAILED",
		"FLEET_BINARY_DOWNLOAD_FAILED",
		"FLEET_VALIDATION_LAUNCH_PATH_NOT_FOUND",
		"FLEET_VALIDATION_EXECUTABLE_RUNTIME_FAILURE",
		"FLEET_VALIDATION_TIMED_OUT",
		"FLEET_ACTIVATION_FAILED",
		"FLEET_ACTIVATION_FAILED_NO_INSTANCES",
		"FLEET_NEW_GAME_SESSION_PROTECTION_POLICY_UPDATED",
		"SERVER_PROCESS_INVALID_PATH",
		"SERVER_PROCESS_SDK_INITIALIZATION_TIMEOUT",
		"SERVER_PROCESS_PROCESS_READY_TIMEOUT",
		"SERVER_PROCESS_CRASHED",
		"SERVER_PROCESS_TERMINATED_UNHEALTHY",
		"SERVER_PROCESS_FORCE_TERMINATED",
		"SERVER_PROCESS_PROCESS_EXIT_TIMEOUT",
		"GAME_SESSION_ACTIVATION_TIMEOUT",
		"FLEET_CREATION_EXTRACTING_BUILD",
		"FLEET_CREATION_RUNNING_INSTALLER",
		"FLEET_CREATION_VALIDATING_RUNTIME_CONFIG",
		"FLEET_VPC_PEERING_SUCCEEDED",
		"FLEET_VPC_PEERING_FAILED",
		"FLEET_VPC_PEERING_DELETED",
		"INSTANCE_INTERRUPTED",
	}
}

type FleetAction string

// Enum values for FleetAction
const (
	FleetActionAutoscaling FleetAction = "AUTO_SCALING"
)

// Values returns all known values for FleetAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FleetAction) Values() []FleetAction {
	return []FleetAction{
		"AUTO_SCALING",
	}
}

type FleetStatus string

// Enum values for FleetStatus
const (
	FleetStatusNew         FleetStatus = "NEW"
	FleetStatusDownloading FleetStatus = "DOWNLOADING"
	FleetStatusValidating  FleetStatus = "VALIDATING"
	FleetStatusBuilding    FleetStatus = "BUILDING"
	FleetStatusActivating  FleetStatus = "ACTIVATING"
	FleetStatusActive      FleetStatus = "ACTIVE"
	FleetStatusDeleting    FleetStatus = "DELETING"
	FleetStatusError       FleetStatus = "ERROR"
	FleetStatusTerminated  FleetStatus = "TERMINATED"
)

// Values returns all known values for FleetStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FleetStatus) Values() []FleetStatus {
	return []FleetStatus{
		"NEW",
		"DOWNLOADING",
		"VALIDATING",
		"BUILDING",
		"ACTIVATING",
		"ACTIVE",
		"DELETING",
		"ERROR",
		"TERMINATED",
	}
}

type FleetType string

// Enum values for FleetType
const (
	FleetTypeOndemand FleetType = "ON_DEMAND"
	FleetTypeSpot     FleetType = "SPOT"
)

// Values returns all known values for FleetType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (FleetType) Values() []FleetType {
	return []FleetType{
		"ON_DEMAND",
		"SPOT",
	}
}

type GameServerClaimStatus string

// Enum values for GameServerClaimStatus
const (
	GameServerClaimStatusClaimed GameServerClaimStatus = "CLAIMED"
)

// Values returns all known values for GameServerClaimStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameServerClaimStatus) Values() []GameServerClaimStatus {
	return []GameServerClaimStatus{
		"CLAIMED",
	}
}

type GameServerGroupAction string

// Enum values for GameServerGroupAction
const (
	GameServerGroupActionReplaceInstanceTypes GameServerGroupAction = "REPLACE_INSTANCE_TYPES"
)

// Values returns all known values for GameServerGroupAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameServerGroupAction) Values() []GameServerGroupAction {
	return []GameServerGroupAction{
		"REPLACE_INSTANCE_TYPES",
	}
}

type GameServerGroupDeleteOption string

// Enum values for GameServerGroupDeleteOption
const (
	GameServerGroupDeleteOptionSafeDelete  GameServerGroupDeleteOption = "SAFE_DELETE"
	GameServerGroupDeleteOptionForceDelete GameServerGroupDeleteOption = "FORCE_DELETE"
	GameServerGroupDeleteOptionRetain      GameServerGroupDeleteOption = "RETAIN"
)

// Values returns all known values for GameServerGroupDeleteOption. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GameServerGroupDeleteOption) Values() []GameServerGroupDeleteOption {
	return []GameServerGroupDeleteOption{
		"SAFE_DELETE",
		"FORCE_DELETE",
		"RETAIN",
	}
}

type GameServerGroupInstanceType string

// Enum values for GameServerGroupInstanceType
const (
	GameServerGroupInstanceTypeC4Large    GameServerGroupInstanceType = "c4.large"
	GameServerGroupInstanceTypeC4Xlarge   GameServerGroupInstanceType = "c4.xlarge"
	GameServerGroupInstanceTypeC42xlarge  GameServerGroupInstanceType = "c4.2xlarge"
	GameServerGroupInstanceTypeC44xlarge  GameServerGroupInstanceType = "c4.4xlarge"
	GameServerGroupInstanceTypeC48xlarge  GameServerGroupInstanceType = "c4.8xlarge"
	GameServerGroupInstanceTypeC5Large    GameServerGroupInstanceType = "c5.large"
	GameServerGroupInstanceTypeC5Xlarge   GameServerGroupInstanceType = "c5.xlarge"
	GameServerGroupInstanceTypeC52xlarge  GameServerGroupInstanceType = "c5.2xlarge"
	GameServerGroupInstanceTypeC54xlarge  GameServerGroupInstanceType = "c5.4xlarge"
	GameServerGroupInstanceTypeC59xlarge  GameServerGroupInstanceType = "c5.9xlarge"
	GameServerGroupInstanceTypeC512xlarge GameServerGroupInstanceType = "c5.12xlarge"
	GameServerGroupInstanceTypeC518xlarge GameServerGroupInstanceType = "c5.18xlarge"
	GameServerGroupInstanceTypeC524xlarge GameServerGroupInstanceType = "c5.24xlarge"
	GameServerGroupInstanceTypeR4Large    GameServerGroupInstanceType = "r4.large"
	GameServerGroupInstanceTypeR4Xlarge   GameServerGroupInstanceType = "r4.xlarge"
	GameServerGroupInstanceTypeR42xlarge  GameServerGroupInstanceType = "r4.2xlarge"
	GameServerGroupInstanceTypeR44xlarge  GameServerGroupInstanceType = "r4.4xlarge"
	GameServerGroupInstanceTypeR48xlarge  GameServerGroupInstanceType = "r4.8xlarge"
	GameServerGroupInstanceTypeR416xlarge GameServerGroupInstanceType = "r4.16xlarge"
	GameServerGroupInstanceTypeR5Large    GameServerGroupInstanceType = "r5.large"
	GameServerGroupInstanceTypeR5Xlarge   GameServerGroupInstanceType = "r5.xlarge"
	GameServerGroupInstanceTypeR52xlarge  GameServerGroupInstanceType = "r5.2xlarge"
	GameServerGroupInstanceTypeR54xlarge  GameServerGroupInstanceType = "r5.4xlarge"
	GameServerGroupInstanceTypeR58xlarge  GameServerGroupInstanceType = "r5.8xlarge"
	GameServerGroupInstanceTypeR512xlarge GameServerGroupInstanceType = "r5.12xlarge"
	GameServerGroupInstanceTypeR516xlarge GameServerGroupInstanceType = "r5.16xlarge"
	GameServerGroupInstanceTypeR524xlarge GameServerGroupInstanceType = "r5.24xlarge"
	GameServerGroupInstanceTypeM4Large    GameServerGroupInstanceType = "m4.large"
	GameServerGroupInstanceTypeM4Xlarge   GameServerGroupInstanceType = "m4.xlarge"
	GameServerGroupInstanceTypeM42xlarge  GameServerGroupInstanceType = "m4.2xlarge"
	GameServerGroupInstanceTypeM44xlarge  GameServerGroupInstanceType = "m4.4xlarge"
	GameServerGroupInstanceTypeM410xlarge GameServerGroupInstanceType = "m4.10xlarge"
	GameServerGroupInstanceTypeM5Large    GameServerGroupInstanceType = "m5.large"
	GameServerGroupInstanceTypeM5Xlarge   GameServerGroupInstanceType = "m5.xlarge"
	GameServerGroupInstanceTypeM52xlarge  GameServerGroupInstanceType = "m5.2xlarge"
	GameServerGroupInstanceTypeM54xlarge  GameServerGroupInstanceType = "m5.4xlarge"
	GameServerGroupInstanceTypeM58xlarge  GameServerGroupInstanceType = "m5.8xlarge"
	GameServerGroupInstanceTypeM512xlarge GameServerGroupInstanceType = "m5.12xlarge"
	GameServerGroupInstanceTypeM516xlarge GameServerGroupInstanceType = "m5.16xlarge"
	GameServerGroupInstanceTypeM524xlarge GameServerGroupInstanceType = "m5.24xlarge"
)

// Values returns all known values for GameServerGroupInstanceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GameServerGroupInstanceType) Values() []GameServerGroupInstanceType {
	return []GameServerGroupInstanceType{
		"c4.large",
		"c4.xlarge",
		"c4.2xlarge",
		"c4.4xlarge",
		"c4.8xlarge",
		"c5.large",
		"c5.xlarge",
		"c5.2xlarge",
		"c5.4xlarge",
		"c5.9xlarge",
		"c5.12xlarge",
		"c5.18xlarge",
		"c5.24xlarge",
		"r4.large",
		"r4.xlarge",
		"r4.2xlarge",
		"r4.4xlarge",
		"r4.8xlarge",
		"r4.16xlarge",
		"r5.large",
		"r5.xlarge",
		"r5.2xlarge",
		"r5.4xlarge",
		"r5.8xlarge",
		"r5.12xlarge",
		"r5.16xlarge",
		"r5.24xlarge",
		"m4.large",
		"m4.xlarge",
		"m4.2xlarge",
		"m4.4xlarge",
		"m4.10xlarge",
		"m5.large",
		"m5.xlarge",
		"m5.2xlarge",
		"m5.4xlarge",
		"m5.8xlarge",
		"m5.12xlarge",
		"m5.16xlarge",
		"m5.24xlarge",
	}
}

type GameServerGroupStatus string

// Enum values for GameServerGroupStatus
const (
	GameServerGroupStatusNew             GameServerGroupStatus = "NEW"
	GameServerGroupStatusActivating      GameServerGroupStatus = "ACTIVATING"
	GameServerGroupStatusActive          GameServerGroupStatus = "ACTIVE"
	GameServerGroupStatusDeleteScheduled GameServerGroupStatus = "DELETE_SCHEDULED"
	GameServerGroupStatusDeleting        GameServerGroupStatus = "DELETING"
	GameServerGroupStatusDeleted         GameServerGroupStatus = "DELETED"
	GameServerGroupStatusError           GameServerGroupStatus = "ERROR"
)

// Values returns all known values for GameServerGroupStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameServerGroupStatus) Values() []GameServerGroupStatus {
	return []GameServerGroupStatus{
		"NEW",
		"ACTIVATING",
		"ACTIVE",
		"DELETE_SCHEDULED",
		"DELETING",
		"DELETED",
		"ERROR",
	}
}

type GameServerHealthCheck string

// Enum values for GameServerHealthCheck
const (
	GameServerHealthCheckHealthy GameServerHealthCheck = "HEALTHY"
)

// Values returns all known values for GameServerHealthCheck. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameServerHealthCheck) Values() []GameServerHealthCheck {
	return []GameServerHealthCheck{
		"HEALTHY",
	}
}

type GameServerInstanceStatus string

// Enum values for GameServerInstanceStatus
const (
	GameServerInstanceStatusActive          GameServerInstanceStatus = "ACTIVE"
	GameServerInstanceStatusDraining        GameServerInstanceStatus = "DRAINING"
	GameServerInstanceStatusSpotTerminating GameServerInstanceStatus = "SPOT_TERMINATING"
)

// Values returns all known values for GameServerInstanceStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameServerInstanceStatus) Values() []GameServerInstanceStatus {
	return []GameServerInstanceStatus{
		"ACTIVE",
		"DRAINING",
		"SPOT_TERMINATING",
	}
}

type GameServerProtectionPolicy string

// Enum values for GameServerProtectionPolicy
const (
	GameServerProtectionPolicyNoProtection   GameServerProtectionPolicy = "NO_PROTECTION"
	GameServerProtectionPolicyFullProtection GameServerProtectionPolicy = "FULL_PROTECTION"
)

// Values returns all known values for GameServerProtectionPolicy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GameServerProtectionPolicy) Values() []GameServerProtectionPolicy {
	return []GameServerProtectionPolicy{
		"NO_PROTECTION",
		"FULL_PROTECTION",
	}
}

type GameServerUtilizationStatus string

// Enum values for GameServerUtilizationStatus
const (
	GameServerUtilizationStatusAvailable GameServerUtilizationStatus = "AVAILABLE"
	GameServerUtilizationStatusUtilized  GameServerUtilizationStatus = "UTILIZED"
)

// Values returns all known values for GameServerUtilizationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GameServerUtilizationStatus) Values() []GameServerUtilizationStatus {
	return []GameServerUtilizationStatus{
		"AVAILABLE",
		"UTILIZED",
	}
}

type GameSessionPlacementState string

// Enum values for GameSessionPlacementState
const (
	GameSessionPlacementStatePending   GameSessionPlacementState = "PENDING"
	GameSessionPlacementStateFulfilled GameSessionPlacementState = "FULFILLED"
	GameSessionPlacementStateCancelled GameSessionPlacementState = "CANCELLED"
	GameSessionPlacementStateTimedOut  GameSessionPlacementState = "TIMED_OUT"
	GameSessionPlacementStateFailed    GameSessionPlacementState = "FAILED"
)

// Values returns all known values for GameSessionPlacementState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GameSessionPlacementState) Values() []GameSessionPlacementState {
	return []GameSessionPlacementState{
		"PENDING",
		"FULFILLED",
		"CANCELLED",
		"TIMED_OUT",
		"FAILED",
	}
}

type GameSessionStatus string

// Enum values for GameSessionStatus
const (
	GameSessionStatusActive      GameSessionStatus = "ACTIVE"
	GameSessionStatusActivating  GameSessionStatus = "ACTIVATING"
	GameSessionStatusTerminated  GameSessionStatus = "TERMINATED"
	GameSessionStatusTerminating GameSessionStatus = "TERMINATING"
	GameSessionStatusError       GameSessionStatus = "ERROR"
)

// Values returns all known values for GameSessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameSessionStatus) Values() []GameSessionStatus {
	return []GameSessionStatus{
		"ACTIVE",
		"ACTIVATING",
		"TERMINATED",
		"TERMINATING",
		"ERROR",
	}
}

type GameSessionStatusReason string

// Enum values for GameSessionStatusReason
const (
	GameSessionStatusReasonInterrupted GameSessionStatusReason = "INTERRUPTED"
)

// Values returns all known values for GameSessionStatusReason. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GameSessionStatusReason) Values() []GameSessionStatusReason {
	return []GameSessionStatusReason{
		"INTERRUPTED",
	}
}

type InstanceStatus string

// Enum values for InstanceStatus
const (
	InstanceStatusPending     InstanceStatus = "PENDING"
	InstanceStatusActive      InstanceStatus = "ACTIVE"
	InstanceStatusTerminating InstanceStatus = "TERMINATING"
)

// Values returns all known values for InstanceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceStatus) Values() []InstanceStatus {
	return []InstanceStatus{
		"PENDING",
		"ACTIVE",
		"TERMINATING",
	}
}

type IpProtocol string

// Enum values for IpProtocol
const (
	IpProtocolTcp IpProtocol = "TCP"
	IpProtocolUdp IpProtocol = "UDP"
)

// Values returns all known values for IpProtocol. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IpProtocol) Values() []IpProtocol {
	return []IpProtocol{
		"TCP",
		"UDP",
	}
}

type MatchmakingConfigurationStatus string

// Enum values for MatchmakingConfigurationStatus
const (
	MatchmakingConfigurationStatusCancelled          MatchmakingConfigurationStatus = "CANCELLED"
	MatchmakingConfigurationStatusCompleted          MatchmakingConfigurationStatus = "COMPLETED"
	MatchmakingConfigurationStatusFailed             MatchmakingConfigurationStatus = "FAILED"
	MatchmakingConfigurationStatusPlacing            MatchmakingConfigurationStatus = "PLACING"
	MatchmakingConfigurationStatusQueued             MatchmakingConfigurationStatus = "QUEUED"
	MatchmakingConfigurationStatusRequiresAcceptance MatchmakingConfigurationStatus = "REQUIRES_ACCEPTANCE"
	MatchmakingConfigurationStatusSearching          MatchmakingConfigurationStatus = "SEARCHING"
	MatchmakingConfigurationStatusTimedOut           MatchmakingConfigurationStatus = "TIMED_OUT"
)

// Values returns all known values for MatchmakingConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (MatchmakingConfigurationStatus) Values() []MatchmakingConfigurationStatus {
	return []MatchmakingConfigurationStatus{
		"CANCELLED",
		"COMPLETED",
		"FAILED",
		"PLACING",
		"QUEUED",
		"REQUIRES_ACCEPTANCE",
		"SEARCHING",
		"TIMED_OUT",
	}
}

type MetricName string

// Enum values for MetricName
const (
	MetricNameActivatinggamesessions       MetricName = "ActivatingGameSessions"
	MetricNameActivegamesessions           MetricName = "ActiveGameSessions"
	MetricNameActiveinstances              MetricName = "ActiveInstances"
	MetricNameAvailablegamesessions        MetricName = "AvailableGameSessions"
	MetricNameAvailableplayersessions      MetricName = "AvailablePlayerSessions"
	MetricNameCurrentplayersessions        MetricName = "CurrentPlayerSessions"
	MetricNameIdleinstances                MetricName = "IdleInstances"
	MetricNamePercentavailablegamesessions MetricName = "PercentAvailableGameSessions"
	MetricNamePercentidleinstances         MetricName = "PercentIdleInstances"
	MetricNameQueuedepth                   MetricName = "QueueDepth"
	MetricNameWaittime                     MetricName = "WaitTime"
)

// Values returns all known values for MetricName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MetricName) Values() []MetricName {
	return []MetricName{
		"ActivatingGameSessions",
		"ActiveGameSessions",
		"ActiveInstances",
		"AvailableGameSessions",
		"AvailablePlayerSessions",
		"CurrentPlayerSessions",
		"IdleInstances",
		"PercentAvailableGameSessions",
		"PercentIdleInstances",
		"QueueDepth",
		"WaitTime",
	}
}

type OperatingSystem string

// Enum values for OperatingSystem
const (
	OperatingSystemWindows2012  OperatingSystem = "WINDOWS_2012"
	OperatingSystemAmazonLinux  OperatingSystem = "AMAZON_LINUX"
	OperatingSystemAmazonLinux2 OperatingSystem = "AMAZON_LINUX_2"
)

// Values returns all known values for OperatingSystem. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperatingSystem) Values() []OperatingSystem {
	return []OperatingSystem{
		"WINDOWS_2012",
		"AMAZON_LINUX",
		"AMAZON_LINUX_2",
	}
}

type PlayerSessionCreationPolicy string

// Enum values for PlayerSessionCreationPolicy
const (
	PlayerSessionCreationPolicyAcceptAll PlayerSessionCreationPolicy = "ACCEPT_ALL"
	PlayerSessionCreationPolicyDenyAll   PlayerSessionCreationPolicy = "DENY_ALL"
)

// Values returns all known values for PlayerSessionCreationPolicy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PlayerSessionCreationPolicy) Values() []PlayerSessionCreationPolicy {
	return []PlayerSessionCreationPolicy{
		"ACCEPT_ALL",
		"DENY_ALL",
	}
}

type PlayerSessionStatus string

// Enum values for PlayerSessionStatus
const (
	PlayerSessionStatusReserved  PlayerSessionStatus = "RESERVED"
	PlayerSessionStatusActive    PlayerSessionStatus = "ACTIVE"
	PlayerSessionStatusCompleted PlayerSessionStatus = "COMPLETED"
	PlayerSessionStatusTimedout  PlayerSessionStatus = "TIMEDOUT"
)

// Values returns all known values for PlayerSessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlayerSessionStatus) Values() []PlayerSessionStatus {
	return []PlayerSessionStatus{
		"RESERVED",
		"ACTIVE",
		"COMPLETED",
		"TIMEDOUT",
	}
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeRulebased   PolicyType = "RuleBased"
	PolicyTypeTargetbased PolicyType = "TargetBased"
)

// Values returns all known values for PolicyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PolicyType) Values() []PolicyType {
	return []PolicyType{
		"RuleBased",
		"TargetBased",
	}
}

type ProtectionPolicy string

// Enum values for ProtectionPolicy
const (
	ProtectionPolicyNoprotection   ProtectionPolicy = "NoProtection"
	ProtectionPolicyFullprotection ProtectionPolicy = "FullProtection"
)

// Values returns all known values for ProtectionPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProtectionPolicy) Values() []ProtectionPolicy {
	return []ProtectionPolicy{
		"NoProtection",
		"FullProtection",
	}
}

type RoutingStrategyType string

// Enum values for RoutingStrategyType
const (
	RoutingStrategyTypeSimple   RoutingStrategyType = "SIMPLE"
	RoutingStrategyTypeTerminal RoutingStrategyType = "TERMINAL"
)

// Values returns all known values for RoutingStrategyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RoutingStrategyType) Values() []RoutingStrategyType {
	return []RoutingStrategyType{
		"SIMPLE",
		"TERMINAL",
	}
}

type ScalingAdjustmentType string

// Enum values for ScalingAdjustmentType
const (
	ScalingAdjustmentTypeChangeincapacity        ScalingAdjustmentType = "ChangeInCapacity"
	ScalingAdjustmentTypeExactcapacity           ScalingAdjustmentType = "ExactCapacity"
	ScalingAdjustmentTypePercentchangeincapacity ScalingAdjustmentType = "PercentChangeInCapacity"
)

// Values returns all known values for ScalingAdjustmentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScalingAdjustmentType) Values() []ScalingAdjustmentType {
	return []ScalingAdjustmentType{
		"ChangeInCapacity",
		"ExactCapacity",
		"PercentChangeInCapacity",
	}
}

type ScalingStatusType string

// Enum values for ScalingStatusType
const (
	ScalingStatusTypeActive          ScalingStatusType = "ACTIVE"
	ScalingStatusTypeUpdateRequested ScalingStatusType = "UPDATE_REQUESTED"
	ScalingStatusTypeUpdating        ScalingStatusType = "UPDATING"
	ScalingStatusTypeDeleteRequested ScalingStatusType = "DELETE_REQUESTED"
	ScalingStatusTypeDeleting        ScalingStatusType = "DELETING"
	ScalingStatusTypeDeleted         ScalingStatusType = "DELETED"
	ScalingStatusTypeError           ScalingStatusType = "ERROR"
)

// Values returns all known values for ScalingStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScalingStatusType) Values() []ScalingStatusType {
	return []ScalingStatusType{
		"ACTIVE",
		"UPDATE_REQUESTED",
		"UPDATING",
		"DELETE_REQUESTED",
		"DELETING",
		"DELETED",
		"ERROR",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

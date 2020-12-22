// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified CIDR IP or Amazon EC2 security group isn't authorized for the
// specified security group. Amazon DocumentDB also might not be authorized to
// perform necessary actions on your behalf using IAM.
type AuthorizationNotFoundFault struct {
	Message *string
}

func (e *AuthorizationNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationNotFoundFault) ErrorCode() string             { return "AuthorizationNotFoundFault" }
func (e *AuthorizationNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// CertificateIdentifier doesn't refer to an existing certificate.
type CertificateNotFoundFault struct {
	Message *string
}

func (e *CertificateNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateNotFoundFault) ErrorCode() string             { return "CertificateNotFoundFault" }
func (e *CertificateNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a cluster with the given identifier.
type DBClusterAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterAlreadyExistsFault) ErrorCode() string             { return "DBClusterAlreadyExistsFault" }
func (e *DBClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBClusterIdentifier doesn't refer to an existing cluster.
type DBClusterNotFoundFault struct {
	Message *string
}

func (e *DBClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterNotFoundFault) ErrorCode() string             { return "DBClusterNotFoundFault" }
func (e *DBClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBClusterParameterGroupName doesn't refer to an existing cluster parameter
// group.
type DBClusterParameterGroupNotFoundFault struct {
	Message *string
}

func (e *DBClusterParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorCode() string {
	return "DBClusterParameterGroupNotFoundFault"
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The cluster can't be created because you have reached the maximum allowed quota
// of clusters.
type DBClusterQuotaExceededFault struct {
	Message *string
}

func (e *DBClusterQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterQuotaExceededFault) ErrorCode() string             { return "DBClusterQuotaExceededFault" }
func (e *DBClusterQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a cluster snapshot with the given identifier.
type DBClusterSnapshotAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorCode() string {
	return "DBClusterSnapshotAlreadyExistsFault"
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBClusterSnapshotIdentifier doesn't refer to an existing cluster snapshot.
type DBClusterSnapshotNotFoundFault struct {
	Message *string
}

func (e *DBClusterSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotNotFoundFault) ErrorCode() string             { return "DBClusterSnapshotNotFoundFault" }
func (e *DBClusterSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a instance with the given identifier.
type DBInstanceAlreadyExistsFault struct {
	Message *string
}

func (e *DBInstanceAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceAlreadyExistsFault) ErrorCode() string             { return "DBInstanceAlreadyExistsFault" }
func (e *DBInstanceAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBInstanceIdentifier doesn't refer to an existing instance.
type DBInstanceNotFoundFault struct {
	Message *string
}

func (e *DBInstanceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceNotFoundFault) ErrorCode() string             { return "DBInstanceNotFoundFault" }
func (e *DBInstanceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A parameter group with the same name already exists.
type DBParameterGroupAlreadyExistsFault struct {
	Message *string
}

func (e *DBParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorCode() string {
	return "DBParameterGroupAlreadyExistsFault"
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBParameterGroupName doesn't refer to an existing parameter group.
type DBParameterGroupNotFoundFault struct {
	Message *string
}

func (e *DBParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupNotFoundFault) ErrorCode() string             { return "DBParameterGroupNotFoundFault" }
func (e *DBParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This request would cause you to exceed the allowed number of parameter groups.
type DBParameterGroupQuotaExceededFault struct {
	Message *string
}

func (e *DBParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupQuotaExceededFault) ErrorCode() string {
	return "DBParameterGroupQuotaExceededFault"
}
func (e *DBParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSecurityGroupName doesn't refer to an existing security group.
type DBSecurityGroupNotFoundFault struct {
	Message *string
}

func (e *DBSecurityGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSecurityGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSecurityGroupNotFoundFault) ErrorCode() string             { return "DBSecurityGroupNotFoundFault" }
func (e *DBSecurityGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier is already being used by an existing snapshot.
type DBSnapshotAlreadyExistsFault struct {
	Message *string
}

func (e *DBSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotAlreadyExistsFault) ErrorCode() string             { return "DBSnapshotAlreadyExistsFault" }
func (e *DBSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier doesn't refer to an existing snapshot.
type DBSnapshotNotFoundFault struct {
	Message *string
}

func (e *DBSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotNotFoundFault) ErrorCode() string             { return "DBSnapshotNotFoundFault" }
func (e *DBSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSubnetGroupName is already being used by an existing subnet group.
type DBSubnetGroupAlreadyExistsFault struct {
	Message *string
}

func (e *DBSubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorCode() string {
	return "DBSubnetGroupAlreadyExistsFault"
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Subnets in the subnet group should cover at least two Availability Zones unless
// there is only one Availability Zone.
type DBSubnetGroupDoesNotCoverEnoughAZs struct {
	Message *string
}

func (e *DBSubnetGroupDoesNotCoverEnoughAZs) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorCode() string {
	return "DBSubnetGroupDoesNotCoverEnoughAZs"
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSubnetGroupName doesn't refer to an existing subnet group.
type DBSubnetGroupNotFoundFault struct {
	Message *string
}

func (e *DBSubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupNotFoundFault) ErrorCode() string             { return "DBSubnetGroupNotFoundFault" }
func (e *DBSubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of subnet groups.
type DBSubnetGroupQuotaExceededFault struct {
	Message *string
}

func (e *DBSubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorCode() string {
	return "DBSubnetGroupQuotaExceededFault"
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of subnets in a subnet
// group.
type DBSubnetQuotaExceededFault struct {
	Message *string
}

func (e *DBSubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetQuotaExceededFault) ErrorCode() string             { return "DBSubnetQuotaExceededFault" }
func (e *DBSubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The upgrade failed because a resource that the depends on can't be modified.
type DBUpgradeDependencyFailureFault struct {
	Message *string
}

func (e *DBUpgradeDependencyFailureFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBUpgradeDependencyFailureFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBUpgradeDependencyFailureFault) ErrorCode() string {
	return "DBUpgradeDependencyFailureFault"
}
func (e *DBUpgradeDependencyFailureFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of instances.
type InstanceQuotaExceededFault struct {
	Message *string
}

func (e *InstanceQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InstanceQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InstanceQuotaExceededFault) ErrorCode() string             { return "InstanceQuotaExceededFault" }
func (e *InstanceQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The cluster doesn't have enough capacity for the current operation.
type InsufficientDBClusterCapacityFault struct {
	Message *string
}

func (e *InsufficientDBClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBClusterCapacityFault) ErrorCode() string {
	return "InsufficientDBClusterCapacityFault"
}
func (e *InsufficientDBClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified instance class isn't available in the specified Availability Zone.
type InsufficientDBInstanceCapacityFault struct {
	Message *string
}

func (e *InsufficientDBInstanceCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBInstanceCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBInstanceCapacityFault) ErrorCode() string {
	return "InsufficientDBInstanceCapacityFault"
}
func (e *InsufficientDBInstanceCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There is not enough storage available for the current action. You might be able
// to resolve this error by updating your subnet group to use different
// Availability Zones that have more storage available.
type InsufficientStorageClusterCapacityFault struct {
	Message *string
}

func (e *InsufficientStorageClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientStorageClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientStorageClusterCapacityFault) ErrorCode() string {
	return "InsufficientStorageClusterCapacityFault"
}
func (e *InsufficientStorageClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The provided value isn't a valid cluster snapshot state.
type InvalidDBClusterSnapshotStateFault struct {
	Message *string
}

func (e *InvalidDBClusterSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorCode() string {
	return "InvalidDBClusterSnapshotStateFault"
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The cluster isn't in a valid state.
type InvalidDBClusterStateFault struct {
	Message *string
}

func (e *InvalidDBClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterStateFault) ErrorCode() string             { return "InvalidDBClusterStateFault" }
func (e *InvalidDBClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified instance isn't in the available state.
type InvalidDBInstanceStateFault struct {
	Message *string
}

func (e *InvalidDBInstanceStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBInstanceStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBInstanceStateFault) ErrorCode() string             { return "InvalidDBInstanceStateFault" }
func (e *InvalidDBInstanceStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The parameter group is in use, or it is in a state that is not valid. If you are
// trying to delete the parameter group, you can't delete it when the parameter
// group is in this state.
type InvalidDBParameterGroupStateFault struct {
	Message *string
}

func (e *InvalidDBParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBParameterGroupStateFault) ErrorCode() string {
	return "InvalidDBParameterGroupStateFault"
}
func (e *InvalidDBParameterGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the security group doesn't allow deletion.
type InvalidDBSecurityGroupStateFault struct {
	Message *string
}

func (e *InvalidDBSecurityGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSecurityGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSecurityGroupStateFault) ErrorCode() string {
	return "InvalidDBSecurityGroupStateFault"
}
func (e *InvalidDBSecurityGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the snapshot doesn't allow deletion.
type InvalidDBSnapshotStateFault struct {
	Message *string
}

func (e *InvalidDBSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSnapshotStateFault) ErrorCode() string             { return "InvalidDBSnapshotStateFault" }
func (e *InvalidDBSnapshotStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet group can't be deleted because it's in use.
type InvalidDBSubnetGroupStateFault struct {
	Message *string
}

func (e *InvalidDBSubnetGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetGroupStateFault) ErrorCode() string             { return "InvalidDBSubnetGroupStateFault" }
func (e *InvalidDBSubnetGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet isn't in the available state.
type InvalidDBSubnetStateFault struct {
	Message *string
}

func (e *InvalidDBSubnetStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetStateFault) ErrorCode() string             { return "InvalidDBSubnetStateFault" }
func (e *InvalidDBSubnetStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You cannot restore from a virtual private cloud (VPC) backup to a non-VPC DB
// instance.
type InvalidRestoreFault struct {
	Message *string
}

func (e *InvalidRestoreFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRestoreFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRestoreFault) ErrorCode() string             { return "InvalidRestoreFault" }
func (e *InvalidRestoreFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet is not valid, or multiple subnets were requested that are
// not all in a common virtual private cloud (VPC).
type InvalidSubnet struct {
	Message *string
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string             { return "InvalidSubnet" }
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet group doesn't cover all Availability Zones after it is created
// because of changes that were made.
type InvalidVPCNetworkStateFault struct {
	Message *string
}

func (e *InvalidVPCNetworkStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCNetworkStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCNetworkStateFault) ErrorCode() string             { return "InvalidVPCNetworkStateFault" }
func (e *InvalidVPCNetworkStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred when accessing an AWS KMS key.
type KMSKeyNotAccessibleFault struct {
	Message *string
}

func (e *KMSKeyNotAccessibleFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSKeyNotAccessibleFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSKeyNotAccessibleFault) ErrorCode() string             { return "KMSKeyNotAccessibleFault" }
func (e *KMSKeyNotAccessibleFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource ID was not found.
type ResourceNotFoundFault struct {
	Message *string
}

func (e *ResourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundFault) ErrorCode() string             { return "ResourceNotFoundFault" }
func (e *ResourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of accounts that you can share a manual DB
// snapshot with.
type SharedSnapshotQuotaExceededFault struct {
	Message *string
}

func (e *SharedSnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SharedSnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SharedSnapshotQuotaExceededFault) ErrorCode() string {
	return "SharedSnapshotQuotaExceededFault"
}
func (e *SharedSnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of snapshots.
type SnapshotQuotaExceededFault struct {
	Message *string
}

func (e *SnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotQuotaExceededFault) ErrorCode() string             { return "SnapshotQuotaExceededFault" }
func (e *SnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed amount of storage available
// across all instances.
type StorageQuotaExceededFault struct {
	Message *string
}

func (e *StorageQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageQuotaExceededFault) ErrorCode() string             { return "StorageQuotaExceededFault" }
func (e *StorageQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Storage of the specified StorageType can't be associated with the DB instance.
type StorageTypeNotSupportedFault struct {
	Message *string
}

func (e *StorageTypeNotSupportedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageTypeNotSupportedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageTypeNotSupportedFault) ErrorCode() string             { return "StorageTypeNotSupportedFault" }
func (e *StorageTypeNotSupportedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet is already in use in the Availability Zone.
type SubnetAlreadyInUse struct {
	Message *string
}

func (e *SubnetAlreadyInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetAlreadyInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetAlreadyInUse) ErrorCode() string             { return "SubnetAlreadyInUse" }
func (e *SubnetAlreadyInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

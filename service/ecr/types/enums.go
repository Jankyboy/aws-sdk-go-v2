// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeAes256 EncryptionType = "AES256"
	EncryptionTypeKms    EncryptionType = "KMS"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"AES256",
		"KMS",
	}
}

type FindingSeverity string

// Enum values for FindingSeverity
const (
	FindingSeverityInformational FindingSeverity = "INFORMATIONAL"
	FindingSeverityLow           FindingSeverity = "LOW"
	FindingSeverityMedium        FindingSeverity = "MEDIUM"
	FindingSeverityHigh          FindingSeverity = "HIGH"
	FindingSeverityCritical      FindingSeverity = "CRITICAL"
	FindingSeverityUndefined     FindingSeverity = "UNDEFINED"
)

// Values returns all known values for FindingSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingSeverity) Values() []FindingSeverity {
	return []FindingSeverity{
		"INFORMATIONAL",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
		"UNDEFINED",
	}
}

type ImageActionType string

// Enum values for ImageActionType
const (
	ImageActionTypeExpire ImageActionType = "EXPIRE"
)

// Values returns all known values for ImageActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageActionType) Values() []ImageActionType {
	return []ImageActionType{
		"EXPIRE",
	}
}

type ImageFailureCode string

// Enum values for ImageFailureCode
const (
	ImageFailureCodeInvalidimagedigest            ImageFailureCode = "InvalidImageDigest"
	ImageFailureCodeInvalidimagetag               ImageFailureCode = "InvalidImageTag"
	ImageFailureCodeImagetagdoesnotmatchdigest    ImageFailureCode = "ImageTagDoesNotMatchDigest"
	ImageFailureCodeImagenotfound                 ImageFailureCode = "ImageNotFound"
	ImageFailureCodeMissingdigestandtag           ImageFailureCode = "MissingDigestAndTag"
	ImageFailureCodeImagereferencedbymanifestlist ImageFailureCode = "ImageReferencedByManifestList"
	ImageFailureCodeKmserror                      ImageFailureCode = "KmsError"
)

// Values returns all known values for ImageFailureCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageFailureCode) Values() []ImageFailureCode {
	return []ImageFailureCode{
		"InvalidImageDigest",
		"InvalidImageTag",
		"ImageTagDoesNotMatchDigest",
		"ImageNotFound",
		"MissingDigestAndTag",
		"ImageReferencedByManifestList",
		"KmsError",
	}
}

type ImageTagMutability string

// Enum values for ImageTagMutability
const (
	ImageTagMutabilityMutable   ImageTagMutability = "MUTABLE"
	ImageTagMutabilityImmutable ImageTagMutability = "IMMUTABLE"
)

// Values returns all known values for ImageTagMutability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageTagMutability) Values() []ImageTagMutability {
	return []ImageTagMutability{
		"MUTABLE",
		"IMMUTABLE",
	}
}

type LayerAvailability string

// Enum values for LayerAvailability
const (
	LayerAvailabilityAvailable   LayerAvailability = "AVAILABLE"
	LayerAvailabilityUnavailable LayerAvailability = "UNAVAILABLE"
)

// Values returns all known values for LayerAvailability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LayerAvailability) Values() []LayerAvailability {
	return []LayerAvailability{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

type LayerFailureCode string

// Enum values for LayerFailureCode
const (
	LayerFailureCodeInvalidlayerdigest LayerFailureCode = "InvalidLayerDigest"
	LayerFailureCodeMissinglayerdigest LayerFailureCode = "MissingLayerDigest"
)

// Values returns all known values for LayerFailureCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LayerFailureCode) Values() []LayerFailureCode {
	return []LayerFailureCode{
		"InvalidLayerDigest",
		"MissingLayerDigest",
	}
}

type LifecyclePolicyPreviewStatus string

// Enum values for LifecyclePolicyPreviewStatus
const (
	LifecyclePolicyPreviewStatusInProgress LifecyclePolicyPreviewStatus = "IN_PROGRESS"
	LifecyclePolicyPreviewStatusComplete   LifecyclePolicyPreviewStatus = "COMPLETE"
	LifecyclePolicyPreviewStatusExpired    LifecyclePolicyPreviewStatus = "EXPIRED"
	LifecyclePolicyPreviewStatusFailed     LifecyclePolicyPreviewStatus = "FAILED"
)

// Values returns all known values for LifecyclePolicyPreviewStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecyclePolicyPreviewStatus) Values() []LifecyclePolicyPreviewStatus {
	return []LifecyclePolicyPreviewStatus{
		"IN_PROGRESS",
		"COMPLETE",
		"EXPIRED",
		"FAILED",
	}
}

type ScanStatus string

// Enum values for ScanStatus
const (
	ScanStatusInProgress ScanStatus = "IN_PROGRESS"
	ScanStatusComplete   ScanStatus = "COMPLETE"
	ScanStatusFailed     ScanStatus = "FAILED"
)

// Values returns all known values for ScanStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScanStatus) Values() []ScanStatus {
	return []ScanStatus{
		"IN_PROGRESS",
		"COMPLETE",
		"FAILED",
	}
}

type TagStatus string

// Enum values for TagStatus
const (
	TagStatusTagged   TagStatus = "TAGGED"
	TagStatusUntagged TagStatus = "UNTAGGED"
	TagStatusAny      TagStatus = "ANY"
)

// Values returns all known values for TagStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TagStatus) Values() []TagStatus {
	return []TagStatus{
		"TAGGED",
		"UNTAGGED",
		"ANY",
	}
}

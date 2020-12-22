// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ClipFragmentSelectorType string

// Enum values for ClipFragmentSelectorType
const (
	ClipFragmentSelectorTypeProducerTimestamp ClipFragmentSelectorType = "PRODUCER_TIMESTAMP"
	ClipFragmentSelectorTypeServerTimestamp   ClipFragmentSelectorType = "SERVER_TIMESTAMP"
)

// Values returns all known values for ClipFragmentSelectorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClipFragmentSelectorType) Values() []ClipFragmentSelectorType {
	return []ClipFragmentSelectorType{
		"PRODUCER_TIMESTAMP",
		"SERVER_TIMESTAMP",
	}
}

type ContainerFormat string

// Enum values for ContainerFormat
const (
	ContainerFormatFragmentedMp4 ContainerFormat = "FRAGMENTED_MP4"
	ContainerFormatMpegTs        ContainerFormat = "MPEG_TS"
)

// Values returns all known values for ContainerFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerFormat) Values() []ContainerFormat {
	return []ContainerFormat{
		"FRAGMENTED_MP4",
		"MPEG_TS",
	}
}

type DASHDisplayFragmentNumber string

// Enum values for DASHDisplayFragmentNumber
const (
	DASHDisplayFragmentNumberAlways DASHDisplayFragmentNumber = "ALWAYS"
	DASHDisplayFragmentNumberNever  DASHDisplayFragmentNumber = "NEVER"
)

// Values returns all known values for DASHDisplayFragmentNumber. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DASHDisplayFragmentNumber) Values() []DASHDisplayFragmentNumber {
	return []DASHDisplayFragmentNumber{
		"ALWAYS",
		"NEVER",
	}
}

type DASHDisplayFragmentTimestamp string

// Enum values for DASHDisplayFragmentTimestamp
const (
	DASHDisplayFragmentTimestampAlways DASHDisplayFragmentTimestamp = "ALWAYS"
	DASHDisplayFragmentTimestampNever  DASHDisplayFragmentTimestamp = "NEVER"
)

// Values returns all known values for DASHDisplayFragmentTimestamp. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DASHDisplayFragmentTimestamp) Values() []DASHDisplayFragmentTimestamp {
	return []DASHDisplayFragmentTimestamp{
		"ALWAYS",
		"NEVER",
	}
}

type DASHFragmentSelectorType string

// Enum values for DASHFragmentSelectorType
const (
	DASHFragmentSelectorTypeProducerTimestamp DASHFragmentSelectorType = "PRODUCER_TIMESTAMP"
	DASHFragmentSelectorTypeServerTimestamp   DASHFragmentSelectorType = "SERVER_TIMESTAMP"
)

// Values returns all known values for DASHFragmentSelectorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DASHFragmentSelectorType) Values() []DASHFragmentSelectorType {
	return []DASHFragmentSelectorType{
		"PRODUCER_TIMESTAMP",
		"SERVER_TIMESTAMP",
	}
}

type DASHPlaybackMode string

// Enum values for DASHPlaybackMode
const (
	DASHPlaybackModeLive       DASHPlaybackMode = "LIVE"
	DASHPlaybackModeLiveReplay DASHPlaybackMode = "LIVE_REPLAY"
	DASHPlaybackModeOnDemand   DASHPlaybackMode = "ON_DEMAND"
)

// Values returns all known values for DASHPlaybackMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DASHPlaybackMode) Values() []DASHPlaybackMode {
	return []DASHPlaybackMode{
		"LIVE",
		"LIVE_REPLAY",
		"ON_DEMAND",
	}
}

type FragmentSelectorType string

// Enum values for FragmentSelectorType
const (
	FragmentSelectorTypeProducerTimestamp FragmentSelectorType = "PRODUCER_TIMESTAMP"
	FragmentSelectorTypeServerTimestamp   FragmentSelectorType = "SERVER_TIMESTAMP"
)

// Values returns all known values for FragmentSelectorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FragmentSelectorType) Values() []FragmentSelectorType {
	return []FragmentSelectorType{
		"PRODUCER_TIMESTAMP",
		"SERVER_TIMESTAMP",
	}
}

type HLSDiscontinuityMode string

// Enum values for HLSDiscontinuityMode
const (
	HLSDiscontinuityModeAlways          HLSDiscontinuityMode = "ALWAYS"
	HLSDiscontinuityModeNever           HLSDiscontinuityMode = "NEVER"
	HLSDiscontinuityModeOnDiscontinuity HLSDiscontinuityMode = "ON_DISCONTINUITY"
)

// Values returns all known values for HLSDiscontinuityMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HLSDiscontinuityMode) Values() []HLSDiscontinuityMode {
	return []HLSDiscontinuityMode{
		"ALWAYS",
		"NEVER",
		"ON_DISCONTINUITY",
	}
}

type HLSDisplayFragmentTimestamp string

// Enum values for HLSDisplayFragmentTimestamp
const (
	HLSDisplayFragmentTimestampAlways HLSDisplayFragmentTimestamp = "ALWAYS"
	HLSDisplayFragmentTimestampNever  HLSDisplayFragmentTimestamp = "NEVER"
)

// Values returns all known values for HLSDisplayFragmentTimestamp. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (HLSDisplayFragmentTimestamp) Values() []HLSDisplayFragmentTimestamp {
	return []HLSDisplayFragmentTimestamp{
		"ALWAYS",
		"NEVER",
	}
}

type HLSFragmentSelectorType string

// Enum values for HLSFragmentSelectorType
const (
	HLSFragmentSelectorTypeProducerTimestamp HLSFragmentSelectorType = "PRODUCER_TIMESTAMP"
	HLSFragmentSelectorTypeServerTimestamp   HLSFragmentSelectorType = "SERVER_TIMESTAMP"
)

// Values returns all known values for HLSFragmentSelectorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HLSFragmentSelectorType) Values() []HLSFragmentSelectorType {
	return []HLSFragmentSelectorType{
		"PRODUCER_TIMESTAMP",
		"SERVER_TIMESTAMP",
	}
}

type HLSPlaybackMode string

// Enum values for HLSPlaybackMode
const (
	HLSPlaybackModeLive       HLSPlaybackMode = "LIVE"
	HLSPlaybackModeLiveReplay HLSPlaybackMode = "LIVE_REPLAY"
	HLSPlaybackModeOnDemand   HLSPlaybackMode = "ON_DEMAND"
)

// Values returns all known values for HLSPlaybackMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HLSPlaybackMode) Values() []HLSPlaybackMode {
	return []HLSPlaybackMode{
		"LIVE",
		"LIVE_REPLAY",
		"ON_DEMAND",
	}
}

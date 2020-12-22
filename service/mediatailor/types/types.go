// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The configuration for Ad Marker Passthrough. Ad marker passthrough can be used
// to pass ad markers from the origin to the customized manifest.
type AdMarkerPassthrough struct {

	// For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN,
	// EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest
	// to the MediaTailor personalized manifest.No logic is applied to these ad
	// markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled
	// for that ad break, MediaTailor will not set the value to 0.
	Enabled bool
}

// The configuration for Avail Suppression. Ad suppression can be used to turn off
// ad personalization in a long manifest, or if a viewer joins mid-break.
type AvailSuppression struct {

	// Sets the mode for avail suppression, also known as ad suppression. By default,
	// ad suppression is off and all ad breaks are filled by MediaTailor with ads or
	// slate.
	Mode Mode

	// The avail suppression value is a live edge offset time in HH:MM:SS. MediaTailor
	// won't fill ad breaks on or behind this time in the manifest lookback window.
	Value *string
}

// The configuration for bumpers. Bumpers are short audio or video clips that play
// at the start or before the end of an ad break.
type Bumper struct {

	// The URL for the end bumper asset.
	EndUrl *string

	// The URL for the start bumper asset.
	StartUrl *string
}

// The configuration for using a content delivery network (CDN), like Amazon
// CloudFront, for content and ad segment management.
type CdnConfiguration struct {

	// A non-default content delivery network (CDN) to serve ad segments. By default,
	// AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as
	// its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN
	// for the following origin: ads.mediatailor.<region>.amazonaws.com. Then specify
	// the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor
	// serves a manifest, it reports your CDN as the source for ad segments.
	AdSegmentUrlPrefix *string

	// A content delivery network (CDN) to cache content segments, so that content
	// requests don’t always have to go to the origin server. First, create a rule in
	// your CDN for the content segment origin server. Then specify the rule's name in
	// this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest,
	// it reports your CDN as the source for content segments.
	ContentSegmentUrlPrefix *string
}

// The configuration for DASH content.
type DashConfiguration struct {

	// The URL generated by MediaTailor to initiate a playback session. The session
	// uses server-side reporting. This setting is ignored in PUT operations.
	ManifestEndpointPrefix *string

	// The setting that controls whether MediaTailor includes the Location tag in DASH
	// manifests. MediaTailor populates the Location tag with the URL for manifest
	// update requests, to be used by players that don't support sticky redirects.
	// Disable this if you have CDN routing rules set up for accessing MediaTailor
	// manifests, and you are either using client-side reporting or your players
	// support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The
	// EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
	MpdLocation *string

	// The setting that controls whether MediaTailor handles manifests from the origin
	// server as multi-period manifests or single-period manifests. If your origin
	// server produces single-period manifests, set this to SINGLE_PERIOD. The default
	// setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it
	// to MULTI_PERIOD.
	OriginManifestType OriginManifestType
}

// The configuration for DASH PUT operations.
type DashConfigurationForPut struct {

	// The setting that controls whether MediaTailor includes the Location tag in DASH
	// manifests. MediaTailor populates the Location tag with the URL for manifest
	// update requests, to be used by players that don't support sticky redirects.
	// Disable this if you have CDN routing rules set up for accessing MediaTailor
	// manifests, and you are either using client-side reporting or your players
	// support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The
	// EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
	MpdLocation *string

	// The setting that controls whether MediaTailor handles manifests from the origin
	// server as multi-period manifests or single-period manifests. If your origin
	// server produces single-period manifests, set this to SINGLE_PERIOD. The default
	// setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it
	// to MULTI_PERIOD.
	OriginManifestType OriginManifestType
}

// The configuration for HLS content.
type HlsConfiguration struct {

	// The URL that is used to initiate a playback session for devices that support
	// Apple HLS. The session uses server-side reporting.
	ManifestEndpointPrefix *string
}

// The configuration for pre-roll ad insertion.
type LivePreRollConfiguration struct {

	// The URL for the ad decision server (ADS) for pre-roll ads. This includes the
	// specification of static parameters and placeholders for dynamic parameters. AWS
	// Elemental MediaTailor substitutes player-specific and session-specific
	// parameters as needed when calling the ADS. Alternately, for testing, you can
	// provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The maximum allowed duration for the pre-roll ad avail. AWS Elemental
	// MediaTailor won't play pre-roll ads to exceed this duration, regardless of the
	// total duration of ads that the ADS returns.
	MaxDurationSeconds int32
}

// The configuration for manifest processing rules. Manifest processing rules
// enable customization of the personalized manifests created by MediaTailor.
type ManifestProcessingRules struct {

	// The configuration for Ad Marker Passthrough. Ad marker passthrough can be used
	// to pass ad markers from the origin to the customized manifest.
	AdMarkerPassthrough *AdMarkerPassthrough
}

// The AWSMediaTailor configuration.
type PlaybackConfiguration struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for Avail Suppression. Ad suppression can be used to turn off
	// ad personalization in a long manifest, or if a viewer joins mid-break.
	AvailSuppression *AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break.
	Bumper *Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration

	// The configuration for DASH content.
	DashConfiguration *DashConfiguration

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *ManifestProcessingRules

	// The identifier for the playback configuration.
	Name *string

	// The maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	PersonalizationThresholdSeconds int32

	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string

	// The URL that the player accesses to get a manifest from AWS Elemental
	// MediaTailor. This session will use server-side reporting.
	PlaybackEndpointPrefix *string

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID playback
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string

	// The tags assigned to the playback configuration.
	Tags map[string]string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the master playlist for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string
}

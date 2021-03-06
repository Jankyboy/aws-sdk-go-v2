// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateMatchmakingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether a match that was created with this configuration
	// must be accepted by the matched players. To require acceptance, set to TRUE.
	AcceptanceRequired *bool `type:"boolean"`

	// The length of time (in seconds) to wait for players to accept a proposed
	// match. If any player rejects the match or fails to accept before the timeout,
	// the ticket continues to look for an acceptable match.
	AcceptanceTimeoutSeconds *int64 `min:"1" type:"integer"`

	// The number of player slots in a match to keep open for future players. For
	// example, assume that the configuration's rule set specifies a match for a
	// single 12-person team. If the additional player count is set to 2, only 10
	// players are initially selected for the match.
	AdditionalPlayerCount *int64 `type:"integer"`

	// The method that is used to backfill game sessions created with this matchmaking
	// configuration. Specify MANUAL when your game manages backfill requests manually
	// or does not use the match backfill feature. Specify AUTOMATIC to have GameLift
	// create a StartMatchBackfill request whenever a game session has one or more
	// open slots. Learn more about manual and automatic backfill in Backfill Existing
	// Games with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-backfill.html).
	BackfillMode BackfillMode `type:"string" enum:"true"`

	// Information to add to all events related to the matchmaking configuration.
	CustomEventData *string `type:"string"`

	// A descriptive label that is associated with matchmaking configuration.
	Description *string `min:"1" type:"string"`

	// A set of custom properties for a game session, formatted as key-value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for
	// a successful match.
	GameProperties []GameProperty `type:"list"`

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with
	// a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for
	// a successful match.
	GameSessionData *string `min:"1" type:"string"`

	// Amazon Resource Name (ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html))
	// that is assigned to a GameLift game session queue resource and uniquely identifies
	// it. ARNs are unique across all Regions. These queues are used when placing
	// game sessions for matches that are created with this matchmaking configuration.
	// Queues can be located in any Region.
	GameSessionQueueArns []string `type:"list"`

	// A unique identifier for a matchmaking configuration to update. You can use
	// either the configuration name or ARN value.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// An SNS topic ARN that is set up to receive matchmaking notifications. See
	// Setting up Notifications for Matchmaking (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-notification.html)
	// for more information.
	NotificationTarget *string `type:"string"`

	// The maximum duration, in seconds, that a matchmaking ticket can remain in
	// process before timing out. Requests that fail due to timing out can be resubmitted
	// as needed.
	RequestTimeoutSeconds *int64 `min:"1" type:"integer"`

	// A unique identifier for a matchmaking rule set to use with this configuration.
	// You can use either the rule set name or ARN value. A matchmaking configuration
	// can only use rule sets that are defined in the same Region.
	RuleSetName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateMatchmakingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMatchmakingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMatchmakingConfigurationInput"}
	if s.AcceptanceTimeoutSeconds != nil && *s.AcceptanceTimeoutSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("AcceptanceTimeoutSeconds", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.GameSessionData != nil && len(*s.GameSessionData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionData", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.RequestTimeoutSeconds != nil && *s.RequestTimeoutSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("RequestTimeoutSeconds", 1))
	}
	if s.RuleSetName != nil && len(*s.RuleSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleSetName", 1))
	}
	if s.GameProperties != nil {
		for i, v := range s.GameProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GameProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateMatchmakingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The updated matchmaking configuration.
	Configuration *MatchmakingConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateMatchmakingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMatchmakingConfiguration = "UpdateMatchmakingConfiguration"

// UpdateMatchmakingConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates settings for a FlexMatch matchmaking configuration. These changes
// affect all matches and game sessions that are created after the update. To
// update settings, specify the configuration name to be updated and provide
// the new settings.
//
// Learn more
//
//  Design a FlexMatch Matchmaker (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using UpdateMatchmakingConfigurationRequest.
//    req := client.UpdateMatchmakingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateMatchmakingConfiguration
func (c *Client) UpdateMatchmakingConfigurationRequest(input *UpdateMatchmakingConfigurationInput) UpdateMatchmakingConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateMatchmakingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMatchmakingConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateMatchmakingConfigurationOutput{})

	return UpdateMatchmakingConfigurationRequest{Request: req, Input: input, Copy: c.UpdateMatchmakingConfigurationRequest}
}

// UpdateMatchmakingConfigurationRequest is the request type for the
// UpdateMatchmakingConfiguration API operation.
type UpdateMatchmakingConfigurationRequest struct {
	*aws.Request
	Input *UpdateMatchmakingConfigurationInput
	Copy  func(*UpdateMatchmakingConfigurationInput) UpdateMatchmakingConfigurationRequest
}

// Send marshals and sends the UpdateMatchmakingConfiguration API request.
func (r UpdateMatchmakingConfigurationRequest) Send(ctx context.Context) (*UpdateMatchmakingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMatchmakingConfigurationResponse{
		UpdateMatchmakingConfigurationOutput: r.Request.Data.(*UpdateMatchmakingConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMatchmakingConfigurationResponse is the response type for the
// UpdateMatchmakingConfiguration API operation.
type UpdateMatchmakingConfigurationResponse struct {
	*UpdateMatchmakingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMatchmakingConfiguration request.
func (r *UpdateMatchmakingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

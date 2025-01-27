// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/DeleteMissionProfileRequest
type DeleteMissionProfileInput struct {
	_ struct{} `type:"structure"`

	// MissionProfileId is a required field
	MissionProfileId *string `location:"uri" locationName:"missionProfileId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMissionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMissionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMissionProfileInput"}

	if s.MissionProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MissionProfileId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMissionProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/MissionProfileIdResponse
type DeleteMissionProfileOutput struct {
	_ struct{} `type:"structure"`

	MissionProfileId *string `locationName:"missionProfileId" type:"string"`
}

// String returns the string representation
func (s DeleteMissionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMissionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteMissionProfile = "DeleteMissionProfile"

// DeleteMissionProfileRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Deletes a mission profile.
//
//    // Example sending a request using DeleteMissionProfileRequest.
//    req := client.DeleteMissionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/DeleteMissionProfile
func (c *Client) DeleteMissionProfileRequest(input *DeleteMissionProfileInput) DeleteMissionProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteMissionProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/missionprofile/{missionProfileId}",
	}

	if input == nil {
		input = &DeleteMissionProfileInput{}
	}

	req := c.newRequest(op, input, &DeleteMissionProfileOutput{})
	return DeleteMissionProfileRequest{Request: req, Input: input, Copy: c.DeleteMissionProfileRequest}
}

// DeleteMissionProfileRequest is the request type for the
// DeleteMissionProfile API operation.
type DeleteMissionProfileRequest struct {
	*aws.Request
	Input *DeleteMissionProfileInput
	Copy  func(*DeleteMissionProfileInput) DeleteMissionProfileRequest
}

// Send marshals and sends the DeleteMissionProfile API request.
func (r DeleteMissionProfileRequest) Send(ctx context.Context) (*DeleteMissionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMissionProfileResponse{
		DeleteMissionProfileOutput: r.Request.Data.(*DeleteMissionProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMissionProfileResponse is the response type for the
// DeleteMissionProfile API operation.
type DeleteMissionProfileResponse struct {
	*DeleteMissionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMissionProfile request.
func (r *DeleteMissionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/UpdateDeviceStateRequest
type UpdateDeviceStateInput struct {
	_ struct{} `type:"structure"`

	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	// If true, the device is enabled. If false, the device is disabled.
	Enabled *bool `locationName:"enabled" type:"boolean"`
}

// String returns the string representation
func (s UpdateDeviceStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDeviceStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDeviceStateInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeviceStateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/UpdateDeviceStateResponse
type UpdateDeviceStateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDeviceStateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeviceStateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateDeviceState = "UpdateDeviceState"

// UpdateDeviceStateRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Using a Boolean value (true or false), this operation enables or disables
// the device given a device ID.
//
//    // Example sending a request using UpdateDeviceStateRequest.
//    req := client.UpdateDeviceStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/UpdateDeviceState
func (c *Client) UpdateDeviceStateRequest(input *UpdateDeviceStateInput) UpdateDeviceStateRequest {
	op := &aws.Operation{
		Name:       opUpdateDeviceState,
		HTTPMethod: "PUT",
		HTTPPath:   "/devices/{deviceId}/state",
	}

	if input == nil {
		input = &UpdateDeviceStateInput{}
	}

	req := c.newRequest(op, input, &UpdateDeviceStateOutput{})
	return UpdateDeviceStateRequest{Request: req, Input: input, Copy: c.UpdateDeviceStateRequest}
}

// UpdateDeviceStateRequest is the request type for the
// UpdateDeviceState API operation.
type UpdateDeviceStateRequest struct {
	*aws.Request
	Input *UpdateDeviceStateInput
	Copy  func(*UpdateDeviceStateInput) UpdateDeviceStateRequest
}

// Send marshals and sends the UpdateDeviceState API request.
func (r UpdateDeviceStateRequest) Send(ctx context.Context) (*UpdateDeviceStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDeviceStateResponse{
		UpdateDeviceStateOutput: r.Request.Data.(*UpdateDeviceStateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDeviceStateResponse is the response type for the
// UpdateDeviceState API operation.
type UpdateDeviceStateResponse struct {
	*UpdateDeviceStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDeviceState request.
func (r *UpdateDeviceStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

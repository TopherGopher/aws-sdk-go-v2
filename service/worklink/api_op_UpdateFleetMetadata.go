// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateFleetMetadataRequest
type UpdateFleetMetadataInput struct {
	_ struct{} `type:"structure"`

	// The fleet name to display. The existing DisplayName is unset if null is passed.
	DisplayName *string `type:"string"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`

	// The option to optimize for better performance by routing traffic through
	// the closest AWS Region to users, which may be outside of your home Region.
	OptimizeForEndUserLocation *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateFleetMetadataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFleetMetadataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFleetMetadataInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFleetMetadataInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisplayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OptimizeForEndUserLocation != nil {
		v := *s.OptimizeForEndUserLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OptimizeForEndUserLocation", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateFleetMetadataResponse
type UpdateFleetMetadataOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateFleetMetadataOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFleetMetadataOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateFleetMetadata = "UpdateFleetMetadata"

// UpdateFleetMetadataRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Updates fleet metadata, such as DisplayName.
//
//    // Example sending a request using UpdateFleetMetadataRequest.
//    req := client.UpdateFleetMetadataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/UpdateFleetMetadata
func (c *Client) UpdateFleetMetadataRequest(input *UpdateFleetMetadataInput) UpdateFleetMetadataRequest {
	op := &aws.Operation{
		Name:       opUpdateFleetMetadata,
		HTTPMethod: "POST",
		HTTPPath:   "/UpdateFleetMetadata",
	}

	if input == nil {
		input = &UpdateFleetMetadataInput{}
	}

	req := c.newRequest(op, input, &UpdateFleetMetadataOutput{})
	return UpdateFleetMetadataRequest{Request: req, Input: input, Copy: c.UpdateFleetMetadataRequest}
}

// UpdateFleetMetadataRequest is the request type for the
// UpdateFleetMetadata API operation.
type UpdateFleetMetadataRequest struct {
	*aws.Request
	Input *UpdateFleetMetadataInput
	Copy  func(*UpdateFleetMetadataInput) UpdateFleetMetadataRequest
}

// Send marshals and sends the UpdateFleetMetadata API request.
func (r UpdateFleetMetadataRequest) Send(ctx context.Context) (*UpdateFleetMetadataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFleetMetadataResponse{
		UpdateFleetMetadataOutput: r.Request.Data.(*UpdateFleetMetadataOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFleetMetadataResponse is the response type for the
// UpdateFleetMetadata API operation.
type UpdateFleetMetadataResponse struct {
	*UpdateFleetMetadataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFleetMetadata request.
func (r *UpdateFleetMetadataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

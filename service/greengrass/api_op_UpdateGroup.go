// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateGroupRequest
type UpdateGroupInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateGroupResponse
type UpdateGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateGroup = "UpdateGroup"

// UpdateGroupRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a group.
//
//    // Example sending a request using UpdateGroupRequest.
//    req := client.UpdateGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateGroup
func (c *Client) UpdateGroupRequest(input *UpdateGroupInput) UpdateGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/groups/{GroupId}",
	}

	if input == nil {
		input = &UpdateGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateGroupOutput{})
	return UpdateGroupRequest{Request: req, Input: input, Copy: c.UpdateGroupRequest}
}

// UpdateGroupRequest is the request type for the
// UpdateGroup API operation.
type UpdateGroupRequest struct {
	*aws.Request
	Input *UpdateGroupInput
	Copy  func(*UpdateGroupInput) UpdateGroupRequest
}

// Send marshals and sends the UpdateGroup API request.
func (r UpdateGroupRequest) Send(ctx context.Context) (*UpdateGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGroupResponse{
		UpdateGroupOutput: r.Request.Data.(*UpdateGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGroupResponse is the response type for the
// UpdateGroup API operation.
type UpdateGroupResponse struct {
	*UpdateGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGroup request.
func (r *UpdateGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

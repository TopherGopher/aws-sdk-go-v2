// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/UpdateGroupInput
type UpdateGroupInput struct {
	_ struct{} `type:"structure"`

	// The description of the resource group. Descriptions can have a maximum of
	// 511 characters, including letters, numbers, hyphens, underscores, punctuation,
	// and spaces.
	Description *string `type:"string"`

	// The name of the resource group for which you want to update its description.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/UpdateGroupOutput
type UpdateGroupOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the resource group after it has been updated.
	Group *Group `type:"structure"`
}

// String returns the string representation
func (s UpdateGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Group != nil {
		v := s.Group

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Group", v, metadata)
	}
	return nil
}

const opUpdateGroup = "UpdateGroup"

// UpdateGroupRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Updates an existing group with a new or changed description. You cannot update
// the name of a resource group.
//
//    // Example sending a request using UpdateGroupRequest.
//    req := client.UpdateGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/UpdateGroup
func (c *Client) UpdateGroupRequest(input *UpdateGroupInput) UpdateGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/groups/{GroupName}",
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

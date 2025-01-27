// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/DeleteGroupRequest
type DeleteGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the group that was generated on creation.
	GroupARN *string `min:"1" type:"string"`

	// The case-sensitive name of the group.
	GroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupInput"}
	if s.GroupARN != nil && len(*s.GroupARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupARN", 1))
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
func (s DeleteGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupARN != nil {
		v := *s.GroupARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/DeleteGroupResult
type DeleteGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteGroup = "DeleteGroup"

// DeleteGroupRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Deletes a group resource.
//
//    // Example sending a request using DeleteGroupRequest.
//    req := client.DeleteGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/DeleteGroup
func (c *Client) DeleteGroupRequest(input *DeleteGroupInput) DeleteGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/DeleteGroup",
	}

	if input == nil {
		input = &DeleteGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteGroupOutput{})
	return DeleteGroupRequest{Request: req, Input: input, Copy: c.DeleteGroupRequest}
}

// DeleteGroupRequest is the request type for the
// DeleteGroup API operation.
type DeleteGroupRequest struct {
	*aws.Request
	Input *DeleteGroupInput
	Copy  func(*DeleteGroupInput) DeleteGroupRequest
}

// Send marshals and sends the DeleteGroup API request.
func (r DeleteGroupRequest) Send(ctx context.Context) (*DeleteGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupResponse{
		DeleteGroupOutput: r.Request.Data.(*DeleteGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupResponse is the response type for the
// DeleteGroup API operation.
type DeleteGroupResponse struct {
	*DeleteGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroup request.
func (r *DeleteGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

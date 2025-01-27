// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DisassociateRoleFromGroupRequest
type DisassociateRoleFromGroupInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateRoleFromGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateRoleFromGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateRoleFromGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateRoleFromGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DisassociateRoleFromGroupResponse
type DisassociateRoleFromGroupOutput struct {
	_ struct{} `type:"structure"`

	// The time, in milliseconds since the epoch, when the role was disassociated
	// from the group.
	DisassociatedAt *string `type:"string"`
}

// String returns the string representation
func (s DisassociateRoleFromGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateRoleFromGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DisassociatedAt != nil {
		v := *s.DisassociatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisassociatedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDisassociateRoleFromGroup = "DisassociateRoleFromGroup"

// DisassociateRoleFromGroupRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Disassociates the role from a group.
//
//    // Example sending a request using DisassociateRoleFromGroupRequest.
//    req := client.DisassociateRoleFromGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DisassociateRoleFromGroup
func (c *Client) DisassociateRoleFromGroupRequest(input *DisassociateRoleFromGroupInput) DisassociateRoleFromGroupRequest {
	op := &aws.Operation{
		Name:       opDisassociateRoleFromGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/groups/{GroupId}/role",
	}

	if input == nil {
		input = &DisassociateRoleFromGroupInput{}
	}

	req := c.newRequest(op, input, &DisassociateRoleFromGroupOutput{})
	return DisassociateRoleFromGroupRequest{Request: req, Input: input, Copy: c.DisassociateRoleFromGroupRequest}
}

// DisassociateRoleFromGroupRequest is the request type for the
// DisassociateRoleFromGroup API operation.
type DisassociateRoleFromGroupRequest struct {
	*aws.Request
	Input *DisassociateRoleFromGroupInput
	Copy  func(*DisassociateRoleFromGroupInput) DisassociateRoleFromGroupRequest
}

// Send marshals and sends the DisassociateRoleFromGroup API request.
func (r DisassociateRoleFromGroupRequest) Send(ctx context.Context) (*DisassociateRoleFromGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateRoleFromGroupResponse{
		DisassociateRoleFromGroupOutput: r.Request.Data.(*DisassociateRoleFromGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateRoleFromGroupResponse is the response type for the
// DisassociateRoleFromGroup API operation.
type DisassociateRoleFromGroupResponse struct {
	*DisassociateRoleFromGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateRoleFromGroup request.
func (r *DisassociateRoleFromGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

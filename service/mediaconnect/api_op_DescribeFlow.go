// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/DescribeFlowRequest
type DescribeFlowInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFlowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFlowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFlowInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFlowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful DescribeFlow request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/DescribeFlowResponse
type DescribeFlowOutput struct {
	_ struct{} `type:"structure"`

	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *Flow `locationName:"flow" type:"structure"`

	// Messages that provide the state of the flow.
	Messages *Messages `locationName:"messages" type:"structure"`
}

// String returns the string representation
func (s DescribeFlowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFlowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Flow != nil {
		v := s.Flow

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "flow", v, metadata)
	}
	if s.Messages != nil {
		v := s.Messages

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "messages", v, metadata)
	}
	return nil
}

const opDescribeFlow = "DescribeFlow"

// DescribeFlowRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Displays the details of a flow. The response includes the flow ARN, name,
// and Availability Zone, as well as details about the source, outputs, and
// entitlements.
//
//    // Example sending a request using DescribeFlowRequest.
//    req := client.DescribeFlowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/DescribeFlow
func (c *Client) DescribeFlowRequest(input *DescribeFlowInput) DescribeFlowRequest {
	op := &aws.Operation{
		Name:       opDescribeFlow,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/flows/{flowArn}",
	}

	if input == nil {
		input = &DescribeFlowInput{}
	}

	req := c.newRequest(op, input, &DescribeFlowOutput{})
	return DescribeFlowRequest{Request: req, Input: input, Copy: c.DescribeFlowRequest}
}

// DescribeFlowRequest is the request type for the
// DescribeFlow API operation.
type DescribeFlowRequest struct {
	*aws.Request
	Input *DescribeFlowInput
	Copy  func(*DescribeFlowInput) DescribeFlowRequest
}

// Send marshals and sends the DescribeFlow API request.
func (r DescribeFlowRequest) Send(ctx context.Context) (*DescribeFlowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFlowResponse{
		DescribeFlowOutput: r.Request.Data.(*DescribeFlowOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFlowResponse is the response type for the
// DescribeFlow API operation.
type DescribeFlowResponse struct {
	*DescribeFlowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFlow request.
func (r *DescribeFlowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

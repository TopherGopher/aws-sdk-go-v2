// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeLoadBasedAutoScalingRequest
type DescribeLoadBasedAutoScalingInput struct {
	_ struct{} `type:"structure"`

	// An array of layer IDs.
	//
	// LayerIds is a required field
	LayerIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeLoadBasedAutoScalingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBasedAutoScalingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBasedAutoScalingInput"}

	if s.LayerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeLoadBasedAutoScaling request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeLoadBasedAutoScalingResult
type DescribeLoadBasedAutoScalingOutput struct {
	_ struct{} `type:"structure"`

	// An array of LoadBasedAutoScalingConfiguration objects that describe each
	// layer's configuration.
	LoadBasedAutoScalingConfigurations []LoadBasedAutoScalingConfiguration `type:"list"`
}

// String returns the string representation
func (s DescribeLoadBasedAutoScalingOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLoadBasedAutoScaling = "DescribeLoadBasedAutoScaling"

// DescribeLoadBasedAutoScalingRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes load-based auto scaling configurations for specified layers.
//
// You must specify at least one of the parameters.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeLoadBasedAutoScalingRequest.
//    req := client.DescribeLoadBasedAutoScalingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeLoadBasedAutoScaling
func (c *Client) DescribeLoadBasedAutoScalingRequest(input *DescribeLoadBasedAutoScalingInput) DescribeLoadBasedAutoScalingRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBasedAutoScaling,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBasedAutoScalingInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBasedAutoScalingOutput{})
	return DescribeLoadBasedAutoScalingRequest{Request: req, Input: input, Copy: c.DescribeLoadBasedAutoScalingRequest}
}

// DescribeLoadBasedAutoScalingRequest is the request type for the
// DescribeLoadBasedAutoScaling API operation.
type DescribeLoadBasedAutoScalingRequest struct {
	*aws.Request
	Input *DescribeLoadBasedAutoScalingInput
	Copy  func(*DescribeLoadBasedAutoScalingInput) DescribeLoadBasedAutoScalingRequest
}

// Send marshals and sends the DescribeLoadBasedAutoScaling API request.
func (r DescribeLoadBasedAutoScalingRequest) Send(ctx context.Context) (*DescribeLoadBasedAutoScalingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBasedAutoScalingResponse{
		DescribeLoadBasedAutoScalingOutput: r.Request.Data.(*DescribeLoadBasedAutoScalingOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoadBasedAutoScalingResponse is the response type for the
// DescribeLoadBasedAutoScaling API operation.
type DescribeLoadBasedAutoScalingResponse struct {
	*DescribeLoadBasedAutoScalingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBasedAutoScaling request.
func (r *DescribeLoadBasedAutoScalingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

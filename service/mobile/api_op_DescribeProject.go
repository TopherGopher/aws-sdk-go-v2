// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to request details about a project.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeProjectRequest
type DescribeProjectInput struct {
	_ struct{} `type:"structure"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"querystring" locationName:"projectId" type:"string" required:"true"`

	// If set to true, causes AWS Mobile Hub to synchronize information from other
	// services, e.g., update state of AWS CloudFormation stacks in the AWS Mobile
	// Hub project.
	SyncFromResources *bool `location:"querystring" locationName:"syncFromResources" type:"boolean"`
}

// String returns the string representation
func (s DescribeProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SyncFromResources != nil {
		v := *s.SyncFromResources

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "syncFromResources", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Result structure used for requests of project details.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeProjectResult
type DescribeProjectOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about an AWS Mobile Hub project.
	Details *ProjectDetails `locationName:"details" type:"structure"`
}

// String returns the string representation
func (s DescribeProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Details != nil {
		v := s.Details

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "details", v, metadata)
	}
	return nil
}

const opDescribeProject = "DescribeProject"

// DescribeProjectRequest returns a request value for making API operation for
// AWS Mobile.
//
// Gets details about a project in AWS Mobile Hub.
//
//    // Example sending a request using DescribeProjectRequest.
//    req := client.DescribeProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeProject
func (c *Client) DescribeProjectRequest(input *DescribeProjectInput) DescribeProjectRequest {
	op := &aws.Operation{
		Name:       opDescribeProject,
		HTTPMethod: "GET",
		HTTPPath:   "/project",
	}

	if input == nil {
		input = &DescribeProjectInput{}
	}

	req := c.newRequest(op, input, &DescribeProjectOutput{})
	return DescribeProjectRequest{Request: req, Input: input, Copy: c.DescribeProjectRequest}
}

// DescribeProjectRequest is the request type for the
// DescribeProject API operation.
type DescribeProjectRequest struct {
	*aws.Request
	Input *DescribeProjectInput
	Copy  func(*DescribeProjectInput) DescribeProjectRequest
}

// Send marshals and sends the DescribeProject API request.
func (r DescribeProjectRequest) Send(ctx context.Context) (*DescribeProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProjectResponse{
		DescribeProjectOutput: r.Request.Data.(*DescribeProjectOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProjectResponse is the response type for the
// DescribeProject API operation.
type DescribeProjectResponse struct {
	*DescribeProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProject request.
func (r *DescribeProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

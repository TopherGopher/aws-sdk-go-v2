// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CancelDeploymentJobRequest
type CancelDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// The deployment job ARN to cancel.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelDeploymentJobInput"}

	if s.Job == nil {
		invalidParams.Add(aws.NewErrParamRequired("Job"))
	}
	if s.Job != nil && len(*s.Job) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Job", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelDeploymentJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Job != nil {
		v := *s.Job

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "job", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CancelDeploymentJobResponse
type CancelDeploymentJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelDeploymentJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelDeploymentJob = "CancelDeploymentJob"

// CancelDeploymentJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Cancels the specified deployment job.
//
//    // Example sending a request using CancelDeploymentJobRequest.
//    req := client.CancelDeploymentJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CancelDeploymentJob
func (c *Client) CancelDeploymentJobRequest(input *CancelDeploymentJobInput) CancelDeploymentJobRequest {
	op := &aws.Operation{
		Name:       opCancelDeploymentJob,
		HTTPMethod: "POST",
		HTTPPath:   "/cancelDeploymentJob",
	}

	if input == nil {
		input = &CancelDeploymentJobInput{}
	}

	req := c.newRequest(op, input, &CancelDeploymentJobOutput{})
	return CancelDeploymentJobRequest{Request: req, Input: input, Copy: c.CancelDeploymentJobRequest}
}

// CancelDeploymentJobRequest is the request type for the
// CancelDeploymentJob API operation.
type CancelDeploymentJobRequest struct {
	*aws.Request
	Input *CancelDeploymentJobInput
	Copy  func(*CancelDeploymentJobInput) CancelDeploymentJobRequest
}

// Send marshals and sends the CancelDeploymentJob API request.
func (r CancelDeploymentJobRequest) Send(ctx context.Context) (*CancelDeploymentJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelDeploymentJobResponse{
		CancelDeploymentJobOutput: r.Request.Data.(*CancelDeploymentJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelDeploymentJobResponse is the response type for the
// CancelDeploymentJob API operation.
type CancelDeploymentJobResponse struct {
	*CancelDeploymentJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelDeploymentJob request.
func (r *CancelDeploymentJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

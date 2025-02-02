// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RestartSimulationJobRequest
type RestartSimulationJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the simulation job.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RestartSimulationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestartSimulationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestartSimulationJobInput"}

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
func (s RestartSimulationJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Job != nil {
		v := *s.Job

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "job", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RestartSimulationJobResponse
type RestartSimulationJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RestartSimulationJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RestartSimulationJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRestartSimulationJob = "RestartSimulationJob"

// RestartSimulationJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Restarts a running simulation job.
//
//    // Example sending a request using RestartSimulationJobRequest.
//    req := client.RestartSimulationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/RestartSimulationJob
func (c *Client) RestartSimulationJobRequest(input *RestartSimulationJobInput) RestartSimulationJobRequest {
	op := &aws.Operation{
		Name:       opRestartSimulationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/restartSimulationJob",
	}

	if input == nil {
		input = &RestartSimulationJobInput{}
	}

	req := c.newRequest(op, input, &RestartSimulationJobOutput{})
	return RestartSimulationJobRequest{Request: req, Input: input, Copy: c.RestartSimulationJobRequest}
}

// RestartSimulationJobRequest is the request type for the
// RestartSimulationJob API operation.
type RestartSimulationJobRequest struct {
	*aws.Request
	Input *RestartSimulationJobInput
	Copy  func(*RestartSimulationJobInput) RestartSimulationJobRequest
}

// Send marshals and sends the RestartSimulationJob API request.
func (r RestartSimulationJobRequest) Send(ctx context.Context) (*RestartSimulationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestartSimulationJobResponse{
		RestartSimulationJobOutput: r.Request.Data.(*RestartSimulationJobOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestartSimulationJobResponse is the response type for the
// RestartSimulationJob API operation.
type RestartSimulationJobResponse struct {
	*RestartSimulationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestartSimulationJob request.
func (r *RestartSimulationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

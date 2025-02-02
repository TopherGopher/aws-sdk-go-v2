// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribePipelineRequest
type DescribePipelineInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline whose information is retrieved.
	//
	// PipelineName is a required field
	PipelineName *string `location:"uri" locationName:"pipelineName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePipelineInput"}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePipelineInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PipelineName != nil {
		v := *s.PipelineName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "pipelineName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribePipelineResponse
type DescribePipelineOutput struct {
	_ struct{} `type:"structure"`

	// A "Pipeline" object that contains information about the pipeline.
	Pipeline *Pipeline `locationName:"pipeline" type:"structure"`
}

// String returns the string representation
func (s DescribePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Pipeline != nil {
		v := s.Pipeline

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "pipeline", v, metadata)
	}
	return nil
}

const opDescribePipeline = "DescribePipeline"

// DescribePipelineRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a pipeline.
//
//    // Example sending a request using DescribePipelineRequest.
//    req := client.DescribePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribePipeline
func (c *Client) DescribePipelineRequest(input *DescribePipelineInput) DescribePipelineRequest {
	op := &aws.Operation{
		Name:       opDescribePipeline,
		HTTPMethod: "GET",
		HTTPPath:   "/pipelines/{pipelineName}",
	}

	if input == nil {
		input = &DescribePipelineInput{}
	}

	req := c.newRequest(op, input, &DescribePipelineOutput{})
	return DescribePipelineRequest{Request: req, Input: input, Copy: c.DescribePipelineRequest}
}

// DescribePipelineRequest is the request type for the
// DescribePipeline API operation.
type DescribePipelineRequest struct {
	*aws.Request
	Input *DescribePipelineInput
	Copy  func(*DescribePipelineInput) DescribePipelineRequest
}

// Send marshals and sends the DescribePipeline API request.
func (r DescribePipelineRequest) Send(ctx context.Context) (*DescribePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePipelineResponse{
		DescribePipelineOutput: r.Request.Data.(*DescribePipelineOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePipelineResponse is the response type for the
// DescribePipeline API operation.
type DescribePipelineResponse struct {
	*DescribePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePipeline request.
func (r *DescribePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

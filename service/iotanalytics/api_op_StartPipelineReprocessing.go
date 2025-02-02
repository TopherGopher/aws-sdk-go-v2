// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/StartPipelineReprocessingRequest
type StartPipelineReprocessingInput struct {
	_ struct{} `type:"structure"`

	// The end time (exclusive) of raw message data that is reprocessed.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The name of the pipeline on which to start reprocessing.
	//
	// PipelineName is a required field
	PipelineName *string `location:"uri" locationName:"pipelineName" min:"1" type:"string" required:"true"`

	// The start time (inclusive) of raw message data that is reprocessed.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s StartPipelineReprocessingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartPipelineReprocessingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartPipelineReprocessingInput"}

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
func (s StartPipelineReprocessingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.PipelineName != nil {
		v := *s.PipelineName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "pipelineName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/StartPipelineReprocessingResponse
type StartPipelineReprocessingOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline reprocessing activity that was started.
	ReprocessingId *string `locationName:"reprocessingId" type:"string"`
}

// String returns the string representation
func (s StartPipelineReprocessingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartPipelineReprocessingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ReprocessingId != nil {
		v := *s.ReprocessingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reprocessingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartPipelineReprocessing = "StartPipelineReprocessing"

// StartPipelineReprocessingRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Starts the reprocessing of raw message data through the pipeline.
//
//    // Example sending a request using StartPipelineReprocessingRequest.
//    req := client.StartPipelineReprocessingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/StartPipelineReprocessing
func (c *Client) StartPipelineReprocessingRequest(input *StartPipelineReprocessingInput) StartPipelineReprocessingRequest {
	op := &aws.Operation{
		Name:       opStartPipelineReprocessing,
		HTTPMethod: "POST",
		HTTPPath:   "/pipelines/{pipelineName}/reprocessing",
	}

	if input == nil {
		input = &StartPipelineReprocessingInput{}
	}

	req := c.newRequest(op, input, &StartPipelineReprocessingOutput{})
	return StartPipelineReprocessingRequest{Request: req, Input: input, Copy: c.StartPipelineReprocessingRequest}
}

// StartPipelineReprocessingRequest is the request type for the
// StartPipelineReprocessing API operation.
type StartPipelineReprocessingRequest struct {
	*aws.Request
	Input *StartPipelineReprocessingInput
	Copy  func(*StartPipelineReprocessingInput) StartPipelineReprocessingRequest
}

// Send marshals and sends the StartPipelineReprocessing API request.
func (r StartPipelineReprocessingRequest) Send(ctx context.Context) (*StartPipelineReprocessingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartPipelineReprocessingResponse{
		StartPipelineReprocessingOutput: r.Request.Data.(*StartPipelineReprocessingOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartPipelineReprocessingResponse is the response type for the
// StartPipelineReprocessing API operation.
type StartPipelineReprocessingResponse struct {
	*StartPipelineReprocessingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartPipelineReprocessing request.
func (r *StartPipelineReprocessingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

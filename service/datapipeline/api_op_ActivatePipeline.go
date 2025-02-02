// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ActivatePipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ActivatePipelineInput
type ActivatePipelineInput struct {
	_ struct{} `type:"structure"`

	// A list of parameter values to pass to the pipeline at activation.
	ParameterValues []ParameterValue `locationName:"parameterValues" type:"list"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// The date and time to resume the pipeline. By default, the pipeline resumes
	// from the last completed execution.
	StartTimestamp *time.Time `locationName:"startTimestamp" type:"timestamp"`
}

// String returns the string representation
func (s ActivatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ActivatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ActivatePipelineInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}
	if s.ParameterValues != nil {
		for i, v := range s.ParameterValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ActivatePipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ActivatePipelineOutput
type ActivatePipelineOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ActivatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opActivatePipeline = "ActivatePipeline"

// ActivatePipelineRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Validates the specified pipeline and starts processing pipeline tasks. If
// the pipeline does not pass validation, activation fails.
//
// If you need to pause the pipeline to investigate an issue with a component,
// such as a data source or script, call DeactivatePipeline.
//
// To activate a finished pipeline, modify the end date for the pipeline and
// then activate it.
//
//    // Example sending a request using ActivatePipelineRequest.
//    req := client.ActivatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ActivatePipeline
func (c *Client) ActivatePipelineRequest(input *ActivatePipelineInput) ActivatePipelineRequest {
	op := &aws.Operation{
		Name:       opActivatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ActivatePipelineInput{}
	}

	req := c.newRequest(op, input, &ActivatePipelineOutput{})
	return ActivatePipelineRequest{Request: req, Input: input, Copy: c.ActivatePipelineRequest}
}

// ActivatePipelineRequest is the request type for the
// ActivatePipeline API operation.
type ActivatePipelineRequest struct {
	*aws.Request
	Input *ActivatePipelineInput
	Copy  func(*ActivatePipelineInput) ActivatePipelineRequest
}

// Send marshals and sends the ActivatePipeline API request.
func (r ActivatePipelineRequest) Send(ctx context.Context) (*ActivatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ActivatePipelineResponse{
		ActivatePipelineOutput: r.Request.Data.(*ActivatePipelineOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ActivatePipelineResponse is the response type for the
// ActivatePipeline API operation.
type ActivatePipelineResponse struct {
	*ActivatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ActivatePipeline request.
func (r *ActivatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

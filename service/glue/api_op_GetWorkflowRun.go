// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRunRequest
type GetWorkflowRunInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool `type:"boolean"`

	// Name of the workflow being run.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The ID of the workflow run.
	//
	// RunId is a required field
	RunId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWorkflowRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkflowRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkflowRunInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RunId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RunId"))
	}
	if s.RunId != nil && len(*s.RunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RunId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRunResponse
type GetWorkflowRunOutput struct {
	_ struct{} `type:"structure"`

	// The requested workflow run metadata.
	Run *WorkflowRun `type:"structure"`
}

// String returns the string representation
func (s GetWorkflowRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWorkflowRun = "GetWorkflowRun"

// GetWorkflowRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the metadata for a given workflow run.
//
//    // Example sending a request using GetWorkflowRunRequest.
//    req := client.GetWorkflowRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRun
func (c *Client) GetWorkflowRunRequest(input *GetWorkflowRunInput) GetWorkflowRunRequest {
	op := &aws.Operation{
		Name:       opGetWorkflowRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWorkflowRunInput{}
	}

	req := c.newRequest(op, input, &GetWorkflowRunOutput{})
	return GetWorkflowRunRequest{Request: req, Input: input, Copy: c.GetWorkflowRunRequest}
}

// GetWorkflowRunRequest is the request type for the
// GetWorkflowRun API operation.
type GetWorkflowRunRequest struct {
	*aws.Request
	Input *GetWorkflowRunInput
	Copy  func(*GetWorkflowRunInput) GetWorkflowRunRequest
}

// Send marshals and sends the GetWorkflowRun API request.
func (r GetWorkflowRunRequest) Send(ctx context.Context) (*GetWorkflowRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWorkflowRunResponse{
		GetWorkflowRunOutput: r.Request.Data.(*GetWorkflowRunOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWorkflowRunResponse is the response type for the
// GetWorkflowRun API operation.
type GetWorkflowRunResponse struct {
	*GetWorkflowRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWorkflowRun request.
func (r *GetWorkflowRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

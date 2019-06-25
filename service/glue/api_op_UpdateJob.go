// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateJobRequest
type UpdateJobInput struct {
	_ struct{} `type:"structure"`

	// Name of the job definition to update.
	//
	// JobName is a required field
	JobName *string `min:"1" type:"string" required:"true"`

	// Specifies the values with which to update the job definition.
	//
	// JobUpdate is a required field
	JobUpdate *JobUpdate `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobInput"}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if s.JobUpdate == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobUpdate"))
	}
	if s.JobUpdate != nil {
		if err := s.JobUpdate.Validate(); err != nil {
			invalidParams.AddNested("JobUpdate", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateJobResponse
type UpdateJobOutput struct {
	_ struct{} `type:"structure"`

	// Returns the name of the updated job definition.
	JobName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateJob = "UpdateJob"

// UpdateJobRequest returns a request value for making API operation for
// AWS Glue.
//
// Updates an existing job definition.
//
//    // Example sending a request using UpdateJobRequest.
//    req := client.UpdateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateJob
func (c *Client) UpdateJobRequest(input *UpdateJobInput) UpdateJobRequest {
	op := &aws.Operation{
		Name:       opUpdateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateJobInput{}
	}

	req := c.newRequest(op, input, &UpdateJobOutput{})
	return UpdateJobRequest{Request: req, Input: input, Copy: c.UpdateJobRequest}
}

// UpdateJobRequest is the request type for the
// UpdateJob API operation.
type UpdateJobRequest struct {
	*aws.Request
	Input *UpdateJobInput
	Copy  func(*UpdateJobInput) UpdateJobRequest
}

// Send marshals and sends the UpdateJob API request.
func (r UpdateJobRequest) Send(ctx context.Context) (*UpdateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobResponse{
		UpdateJobOutput: r.Request.Data.(*UpdateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobResponse is the response type for the
// UpdateJob API operation.
type UpdateJobResponse struct {
	*UpdateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJob request.
func (r *UpdateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Query a job by sending a request with the job ID.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/GetJobRequest
type GetJobInput struct {
	_ struct{} `type:"structure"`

	// the job ID of the job.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful get job requests will return an OK message and the job JSON.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/GetJobResponse
type GetJobOutput struct {
	_ struct{} `type:"structure"`

	// Each job converts an input file into an output file or files. For more information,
	// see the User Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
	Job *Job `locationName:"job" type:"structure"`
}

// String returns the string representation
func (s GetJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Job != nil {
		v := s.Job

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "job", v, metadata)
	}
	return nil
}

const opGetJob = "GetJob"

// GetJobRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Retrieve the JSON for a specific completed transcoding job.
//
//    // Example sending a request using GetJobRequest.
//    req := client.GetJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/GetJob
func (c *Client) GetJobRequest(input *GetJobInput) GetJobRequest {
	op := &aws.Operation{
		Name:       opGetJob,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-08-29/jobs/{id}",
	}

	if input == nil {
		input = &GetJobInput{}
	}

	req := c.newRequest(op, input, &GetJobOutput{})
	return GetJobRequest{Request: req, Input: input, Copy: c.GetJobRequest}
}

// GetJobRequest is the request type for the
// GetJob API operation.
type GetJobRequest struct {
	*aws.Request
	Input *GetJobInput
	Copy  func(*GetJobInput) GetJobRequest
}

// Send marshals and sends the GetJob API request.
func (r GetJobRequest) Send(ctx context.Context) (*GetJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobResponse{
		GetJobOutput: r.Request.Data.(*GetJobOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobResponse is the response type for the
// GetJob API operation.
type GetJobResponse struct {
	*GetJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJob request.
func (r *GetJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

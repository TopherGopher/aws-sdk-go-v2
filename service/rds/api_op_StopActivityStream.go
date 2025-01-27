// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StopActivityStreamRequest
type StopActivityStreamInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether or not the database activity stream is to stop as soon
	// as possible, regardless of the maintenance window for the database.
	ApplyImmediately *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the DB cluster for the database activity
	// stream. For example, arn:aws:rds:us-east-1:12345667890:cluster:das-cluster.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopActivityStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopActivityStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopActivityStreamInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StopActivityStreamResponse
type StopActivityStreamOutput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon Kinesis data stream used for the database activity
	// stream.
	KinesisStreamName *string `type:"string"`

	// The AWS KMS key identifier used for encrypting messages in the database activity
	// stream.
	KmsKeyId *string `type:"string"`

	// The status of the database activity stream.
	Status ActivityStreamStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StopActivityStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopActivityStream = "StopActivityStream"

// StopActivityStreamRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Stops a database activity stream that was started using the AWS console,
// the start-activity-stream AWS CLI command, or the StartActivityStream action.
//
// For more information, see Database Activity Streams (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html)
// in the Amazon Aurora User Guide.
//
//    // Example sending a request using StopActivityStreamRequest.
//    req := client.StopActivityStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StopActivityStream
func (c *Client) StopActivityStreamRequest(input *StopActivityStreamInput) StopActivityStreamRequest {
	op := &aws.Operation{
		Name:       opStopActivityStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopActivityStreamInput{}
	}

	req := c.newRequest(op, input, &StopActivityStreamOutput{})
	return StopActivityStreamRequest{Request: req, Input: input, Copy: c.StopActivityStreamRequest}
}

// StopActivityStreamRequest is the request type for the
// StopActivityStream API operation.
type StopActivityStreamRequest struct {
	*aws.Request
	Input *StopActivityStreamInput
	Copy  func(*StopActivityStreamInput) StopActivityStreamRequest
}

// Send marshals and sends the StopActivityStream API request.
func (r StopActivityStreamRequest) Send(ctx context.Context) (*StopActivityStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopActivityStreamResponse{
		StopActivityStreamOutput: r.Request.Data.(*StopActivityStreamOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopActivityStreamResponse is the response type for the
// StopActivityStream API operation.
type StopActivityStreamResponse struct {
	*StopActivityStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopActivityStream request.
func (r *StopActivityStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

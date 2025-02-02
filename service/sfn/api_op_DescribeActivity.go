// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DescribeActivityInput
type DescribeActivityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the activity to describe.
	//
	// ActivityArn is a required field
	ActivityArn *string `locationName:"activityArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeActivityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeActivityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeActivityInput"}

	if s.ActivityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivityArn"))
	}
	if s.ActivityArn != nil && len(*s.ActivityArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ActivityArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DescribeActivityOutput
type DescribeActivityOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the activity.
	//
	// ActivityArn is a required field
	ActivityArn *string `locationName:"activityArn" min:"1" type:"string" required:"true"`

	// The date the activity is created.
	//
	// CreationDate is a required field
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" required:"true"`

	// The name of the activity.
	//
	// A name must not contain:
	//
	//    * whitespace
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeActivityOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeActivity = "DescribeActivity"

// DescribeActivityRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Describes an activity.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using DescribeActivityRequest.
//    req := client.DescribeActivityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DescribeActivity
func (c *Client) DescribeActivityRequest(input *DescribeActivityInput) DescribeActivityRequest {
	op := &aws.Operation{
		Name:       opDescribeActivity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeActivityInput{}
	}

	req := c.newRequest(op, input, &DescribeActivityOutput{})
	return DescribeActivityRequest{Request: req, Input: input, Copy: c.DescribeActivityRequest}
}

// DescribeActivityRequest is the request type for the
// DescribeActivity API operation.
type DescribeActivityRequest struct {
	*aws.Request
	Input *DescribeActivityInput
	Copy  func(*DescribeActivityInput) DescribeActivityRequest
}

// Send marshals and sends the DescribeActivity API request.
func (r DescribeActivityRequest) Send(ctx context.Context) (*DescribeActivityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeActivityResponse{
		DescribeActivityOutput: r.Request.Data.(*DescribeActivityOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeActivityResponse is the response type for the
// DescribeActivity API operation.
type DescribeActivityResponse struct {
	*DescribeActivityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeActivity request.
func (r *DescribeActivityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

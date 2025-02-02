// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeUserRequest
type DescribeUserInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the organization under which the user exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The identifier for the user to be described.
	//
	// UserId is a required field
	UserId *string `min:"12" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeUserResponse
type DescribeUserOutput struct {
	_ struct{} `type:"structure"`

	// The date and time at which the user was disabled for Amazon WorkMail usage,
	// in UNIX epoch time format.
	DisabledDate *time.Time `type:"timestamp"`

	// The display name of the user.
	DisplayName *string `type:"string"`

	// The email of the user.
	Email *string `min:"1" type:"string"`

	// The date and time at which the user was enabled for Amazon WorkMail usage,
	// in UNIX epoch time format.
	EnabledDate *time.Time `type:"timestamp"`

	// The name for the user.
	Name *string `min:"1" type:"string"`

	// The state of a user: enabled (registered to Amazon WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State EntityState `type:"string" enum:"true"`

	// The identifier for the described user.
	UserId *string `min:"12" type:"string"`

	// In certain cases, other entities are modeled as users. If interoperability
	// is enabled, resources are imported into Amazon WorkMail as users. Because
	// different WorkMail organizations rely on different directory types, administrators
	// can distinguish between an unregistered user (account is disabled and has
	// a user role) and the directory administrators. The values are USER, RESOURCE,
	// and SYSTEM_USER.
	UserRole UserRole `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUser = "DescribeUser"

// DescribeUserRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Provides information regarding the user.
//
//    // Example sending a request using DescribeUserRequest.
//    req := client.DescribeUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeUser
func (c *Client) DescribeUserRequest(input *DescribeUserInput) DescribeUserRequest {
	op := &aws.Operation{
		Name:       opDescribeUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserInput{}
	}

	req := c.newRequest(op, input, &DescribeUserOutput{})
	return DescribeUserRequest{Request: req, Input: input, Copy: c.DescribeUserRequest}
}

// DescribeUserRequest is the request type for the
// DescribeUser API operation.
type DescribeUserRequest struct {
	*aws.Request
	Input *DescribeUserInput
	Copy  func(*DescribeUserInput) DescribeUserRequest
}

// Send marshals and sends the DescribeUser API request.
func (r DescribeUserRequest) Send(ctx context.Context) (*DescribeUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserResponse{
		DescribeUserOutput: r.Request.Data.(*DescribeUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserResponse is the response type for the
// DescribeUser API operation.
type DescribeUserResponse struct {
	*DescribeUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUser request.
func (r *DescribeUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

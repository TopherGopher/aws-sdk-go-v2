// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/LogoutUserRequest
type LogoutUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s LogoutUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LogoutUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LogoutUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LogoutUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/LogoutUserResponse
type LogoutUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s LogoutUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LogoutUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opLogoutUser = "LogoutUser"

// LogoutUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Logs out the specified user from all of the devices they are currently logged
// into.
//
//    // Example sending a request using LogoutUserRequest.
//    req := client.LogoutUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/LogoutUser
func (c *Client) LogoutUserRequest(input *LogoutUserInput) LogoutUserRequest {
	op := &aws.Operation{
		Name:       opLogoutUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users/{userId}?operation=logout",
	}

	if input == nil {
		input = &LogoutUserInput{}
	}

	req := c.newRequest(op, input, &LogoutUserOutput{})
	return LogoutUserRequest{Request: req, Input: input, Copy: c.LogoutUserRequest}
}

// LogoutUserRequest is the request type for the
// LogoutUser API operation.
type LogoutUserRequest struct {
	*aws.Request
	Input *LogoutUserInput
	Copy  func(*LogoutUserInput) LogoutUserRequest
}

// Send marshals and sends the LogoutUser API request.
func (r LogoutUserRequest) Send(ctx context.Context) (*LogoutUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LogoutUserResponse{
		LogoutUserOutput: r.Request.Data.(*LogoutUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// LogoutUserResponse is the response type for the
// LogoutUser API operation.
type LogoutUserResponse struct {
	*LogoutUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LogoutUser request.
func (r *LogoutUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

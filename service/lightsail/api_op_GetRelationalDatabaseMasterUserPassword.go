// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseMasterUserPasswordRequest
type GetRelationalDatabaseMasterUserPasswordInput struct {
	_ struct{} `type:"structure"`

	// The password version to return.
	//
	// Specifying CURRENT or PREVIOUS returns the current or previous passwords
	// respectively. Specifying PENDING returns the newest version of the password
	// that will rotate to CURRENT. After the PENDING password rotates to CURRENT,
	// the PENDING password is no longer available.
	//
	// Default: CURRENT
	PasswordVersion RelationalDatabasePasswordVersion `locationName:"passwordVersion" type:"string" enum:"true"`

	// The name of your database for which to get the master user password.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMasterUserPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseMasterUserPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseMasterUserPasswordInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseMasterUserPasswordResult
type GetRelationalDatabaseMasterUserPasswordOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the specified version of the master user password was
	// created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The master user password for the password version specified.
	MasterUserPassword *string `locationName:"masterUserPassword" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseMasterUserPasswordOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseMasterUserPassword = "GetRelationalDatabaseMasterUserPassword"

// GetRelationalDatabaseMasterUserPasswordRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the current, previous, or pending versions of the master user password
// for a Lightsail database.
//
// The asdf operation GetRelationalDatabaseMasterUserPassword supports tag-based
// access control via resource tags applied to the resource identified by relationalDatabaseName.
//
//    // Example sending a request using GetRelationalDatabaseMasterUserPasswordRequest.
//    req := client.GetRelationalDatabaseMasterUserPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseMasterUserPassword
func (c *Client) GetRelationalDatabaseMasterUserPasswordRequest(input *GetRelationalDatabaseMasterUserPasswordInput) GetRelationalDatabaseMasterUserPasswordRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseMasterUserPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseMasterUserPasswordInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseMasterUserPasswordOutput{})
	return GetRelationalDatabaseMasterUserPasswordRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseMasterUserPasswordRequest}
}

// GetRelationalDatabaseMasterUserPasswordRequest is the request type for the
// GetRelationalDatabaseMasterUserPassword API operation.
type GetRelationalDatabaseMasterUserPasswordRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseMasterUserPasswordInput
	Copy  func(*GetRelationalDatabaseMasterUserPasswordInput) GetRelationalDatabaseMasterUserPasswordRequest
}

// Send marshals and sends the GetRelationalDatabaseMasterUserPassword API request.
func (r GetRelationalDatabaseMasterUserPasswordRequest) Send(ctx context.Context) (*GetRelationalDatabaseMasterUserPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseMasterUserPasswordResponse{
		GetRelationalDatabaseMasterUserPasswordOutput: r.Request.Data.(*GetRelationalDatabaseMasterUserPasswordOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseMasterUserPasswordResponse is the response type for the
// GetRelationalDatabaseMasterUserPassword API operation.
type GetRelationalDatabaseMasterUserPasswordResponse struct {
	*GetRelationalDatabaseMasterUserPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseMasterUserPassword request.
func (r *GetRelationalDatabaseMasterUserPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

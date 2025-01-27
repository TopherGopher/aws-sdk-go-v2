// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateManagedInstanceRoleRequest
type UpdateManagedInstanceRoleInput struct {
	_ struct{} `type:"structure"`

	// The IAM role you want to assign or change.
	//
	// IamRole is a required field
	IamRole *string `type:"string" required:"true"`

	// The ID of the managed instance where you want to update the role.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateManagedInstanceRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateManagedInstanceRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateManagedInstanceRoleInput"}

	if s.IamRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamRole"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateManagedInstanceRoleResult
type UpdateManagedInstanceRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateManagedInstanceRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateManagedInstanceRole = "UpdateManagedInstanceRole"

// UpdateManagedInstanceRoleRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Assigns or changes an Amazon Identity and Access Management (IAM) role for
// the managed instance.
//
//    // Example sending a request using UpdateManagedInstanceRoleRequest.
//    req := client.UpdateManagedInstanceRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateManagedInstanceRole
func (c *Client) UpdateManagedInstanceRoleRequest(input *UpdateManagedInstanceRoleInput) UpdateManagedInstanceRoleRequest {
	op := &aws.Operation{
		Name:       opUpdateManagedInstanceRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateManagedInstanceRoleInput{}
	}

	req := c.newRequest(op, input, &UpdateManagedInstanceRoleOutput{})
	return UpdateManagedInstanceRoleRequest{Request: req, Input: input, Copy: c.UpdateManagedInstanceRoleRequest}
}

// UpdateManagedInstanceRoleRequest is the request type for the
// UpdateManagedInstanceRole API operation.
type UpdateManagedInstanceRoleRequest struct {
	*aws.Request
	Input *UpdateManagedInstanceRoleInput
	Copy  func(*UpdateManagedInstanceRoleInput) UpdateManagedInstanceRoleRequest
}

// Send marshals and sends the UpdateManagedInstanceRole API request.
func (r UpdateManagedInstanceRoleRequest) Send(ctx context.Context) (*UpdateManagedInstanceRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateManagedInstanceRoleResponse{
		UpdateManagedInstanceRoleOutput: r.Request.Data.(*UpdateManagedInstanceRoleOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateManagedInstanceRoleResponse is the response type for the
// UpdateManagedInstanceRole API operation.
type UpdateManagedInstanceRoleResponse struct {
	*UpdateManagedInstanceRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateManagedInstanceRole request.
func (r *UpdateManagedInstanceRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
